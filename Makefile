GOCMD=go  

# 目标程序名称  
Target=awsh  

# 默认目标  
all: $(Target)  

# 构建awsh  
$(Target): awsh.go  
	$(GOCMD) build -o $(Target) awsh.go  
	sudo cp $(Target) /usr/bin/  

# 清理目标  
clean:  
	sudo rm -f /usr/bin/$(Target)  
	rm -f $(Target)  

# 声明clean为伪目标  
.PHONY: clean