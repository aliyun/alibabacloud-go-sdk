// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResellerUserQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateResellerUserQuotaResponseBody
	GetCode() *string
	SetData(v bool) *CreateResellerUserQuotaResponseBody
	GetData() *bool
	SetMessage(v string) *CreateResellerUserQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateResellerUserQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateResellerUserQuotaResponseBody
	GetSuccess() *bool
}

type CreateResellerUserQuotaResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateResellerUserQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResellerUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResellerUserQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateResellerUserQuotaResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateResellerUserQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateResellerUserQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResellerUserQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateResellerUserQuotaResponseBody) SetCode(v string) *CreateResellerUserQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetData(v bool) *CreateResellerUserQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetMessage(v string) *CreateResellerUserQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetRequestId(v string) *CreateResellerUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetSuccess(v bool) *CreateResellerUserQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
