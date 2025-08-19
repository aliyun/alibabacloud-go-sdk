// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustedServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminAccountId(v string) *ListTrustedServiceStatusRequest
	GetAdminAccountId() *string
	SetPageNumber(v int32) *ListTrustedServiceStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTrustedServiceStatusRequest
	GetPageSize() *int32
}

type ListTrustedServiceStatusRequest struct {
	// The ID of the management account or delegated administrator account.
	//
	// 	- If you set this parameter to the ID of a management account, the trusted services that are enabled within the management account are queried. The default value of this parameter is the ID of an management account.
	//
	// 	- If you set this parameter to the ID of a delegated administrator account, the trusted services that are enabled within the delegated administrator account are queried.
	//
	// For more information about trusted services and delegated administrator accounts, see [Overview of trusted services](https://help.aliyun.com/document_detail/208133.html) and [Delegated administrator accounts](https://help.aliyun.com/document_detail/208117.html).
	//
	// example:
	//
	// 177242285274****
	AdminAccountId *string `json:"AdminAccountId,omitempty" xml:"AdminAccountId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTrustedServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusRequest) GetAdminAccountId() *string {
	return s.AdminAccountId
}

func (s *ListTrustedServiceStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTrustedServiceStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTrustedServiceStatusRequest) SetAdminAccountId(v string) *ListTrustedServiceStatusRequest {
	s.AdminAccountId = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetPageNumber(v int32) *ListTrustedServiceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetPageSize(v int32) *ListTrustedServiceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
