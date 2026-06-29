// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListTenantsResponseBody
	GetCode() *int32
	SetDetails(v string) *ListTenantsResponseBody
	GetDetails() *string
	SetErrorCode(v string) *ListTenantsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListTenantsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTenantsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTenantsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTenantsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTenantsResponseBody
	GetSuccess() *bool
	SetTenants(v []*SimpleTenant) *ListTenantsResponseBody
	GetTenants() []*SimpleTenant
	SetTotalCount(v int32) *ListTenantsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListTenantsResponseBody
	GetTotalPage() *int32
}

type ListTenantsResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
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
	// Returned error code.
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
	// Page number of the returned tenant list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of tenants displayed per page in the response.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// List of tenants.
	Tenants []*SimpleTenant `json:"Tenants,omitempty" xml:"Tenants,omitempty" type:"Repeated"`
	// Total number of tenants.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages in the tenant list.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTenantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTenantsResponseBody) GetDetails() *string {
	return s.Details
}

func (s *ListTenantsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTenantsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTenantsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTenantsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTenantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTenantsResponseBody) GetTenants() []*SimpleTenant {
	return s.Tenants
}

func (s *ListTenantsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTenantsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListTenantsResponseBody) SetCode(v int32) *ListTenantsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantsResponseBody) SetDetails(v string) *ListTenantsResponseBody {
	s.Details = &v
	return s
}

func (s *ListTenantsResponseBody) SetErrorCode(v string) *ListTenantsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTenantsResponseBody) SetMessage(v string) *ListTenantsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantsResponseBody) SetPageNumber(v int32) *ListTenantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTenantsResponseBody) SetPageSize(v int32) *ListTenantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTenantsResponseBody) SetRequestId(v string) *ListTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantsResponseBody) SetSuccess(v bool) *ListTenantsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTenantsResponseBody) SetTenants(v []*SimpleTenant) *ListTenantsResponseBody {
	s.Tenants = v
	return s
}

func (s *ListTenantsResponseBody) SetTotalCount(v int32) *ListTenantsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTenantsResponseBody) SetTotalPage(v int32) *ListTenantsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListTenantsResponseBody) Validate() error {
	if s.Tenants != nil {
		for _, item := range s.Tenants {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
