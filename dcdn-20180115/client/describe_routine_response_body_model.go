// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeRoutineResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeRoutineResponseBody
	GetRequestId() *string
}

type DescribeRoutineResponseBody struct {
	// The metadata of the routine. The following table describes the fields.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4DBA68F5-04A9-406B-B1E4-F2CB635E103F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeRoutineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoutineResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineResponseBody) SetRequestId(v string) *DescribeRoutineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoutineResponseBody) Validate() error {
	return dara.Validate(s)
}
