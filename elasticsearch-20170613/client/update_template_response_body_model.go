// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTemplateResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateTemplateResponseBody
	GetResult() *bool
}

type UpdateTemplateResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetResult(v bool) *UpdateTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
