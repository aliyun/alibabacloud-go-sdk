// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateServiceConnectionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateServiceConnectionResponseBody
	GetErrorMessage() *string
	SetId(v int64) *CreateServiceConnectionResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateServiceConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceConnectionResponseBody
	GetSuccess() *bool
}

type CreateServiceConnectionResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 19224
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateServiceConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateServiceConnectionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateServiceConnectionResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateServiceConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceConnectionResponseBody) SetErrorCode(v string) *CreateServiceConnectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetErrorMessage(v string) *CreateServiceConnectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetId(v int64) *CreateServiceConnectionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetRequestId(v string) *CreateServiceConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetSuccess(v bool) *CreateServiceConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
