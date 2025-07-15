// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCustomTemplateResponseBody
	GetRequestId() *string
}

type CreateCustomTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0D715397-2E66-4AE1-694h-C546628AD145
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomTemplateResponseBody) SetRequestId(v string) *CreateCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
