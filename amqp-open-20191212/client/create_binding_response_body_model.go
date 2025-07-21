// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateBindingResponseBody
	GetRequestId() *string
}

type CreateBindingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 09768C14-E51C-4F4A-8077-30810032C***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBindingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBindingResponseBody) SetRequestId(v string) *CreateBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
