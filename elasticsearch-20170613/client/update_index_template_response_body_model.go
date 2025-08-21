// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIndexTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIndexTemplateResponseBody
	GetRequestId() *string
	SetResult(v string) *UpdateIndexTemplateResponseBody
	GetResult() *string
}

type UpdateIndexTemplateResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// my-template
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateIndexTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIndexTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIndexTemplateResponseBody) GetResult() *string {
	return s.Result
}

func (s *UpdateIndexTemplateResponseBody) SetRequestId(v string) *UpdateIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIndexTemplateResponseBody) SetResult(v string) *UpdateIndexTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateIndexTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
