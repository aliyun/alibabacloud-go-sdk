// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSuperAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateSuperAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSuperAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSuperAccountResponseBody
	GetSuccess() *bool
}

type CreateSuperAccountResponseBody struct {
	// example:
	//
	// *****
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

func (s CreateSuperAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSuperAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSuperAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSuperAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSuperAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSuperAccountResponseBody) SetMessage(v string) *CreateSuperAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSuperAccountResponseBody) SetRequestId(v string) *CreateSuperAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSuperAccountResponseBody) SetSuccess(v bool) *CreateSuperAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSuperAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
