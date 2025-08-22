// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeRoutineUserInfoResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeRoutineUserInfoResponseBody
	GetRequestId() *string
}

type DescribeRoutineUserInfoResponseBody struct {
	// The content returned by calling the operation.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 691DEEE5-4BDB-47F3-930E-F57176427717
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineUserInfoResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeRoutineUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoutineUserInfoResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineUserInfoResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineUserInfoResponseBody) SetRequestId(v string) *DescribeRoutineUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoutineUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
