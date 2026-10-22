1. wget https://github.com/protocolbuffers/protobuf/releases/download/v23.0/protoc-23. 0-lìnux-x86_64.zip

2. unzip protoc-23.0-linux-x86 64.zip -d protoc

3. sudo mv protoc/bin/protoc /usr/local/bin/

4. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

5. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

6. echo $PATH

7.

8. source ~/.bashrc

9. create the pb folder in your project

10. add this to account.proto option go_package = "./pb";

11. finally run this command - protoc --go out=./pb --go-grpc_out=./pb account.proto