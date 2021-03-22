<p align="left">
  <img width ="200px" src="https://gitee.com/kevwan/static/raw/master/doc/images/cds/logo.png">
</p>

# ClickHouse Data Synchromesh
Data syncing in golang for ClickHouse.

围绕`ClickHouse` 的自动化数据同步组件。

galaxy控制台服务基于 [go-zero](https://github.com/tal-tech/go-zero) 构建。 

### 系统架构

下图展示了以clickhouse为存储和计算引擎的数仓架构。

![avatar](https://gitee.com/kevwan/static/raw/master/doc/images/cds/clickhouse_arch1.png)

### 数据同步设计

该部分实现了从MySQL/MongoDB数据源自动实时同步数据到ClickHouse集群的功能。

![同步drawio](https://gitee.com/kevwan/static/raw/master/doc/images/cds/同步drawio.png)

#### 快速开始

[快速开始](doc/quickstart.md)

### 数据模型
[CDS中ClickHouse使用的建表方案](doc/CDS中ClickHouse使用的建表方案.md)


欢迎大家使用和 `star` 🤝

交流群

<img src="https://gitee.com/zyz01/cds/raw/master/doc/weichat.JPG" alt="cds" width="310" />
