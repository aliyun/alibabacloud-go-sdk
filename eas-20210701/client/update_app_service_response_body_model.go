// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateAppServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAppServiceResponseBody
	GetRequestId() *string
}

type UpdateAppServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAppServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppServiceResponseBody) SetMessage(v string) *UpdateAppServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppServiceResponseBody) SetRequestId(v string) *UpdateAppServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
