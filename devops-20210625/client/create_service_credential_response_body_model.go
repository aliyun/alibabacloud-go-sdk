// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateServiceCredentialResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateServiceCredentialResponseBody
	GetErrorMessage() *string
	SetId(v int64) *CreateServiceCredentialResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateServiceCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceCredentialResponseBody
	GetSuccess() *bool
}

type CreateServiceCredentialResponseBody struct {
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
	// 11222
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

func (s CreateServiceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceCredentialResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateServiceCredentialResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateServiceCredentialResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateServiceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceCredentialResponseBody) SetErrorCode(v string) *CreateServiceCredentialResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateServiceCredentialResponseBody) SetErrorMessage(v string) *CreateServiceCredentialResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateServiceCredentialResponseBody) SetId(v int64) *CreateServiceCredentialResponseBody {
	s.Id = &v
	return s
}

func (s *CreateServiceCredentialResponseBody) SetRequestId(v string) *CreateServiceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceCredentialResponseBody) SetSuccess(v bool) *CreateServiceCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
