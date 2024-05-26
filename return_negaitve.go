/*
In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?
MakeNegative(1)    // return -1
MakeNegative(-5)   // return -5
MakeNegative(0)    // return 0
*/

package kata

func MakeNegative(x int) int {
  if x <= 0 {
    return x
  }
  return -x
}
