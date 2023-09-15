# Remove Element

**Difficulty**: Easy

## Problem Description

Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` in-place. The order of the elements may be changed. Then return the number of elements in `nums` which are not equal to `val`.

Consider the number of elements in `nums` which are not equal to `val` be `k`. To get accepted, you need to do the following things:

1. Change the array `nums` such that the first `k` elements of `nums` contain the elements which are not equal to `val`. The remaining elements of `nums` are not important as well as the size of `nums`.
2. Return `k`.

**Custom Judge**:

The judge will test your solution with the following code:

```python
nums = [...]  # Input array
val = ...     # Value to remove
expectedNums = [...]  # The expected answer with the correct length.
                      # It is sorted with no values equaling val.

k = removeElement(nums, val)  # Calls your implementation

assert k == len(expectedNums)
nums[:k].sort()  # Sort the first k elements of nums
for i in range(k):
    assert nums[i] == expectedNums[i]
```

If all assertions pass, then your solution will be accepted.

## Example 1

**Input**:

```
nums = [3,2,2,3], val = 3
```

**Output**:

```
2
```

**Explanation**:

Your function should return `k = 2`, with the first two elements of `nums` being `2`. It does not matter what you leave beyond the returned `k` (hence they are underscores).

## Example 2

**Input**:

```
nums = [0,1,2,2,3,0,4,2], val = 2
```

**Output**:

```
5
```

**Explanation**:

Your function should return `k = 5`, with the first five elements of `nums` containing `0`, `0`, `1`, `3`, and `4`. Note that the five elements can be returned in any order. It does not matter what you leave beyond the returned `k` (hence they are underscores).

## Constraints

- `0 <= nums.length <= 100`
- `0 <= nums[i] <= 50`
- `0 <= val <= 100`