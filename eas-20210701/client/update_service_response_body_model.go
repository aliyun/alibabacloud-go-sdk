// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceResponseBody
	GetRequestId() *string
}

type UpdateServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to update service [foo] in region [cn-shanghai]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceResponseBody) SetMessage(v string) *UpdateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
