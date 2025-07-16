// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceSafetyLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceSafetyLockResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceSafetyLockResponseBody
	GetRequestId() *string
}

type UpdateServiceSafetyLockResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// service safety lock updated to dangerous
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E089D584-B6F4-50C4-9902-DA2295B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceSafetyLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSafetyLockResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceSafetyLockResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceSafetyLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceSafetyLockResponseBody) SetMessage(v string) *UpdateServiceSafetyLockResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceSafetyLockResponseBody) SetRequestId(v string) *UpdateServiceSafetyLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceSafetyLockResponseBody) Validate() error {
	return dara.Validate(s)
}
