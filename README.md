# go-algorithm

## Features
- [x] Bubble Sort
- [x] Insertion Sort
- [x] Selection Sort
- [x] Shell Sort

## 概述

排序算法是《数据结构与算法》中最基本的算法之一。
排序算法可以分为内部排序和外部排序，内部排序是数据记录在内存中进行排序，而外部排序是因排序的数据很大，一次不能容纳全部的排序记录，在排序过程中需要访问外存。常见的内部排序算法有：__插入排序__、__希尔排序__、__选择排序__、__冒泡排序__、__归并排序__、__快速排序__、__堆排序__、__基数排序__ 等。用一张图概括：

![image](https://user-images.githubusercontent.com/5203608/97692300-92062100-1ada-11eb-9f92-96f923303929.png)


### 关于时间复杂度
* 平方阶 (O(n2)) 排序 各类简单排序：直接插入、直接选择和冒泡排序。
* 线性对数阶 (O(nlog2n)) 排序 快速排序、堆排序和归并排序；
* O(n1+§)) 排序，§ 是介于 0 和 1 之间的常数。 希尔排序
* 线性阶 (O(n)) 排序 基数排序，此外还有桶、箱排序。

### 关于稳定性

* 稳定的排序算法：冒泡排序、插入排序、归并排序和基数排序。

* 不是稳定的排序算法：选择排序、快速排序、希尔排序、堆排序。

### 名词解释
__n__: 数据规模

__k__: 『桶』的个数

__In-place__: 占用常数内存，不占用额外内存

__Out-place__: 占用额外内存

__稳定性__: 排序后2个相等键值的顺序和排序之前它们的顺序相同
