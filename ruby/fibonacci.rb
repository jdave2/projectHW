first_=0
second=1
nxtTerm=0

puts "Enter the number of terms:-"
n=gets.chomp.to_i

puts "The first #{n} terms of Fibonacci series are:-"
c=1
while(c<=n+1)
	if(c<=1)
		nxtTerm=c
	else
		puts nxtTerm
		nxtTerm=first_+second
		first_=second
		second=nxtTerm
	end
	c+=1
end
