// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineCanaryEnvsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeRoutineCanaryEnvsResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeRoutineCanaryEnvsResponseBody
	GetRequestId() *string
}

type DescribeRoutineCanaryEnvsResponseBody struct {
	// The canary release environments that are supported.
	//
	// example:
	//
	// presetCanaryShanghai
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CDCD94C0-F7FE-412F-B8F8-7E3C610C78E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineCanaryEnvsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineCanaryEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCanaryEnvsResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeRoutineCanaryEnvsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoutineCanaryEnvsResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineCanaryEnvsResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineCanaryEnvsResponseBody) SetRequestId(v string) *DescribeRoutineCanaryEnvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoutineCanaryEnvsResponseBody) Validate() error {
	return dara.Validate(s)
}
