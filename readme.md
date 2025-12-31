# ChannelUsage
## channel VS sharing slice （one source one final slice）
Using channel better

Execution time with channel:2.625µs
Execution time with sharing variable:4.375µs
## channel VS sharing slice （channel: multiple sources one final slice;slice: one source one final slice ）
Sharing Slice better

Execution time with channel:94.875µs
Execution time with channel:135.875µs
Execution time with channel:61.916µs
Execution time with channel:93.959µs
Execution time with channel:68.542µs

Execution time with sharing variable:25.583µs
Execution time with sharing variable:6.125µs
Execution time with sharing variable:4.209µs
Execution time with sharing variable:4.083µs
Execution time with sharing variable:5.875µs

Before Wait: Execution time with channel for wating: 5.5µs
After Wait: Execution time with channel for merging: 98.917µs
Before Merge Result: Execution time with channel for merging: 141.917µs
After Merge Result: Execution time with channel:145.208µs

Spend A lot of time for waiting and merging
## channel VS sharing slice （channel: multiple sources one final slice;slice multiple sources one final slice）
Sharing Slice better

Execution time with sharing variable:36.75µs
Execution time with sharing variable:17.084µs