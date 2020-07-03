package examples

// + Section 7.3 = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// f(x1, x2, x3, x4) = 4 * x1 * x2 - 7 * x2 + 3 * x3
//                   = x3 + 3 * x4 - 7 * x5
//                   = f2(x3, x4, x5)

// x2          => x5
// x3          => x4
// 4 * x1 * x2 => x3

// f1(x1, x2) = 4 * x1 * x2
//            = (4 * x1) * x2
//            = p1(x1, x2) * p2(x1, x2)
//            = 24

// p1(x1, x2) = c_{0} + Sigma_{i = 1}^{m - 1} c_{i} * x_{i}
//            = c_{0} + c_{1} * x_{1}
//            = 4 * 3

// p2(x1, x2) = d_{0} + Sigma_{i = 1}^{m - 1} d_{i} * x_{i}
//            = d_{0} + d_{1} * x_{1} + d_{2} * x_{2}
//            = 1 * 2

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q1: I_{1} = { 0, 1, 2, 3 }

// t1(x) = x - r1

// v1_{0}(x) = c_{0} = 0
// v1_{1}(x) = c_{1} = 4
// v1_{2}(x) = c_{2} = 0
// v1_{3}(x)         = 0

// w1_{0}(x) = d_{0} = 0
// w1_{1}(x) = d_{1} = 0
// w1_{2}(x) = d_{2} = 1
// w1_{3}(x)         = 0

// y1_{0}(x)         = 0
// y1_{1}(x)         = 0
// y1_{2}(x)         = 0
// y1_{3}(x)         = 1

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                             *
 *  x1       x2  a1       a2  3       2    p1(a1, a2) * p2(a1, 2) - a3 = (4 * a1) * (a2) - a3  *
 *    \     /      \     /     \     /                                 = 4 * 3 * 2 - 24        *
 *   4 *   /      4 *   /     4 *   /                                  = 0                     *
 *      \ /          \ /         \ /                                                           *
 *       *            *           *                                                            *
 *       |            |           |                                                            *
 *       |            |           |                                                            *
 *       |            |           |                                                            *
 *   f(x1, x2)        a3          24                                                           *
 *                                                                                             *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// f2(x3, x4, x5) = x3 + 3 * x4 - 7 * x5
//                = 1 * (x3 + 3 * x4 - 7 * x5)
//                = p1(x3, x4, x5) * p2(x3, x4, x5)
//                = 24 + 3 * 1 - 7 * 2
//                = 13

// p1(x3, x4, x5) = c_{3} + Sigma_{i = 4}^{m - 1} c_{i} * x_{i}
//                = c_{3}
//                = 1

// p2(x3, x4, x5) = d_{3} + Sigma_{i = 4}^{m - 1} d_{i} * x_{i}
//                = d_{3} + d_{4} * x_{4} + d_{5} * x_{5}
//                = 24 + 3 * 1 - 7 * 2

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q2: I_{2} = { 3, 4, 5 }

// t2(x) = x - r2

// v1_{3}(x) = c_{3} = 1
// v1_{4}(x) = c_{4} = 0
// v1_{5}(x) = c_{5} = 0
// v1_{6}(x)         = 0

// w1_{3}(x) = d_{3} = 24
// w1_{4}(x) = d_{4} = 3
// w1_{5}(x) = d_{5} = -7
// w1_{6}(x)         = 0

// y1_{3}(x)         = 0
// y1_{4}(x)         = 0
// y1_{5}(x)         = 0
// y1_{6}(x)         = 1

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                                                               *
 *        x3       x5      a3       a5     24       2    p(a3, a4, a5) * p(a3, a4, a5) - a6 = (1) * (a3 + 3 * a4 - 7 * a5) - a6  *
 *          \     /          \     /         \     /                                        = (1) * (24 + 3 * 1 - 7 * 2) - 13    *
 *           \   * 7          \   * 7         \   * 7                                       = 0                                  *
 *            \ /              \ /             \ /                                                                               *
 *    x4       -       a4       -       1       -                                                                                *
 *      \     /          \     /         \     /                                                                                 *
 *     3 *   /          3 *   /         3 *   /                                                                                  *
 *        \ /              \ /             \ /                                                                                   *
 *         +                +               +                                                                                    *
 *         |                |               |                                                                                    *
 *         * 1              * 1             * 1                                                                                  *
 *         |                |               |                                                                                    *
 *  f2(x3, x4, x5)          a6              13                                                                                   *
 *                                                                                                                               *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// Example3 generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func Example3() bool {

	return true
}
