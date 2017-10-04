fn main()
{
	f(1);
}

fn f(n: i32)
{
	println!("{}", n);
	f(n+1);
}
