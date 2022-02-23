package bstcount

type BSTNodeWithCount struct {
	*bstNodeWithCount
}

func New(nums ...int) *BSTNodeWithCount {
	var root *bstNodeWithCount // root := (*treeNode)(nil) would also do, but not &treeNode{} (a node with val 0)
	for _, num := range nums {
		root, _ = insert(root, num)
	}
	return &BSTNodeWithCount{root}
}

func (bst *BSTNodeWithCount) Insert(val int) (count int) {
	bst.bstNodeWithCount, count = insert(bst.bstNodeWithCount, val)
	return
}

func (bst *BSTNodeWithCount) Count(val int) *int {
	return bst.valCount(val)
}

type bstNodeWithCount struct {
	val         int
	left, right *bstNodeWithCount

	count int
}

func (bst *bstNodeWithCount) valCount(val int) *int {
	if bst == nil {
		return nil
	}

	if val < bst.val {
		return bst.left.valCount(val)
	} else if val > bst.val {
		return bst.right.valCount(val)
	} else {
		return &bst.count
	}
}

func insert(old *bstNodeWithCount, val int) (new *bstNodeWithCount, count int) {
	if old == nil {
		return &bstNodeWithCount{val: val, count: 1}, 1
	}

	if val < old.val {
		old.left, count = insert(old.left, val)
	} else if val > old.val {
		old.right, count = insert(old.right, val)
	} else {
		old.count++
		count = old.count
	}
	return old, count
}
