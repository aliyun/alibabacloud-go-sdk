// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceInstanceResponseBody
	GetRequestId() *string
}

type UpdateServiceInstanceResponseBody struct {
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

func (s UpdateServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceInstanceResponseBody) SetMessage(v string) *UpdateServiceInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceInstanceResponseBody) SetRequestId(v string) *UpdateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
