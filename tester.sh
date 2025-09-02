
clear
echo "Correct results:"
cd stat-bin
./run.sh math-skills  
cd ..
echo "\nUser results:"
go run . stat-bin/data.txt