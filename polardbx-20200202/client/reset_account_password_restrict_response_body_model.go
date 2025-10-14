// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRestrictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ResetAccountPasswordRestrictResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetAccountPasswordRestrictResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetAccountPasswordRestrictResponseBody
	GetSuccess() *bool
}

type ResetAccountPasswordRestrictResponseBody struct {
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-****-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetAccountPasswordRestrictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRestrictResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRestrictResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetAccountPasswordRestrictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAccountPasswordRestrictResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetAccountPasswordRestrictResponseBody) SetMessage(v string) *ResetAccountPasswordRestrictResponseBody {
	s.Message = &v
	return s
}

func (s *ResetAccountPasswordRestrictResponseBody) SetRequestId(v string) *ResetAccountPasswordRestrictResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordRestrictResponseBody) SetSuccess(v bool) *ResetAccountPasswordRestrictResponseBody {
	s.Success = &v
	return s
}

func (s *ResetAccountPasswordRestrictResponseBody) Validate() error {
	return dara.Validate(s)
}
