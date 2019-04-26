# kafka-topic-usage-exporter
Prometheus exporter that creates text prom file with the a metric for the size in bytes of each topic on the file system.

### example prom file
```
kafka_topic_disk_usage_bytes{kafka_log_dir="/tmp/kafka-logs", kafka_topic="test5", kafka_node="balor.local", kafka_cluster="test1"} 20975990
kafka_topic_disk_usage_bytes{kafka_log_dir="/tmp/kafka-logs", kafka_topic="test2", kafka_node="balor.local", kafka_cluster="test1"} 20976139
kafka_topic_disk_usage_bytes{kafka_log_dir="/tmp/kafka-logs", kafka_topic="test1", kafka_node="balor.local", kafka_cluster="test1"} 20976017
kafka_topic_disk_usage_bytes{kafka_log_dir="/tmp/kafka-logs", kafka_topic="test678", kafka_node="balor.local", kafka_cluster="test1"} 21427440
kafka_topic_disk_usage_bytes{kafka_log_dir="/tmp/kafka-logs", kafka_topic="test3", kafka_node="balor.local", kafka_cluster="test1"} 66002410
kafka_topic_disk_usage_bytes{kafka_log_dir="/tmp/kafka-logs", kafka_topic="test-1", kafka_node="balor.local", kafka_cluster="test1"} 21423377
```
