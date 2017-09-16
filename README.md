# flipcoins

If we are only allowed to flip coins, how do we:

* Pick a random item out of a set with `n` items?
    * Flip `k = math.ceil(log_2(n))` coins to generate bit string. Reflip if `2^k >= n`.
    * Flip `n` coins and eliminate all tails and repeat until 1 item is left. Discard and reflip if all items were to get eliminated.
