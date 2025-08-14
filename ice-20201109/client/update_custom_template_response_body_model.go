// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomTemplateResponseBody
	GetSuccess() *bool
}

type UpdateCustomTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomTemplateResponseBody) SetRequestId(v string) *UpdateCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomTemplateResponseBody) SetSuccess(v bool) *UpdateCustomTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
