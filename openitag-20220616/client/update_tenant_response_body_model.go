// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateTenantResponseBody
	GetCode() *int32
	SetDetails(v string) *UpdateTenantResponseBody
	GetDetails() *string
	SetErrorCode(v string) *UpdateTenantResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateTenantResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTenantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTenantResponseBody
	GetSuccess() *bool
}

type UpdateTenantResponseBody struct {
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
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// ""
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
	// Indicates whether the request succeeded. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTenantResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateTenantResponseBody) GetDetails() *string {
	return s.Details
}

func (s *UpdateTenantResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTenantResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTenantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTenantResponseBody) SetCode(v int32) *UpdateTenantResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTenantResponseBody) SetDetails(v string) *UpdateTenantResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTenantResponseBody) SetErrorCode(v string) *UpdateTenantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTenantResponseBody) SetMessage(v string) *UpdateTenantResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTenantResponseBody) SetRequestId(v string) *UpdateTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTenantResponseBody) SetSuccess(v bool) *UpdateTenantResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTenantResponseBody) Validate() error {
	return dara.Validate(s)
}
