// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAccountResponseBody
	GetSuccess() *bool
}

type DeleteAccountResponseBody struct {
	// example:
	//
	// ****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAccountResponseBody) SetMessage(v string) *DeleteAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) SetSuccess(v bool) *DeleteAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
