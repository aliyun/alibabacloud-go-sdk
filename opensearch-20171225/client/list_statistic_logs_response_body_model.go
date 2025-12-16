// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatisticLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListStatisticLogsResponseBody
	GetRequestId() *string
	SetResult(v []map[string]interface{}) *ListStatisticLogsResponseBody
	GetResult() []map[string]interface{}
	SetTotalCount(v int64) *ListStatisticLogsResponseBody
	GetTotalCount() *int64
}

type ListStatisticLogsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F76ACE3D-E510-EE2C-B7B1-39B3136A61EE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result. For more information, see
	//
	// 	- [Parameters of hotwords rankings](https://help.aliyun.com/document_detail/421248.html).
	//
	// example:
	//
	// []
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStatisticLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStatisticLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStatisticLogsResponseBody) GetResult() []map[string]interface{} {
	return s.Result
}

func (s *ListStatisticLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStatisticLogsResponseBody) SetRequestId(v string) *ListStatisticLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatisticLogsResponseBody) SetResult(v []map[string]interface{}) *ListStatisticLogsResponseBody {
	s.Result = v
	return s
}

func (s *ListStatisticLogsResponseBody) SetTotalCount(v int64) *ListStatisticLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStatisticLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
