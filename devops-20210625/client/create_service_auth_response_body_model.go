// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateServiceAuthResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateServiceAuthResponseBody
	GetErrorMessage() *string
	SetId(v string) *CreateServiceAuthResponseBody
	GetId() *string
	SetRequestId(v string) *CreateServiceAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceAuthResponseBody
	GetSuccess() *bool
}

type CreateServiceAuthResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 1223
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateServiceAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceAuthResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateServiceAuthResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateServiceAuthResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateServiceAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceAuthResponseBody) SetErrorCode(v string) *CreateServiceAuthResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateServiceAuthResponseBody) SetErrorMessage(v string) *CreateServiceAuthResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateServiceAuthResponseBody) SetId(v string) *CreateServiceAuthResponseBody {
	s.Id = &v
	return s
}

func (s *CreateServiceAuthResponseBody) SetRequestId(v string) *CreateServiceAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceAuthResponseBody) SetSuccess(v bool) *CreateServiceAuthResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
