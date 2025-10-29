// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAllNodeResponseBody
	GetRequestId() *string
	SetResult(v []*ListAllNodeResponseBodyResult) *ListAllNodeResponseBody
	GetResult() []*ListAllNodeResponseBodyResult
}

type ListAllNodeResponseBody struct {
	// The zone ID of the node.
	//
	// example:
	//
	// 0D71B597-F3FF-5B56-88D7-74F9D3F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CPU utilization.
	//
	// >  If the **extended*	- request parameter is set to **true*	- and the monitoring information of the nodes in the cluster is being synchronized, the value of the cpuPercent parameter is null. In this case, you need to send a request again after 10 seconds to obtain the value of the cpuPercent parameter.
	Result []*ListAllNodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAllNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllNodeResponseBody) GetResult() []*ListAllNodeResponseBodyResult {
	return s.Result
}

func (s *ListAllNodeResponseBody) SetRequestId(v string) *ListAllNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllNodeResponseBody) SetResult(v []*ListAllNodeResponseBodyResult) *ListAllNodeResponseBody {
	s.Result = v
	return s
}

func (s *ListAllNodeResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllNodeResponseBodyResult struct {
	// The disk usage.
	//
	// example:
	//
	// 4.2%
	CpuPercent *string `json:"cpuPercent,omitempty" xml:"cpuPercent,omitempty"`
	// The health status of the node. Valid values: GREEN, YELLOW, RED, and GRAY.
	//
	// example:
	//
	// 1.0%
	DiskUsedPercent *string `json:"diskUsedPercent,omitempty" xml:"diskUsedPercent,omitempty"`
	// example:
	//
	// GREEN
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 21.6%
	HeapPercent *string `json:"heapPercent,omitempty" xml:"heapPercent,omitempty"`
	// The port that is used to connect to the node.
	//
	// example:
	//
	// 10.15.XX.XX
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// 0.12
	LoadOneM *string `json:"loadOneM,omitempty" xml:"loadOneM,omitempty"`
	// The 1-minute load of the node.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// example:
	//
	// 9200
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The type of the nodes. Valid values:
	//
	// 	- MASTER: dedicated master node
	//
	// 	- WORKER: hot node
	//
	// 	- WORKER_WARM: warm node
	//
	// 	- COORDINATING: client node
	//
	// 	- KIBANA: Kibana node
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListAllNodeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAllNodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAllNodeResponseBodyResult) GetCpuPercent() *string {
	return s.CpuPercent
}

func (s *ListAllNodeResponseBodyResult) GetDiskUsedPercent() *string {
	return s.DiskUsedPercent
}

func (s *ListAllNodeResponseBodyResult) GetHealth() *string {
	return s.Health
}

func (s *ListAllNodeResponseBodyResult) GetHeapPercent() *string {
	return s.HeapPercent
}

func (s *ListAllNodeResponseBodyResult) GetHost() *string {
	return s.Host
}

func (s *ListAllNodeResponseBodyResult) GetLoadOneM() *string {
	return s.LoadOneM
}

func (s *ListAllNodeResponseBodyResult) GetNodeType() *string {
	return s.NodeType
}

func (s *ListAllNodeResponseBodyResult) GetPort() *int32 {
	return s.Port
}

func (s *ListAllNodeResponseBodyResult) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListAllNodeResponseBodyResult) SetCpuPercent(v string) *ListAllNodeResponseBodyResult {
	s.CpuPercent = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetDiskUsedPercent(v string) *ListAllNodeResponseBodyResult {
	s.DiskUsedPercent = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetHealth(v string) *ListAllNodeResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetHeapPercent(v string) *ListAllNodeResponseBodyResult {
	s.HeapPercent = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetHost(v string) *ListAllNodeResponseBodyResult {
	s.Host = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetLoadOneM(v string) *ListAllNodeResponseBodyResult {
	s.LoadOneM = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetNodeType(v string) *ListAllNodeResponseBodyResult {
	s.NodeType = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetPort(v int32) *ListAllNodeResponseBodyResult {
	s.Port = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetZoneId(v string) *ListAllNodeResponseBodyResult {
	s.ZoneId = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
