// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceTaskResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTaskResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceTaskResponseBody
	GetRequestId() *string
	SetServiceTasks(v []map[string]interface{}) *ListServiceTaskResponseBody
	GetServiceTasks() []map[string]interface{}
	SetTotalCount(v int32) *ListServiceTaskResponseBody
	GetTotalCount() *int32
}

type ListServiceTaskResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// d23d8f3f0f0cd1984566b1986c9343122fa0385a05c09694c17fe87709f3eb56d1a7ead56b4a2536
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// [{"taskId":"a1b2c3d4-e5f6-7890-abcd-ef1234567890","taskType":"live_debug_log_probe"}]
	ServiceTasks []map[string]interface{} `json:"serviceTasks,omitempty" xml:"serviceTasks,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServiceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTaskResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTaskResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceTaskResponseBody) GetServiceTasks() []map[string]interface{} {
	return s.ServiceTasks
}

func (s *ListServiceTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceTaskResponseBody) SetMaxResults(v int32) *ListServiceTaskResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTaskResponseBody) SetNextToken(v string) *ListServiceTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceTaskResponseBody) SetRequestId(v string) *ListServiceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTaskResponseBody) SetServiceTasks(v []map[string]interface{}) *ListServiceTaskResponseBody {
	s.ServiceTasks = v
	return s
}

func (s *ListServiceTaskResponseBody) SetTotalCount(v int32) *ListServiceTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
