// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListSearchLogResponseBodyHeaders) *ListSearchLogResponseBody
	GetHeaders() *ListSearchLogResponseBodyHeaders
	SetRequestId(v string) *ListSearchLogResponseBody
	GetRequestId() *string
	SetResult(v []*ListSearchLogResponseBodyResult) *ListSearchLogResponseBody
	GetResult() []*ListSearchLogResponseBodyResult
}

type ListSearchLogResponseBody struct {
	// The level of the log. Valid values:
	//
	// 	- warn: warning log
	//
	// 	- info: information log
	//
	// 	- error: error log
	//
	// 	- trace: trace logs
	//
	// 	- debug: debug logs
	//
	// The level information has been migrated to the contentCollection field.
	Headers *ListSearchLogResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The list of logs returned by the request.
	//
	// example:
	//
	// 7F40EAA1-6F1D-4DD9-8DB8-C5F00C4E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The content of the log entry. Migrated to the contentCollection field.
	Result []*ListSearchLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSearchLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponseBody) GetHeaders() *ListSearchLogResponseBodyHeaders {
	return s.Headers
}

func (s *ListSearchLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchLogResponseBody) GetResult() []*ListSearchLogResponseBodyResult {
	return s.Result
}

func (s *ListSearchLogResponseBody) SetHeaders(v *ListSearchLogResponseBodyHeaders) *ListSearchLogResponseBody {
	s.Headers = v
	return s
}

func (s *ListSearchLogResponseBody) SetRequestId(v string) *ListSearchLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchLogResponseBody) SetResult(v []*ListSearchLogResponseBodyResult) *ListSearchLogResponseBody {
	s.Result = v
	return s
}

func (s *ListSearchLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSearchLogResponseBodyHeaders struct {
	// The IP address of the node that generates the log.
	//
	// example:
	//
	// 1000
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListSearchLogResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLogResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListSearchLogResponseBodyHeaders) SetXTotalCount(v int32) *ListSearchLogResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListSearchLogResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListSearchLogResponseBodyResult struct {
	// The ID of the instance.
	//
	// example:
	//
	// [GC (Allocation Failure) 2018-07-19T17:24:20.682+0800: 7516.513: [ParNew: 6604768K->81121K(7341504K), 0.0760606 secs] 7226662K->703015K(31813056K), 0.0762507 secs] [Times: user=0.52 sys=0.00, real=0.07 secs]
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// {"level": "info", "host": "``192.168.**.**``", "time": "2019-03-18T08:16:12.741Z","content": "[o.e.c.r.a.AllocationService] [MnNASM_] Cluster health status changed from [YELLOW] to [GREEN] (reason: [shards started [[my_index][3]] ...])."}
	ContentCollection map[string]interface{} `json:"contentCollection,omitempty" xml:"contentCollection,omitempty"`
	// Details of the log entry. Different content fields are returned for different log types.
	//
	// example:
	//
	// ``192.168.**.**``
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The timestamp when the log is generated. Unit: ms.
	//
	// example:
	//
	// info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// 1531985112420
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListSearchLogResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ListSearchLogResponseBodyResult) GetContentCollection() map[string]interface{} {
	return s.ContentCollection
}

func (s *ListSearchLogResponseBodyResult) GetHost() *string {
	return s.Host
}

func (s *ListSearchLogResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSearchLogResponseBodyResult) GetLevel() *string {
	return s.Level
}

func (s *ListSearchLogResponseBodyResult) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListSearchLogResponseBodyResult) SetContent(v string) *ListSearchLogResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetContentCollection(v map[string]interface{}) *ListSearchLogResponseBodyResult {
	s.ContentCollection = v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetHost(v string) *ListSearchLogResponseBodyResult {
	s.Host = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetInstanceId(v string) *ListSearchLogResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetLevel(v string) *ListSearchLogResponseBodyResult {
	s.Level = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetTimestamp(v int64) *ListSearchLogResponseBodyResult {
	s.Timestamp = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
