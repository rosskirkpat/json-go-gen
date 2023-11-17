# Input

```json
{
  "kafka": {
    "autoCreateTopics": false,
    "interBrokerProtocolVersion": "3.3",
    "kafkaVersion": "3.3.2",
    "storage": {
      "deleteClaim": false
    },
    "version": "3.3.2"
  },
  "namespace": "apps",
  "tmpDirSizeLimit": "10Mi"
}
```

# Output
```go
package main

type T struct {
	Kafka struct {
		AutoCreateTopics           bool   `json:"autoCreateTopics"`
		InterBrokerProtocolVersion string `json:"interBrokerProtocolVersion"`
		KafkaVersion               string `json:"kafkaVersion"`
		Storage                    struct {
			DeleteClaim bool `json:"deleteClaim"`
		} `json:"storage"`
		Version string `json:"version"`
	} `json:"kafka"`
	Namespace       string `json:"namespace"`
	TmpDirSizeLimit string `json:"tmpDirSizeLimit"`
}
```
