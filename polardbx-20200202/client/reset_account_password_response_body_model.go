// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ResetAccountPasswordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetAccountPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetAccountPasswordResponseBody
	GetSuccess() *bool
}

type ResetAccountPasswordResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73559800-3c8c-11ec-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAccountPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetAccountPasswordResponseBody) SetMessage(v string) *ResetAccountPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetSuccess(v bool) *ResetAccountPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
