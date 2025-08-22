// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeRoutineSpecResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeRoutineSpecResponseBody
	GetRequestId() *string
}

type DescribeRoutineSpecResponseBody struct {
	// The specification of the CPU time slice. Valid values: 5 ms, 50 ms, and 100 ms.
	//
	// example:
	//
	// 5ms, 50ms, 100ms
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AE4E1B80-D5F3-47DB-824A-DA98A21854C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineSpecResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeRoutineSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoutineSpecResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineSpecResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineSpecResponseBody) SetRequestId(v string) *DescribeRoutineSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoutineSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
