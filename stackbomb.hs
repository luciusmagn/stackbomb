module Main where

stackbomb n = do
	print n
	stackbomb (n+1)

main = do stackbomb 1
