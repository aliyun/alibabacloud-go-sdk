// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTenantResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTenantResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTenantResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTenantResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTenantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTenantResponseBody
	GetSuccess() *bool
	SetTenant(v *SingleTenant) *GetTenantResponseBody
	GetTenant() *SingleTenant
}

type GetTenantResponseBody struct {
	// Return encoding. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// -
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// -
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Possible values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Tenant information.
	Tenant *SingleTenant `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
}

func (s GetTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTenantResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTenantResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTenantResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTenantResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTenantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTenantResponseBody) GetTenant() *SingleTenant {
	return s.Tenant
}

func (s *GetTenantResponseBody) SetCode(v int32) *GetTenantResponseBody {
	s.Code = &v
	return s
}

func (s *GetTenantResponseBody) SetDetails(v string) *GetTenantResponseBody {
	s.Details = &v
	return s
}

func (s *GetTenantResponseBody) SetErrorCode(v string) *GetTenantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTenantResponseBody) SetMessage(v string) *GetTenantResponseBody {
	s.Message = &v
	return s
}

func (s *GetTenantResponseBody) SetRequestId(v string) *GetTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTenantResponseBody) SetSuccess(v bool) *GetTenantResponseBody {
	s.Success = &v
	return s
}

func (s *GetTenantResponseBody) SetTenant(v *SingleTenant) *GetTenantResponseBody {
	s.Tenant = v
	return s
}

func (s *GetTenantResponseBody) Validate() error {
	if s.Tenant != nil {
		if err := s.Tenant.Validate(); err != nil {
			return err
		}
	}
	return nil
}
