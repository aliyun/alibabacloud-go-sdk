// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineCodeRevisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeRoutineCodeRevisionResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeRoutineCodeRevisionResponseBody
	GetRequestId() *string
}

type DescribeRoutineCodeRevisionResponseBody struct {
	// The information about the JavaScript code version.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D24F0C48-1B27-4C58-8B84-1A0C001A514E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineCodeRevisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineCodeRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCodeRevisionResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeRoutineCodeRevisionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoutineCodeRevisionResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineCodeRevisionResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineCodeRevisionResponseBody) SetRequestId(v string) *DescribeRoutineCodeRevisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoutineCodeRevisionResponseBody) Validate() error {
	return dara.Validate(s)
}
