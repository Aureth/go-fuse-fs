echo "-------  Start Creating the File System  -------"
dd if=fs of=file_system.txt bs=100m count=1
go run
echo "-------  Done  -------"
