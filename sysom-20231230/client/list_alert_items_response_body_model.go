// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertItemsResponseBody
	GetCode() *string
	SetData(v interface{}) *ListAlertItemsResponseBody
	GetData() interface{}
	SetMessage(v string) *ListAlertItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlertItemsResponseBody
	GetRequestId() *string
}

type ListAlertItemsResponseBody struct {
	// The status code.
	//
	// - If `code == Success`, the authorization is successful.
	//
	// - Other status codes indicate authorization failure. If authorization fails, check the `message` field for detailed error messages.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {
	//
	//     "NODE": {
	//
	//       "Saturation": [
	//
	//         "Node CPU utilization detection",
	//
	//         "Node kernel-mode CPU utilization detection",
	//
	//         "Node softirq CPU utilization detection",
	//
	//         "Node memory usage detection",
	//
	//         "Node kernel memory usage detection",
	//
	//         "Node file descriptor usage detection",
	//
	//         "Node root file system usage detection",
	//
	//         "Node cgroup leak detection",
	//
	//         "Node socket usage detection",
	//
	//         "Node TCP memory usage detection"
	//
	//       ],
	//
	//       "Latency": [
	//
	//         "Node scheduling latency detection",
	//
	//         "Node network latency detection",
	//
	//         "Node disk write latency detection",
	//
	//         "Node disk read latency detection"
	//
	//       ],
	//
	//       "Load": [
	//
	//         "Node disk I/O traffic detection",
	//
	//         "Node load average detection"
	//
	//       ],
	//
	//       "Errors": [
	//
	//         "Node network packet loss detection",
	//
	//         "Node OOM stall prediction and detection"
	//
	//       ]
	//
	//     },
	//
	//     "POD": {
	//
	//       "Saturation": [
	//
	//         "Pod memory usage detection",
	//
	//         "Pod CPU utilization detection"
	//
	//       ],
	//
	//       "Errors": [
	//
	//         "Pod CPU throttling detection",
	//
	//         "Pod OOM event detection"
	//
	//       ]
	//
	//     }
	//
	// }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAlertItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertItemsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListAlertItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertItemsResponseBody) SetCode(v string) *ListAlertItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertItemsResponseBody) SetData(v interface{}) *ListAlertItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertItemsResponseBody) SetMessage(v string) *ListAlertItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertItemsResponseBody) SetRequestId(v string) *ListAlertItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
