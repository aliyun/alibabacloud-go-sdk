// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateDBResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDBResponseBody
	GetSuccess() *bool
}

type CreateDBResponseBody struct {
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

func (s CreateDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDBResponseBody) SetMessage(v string) *CreateDBResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDBResponseBody) SetRequestId(v string) *CreateDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBResponseBody) SetSuccess(v bool) *CreateDBResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDBResponseBody) Validate() error {
	return dara.Validate(s)
}
