// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomTemplateResponseBody
	GetSuccess() *bool
}

type DeleteCustomTemplateResponseBody struct {
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

func (s DeleteCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomTemplateResponseBody) SetRequestId(v string) *DeleteCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTemplateResponseBody) SetSuccess(v bool) *DeleteCustomTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
