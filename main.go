package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := "{\"s3Bucket\":\"otherword-bigcat\",\"s3ObjectKey\":[\"AWSLogs/987654321/CloudTrail/us-west-2/2023/05/27/987654321_CloudTrail_us-west-2_20230527T2325Z_4Ef2boJWDeex6r8s.json.gz\"]}"

	var data map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
