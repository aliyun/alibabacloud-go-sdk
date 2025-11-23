// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbacAuthorizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAbacAuthorizationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAbacAuthorizationsRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListAbacAuthorizationsRequest
	GetPolicyId() *string
	SetPolicySource(v string) *ListAbacAuthorizationsRequest
	GetPolicySource() *string
	SetTid(v int64) *ListAbacAuthorizationsRequest
	GetTid() *int64
}

type ListAbacAuthorizationsRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 12****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The type of the policy. The value can be custom or system.
	//
	// Valid values:
	//
	// 	- USER_DEFINE
	//
	// 	- SYSTEM
	//
	// example:
	//
	// USER_DEFINE
	PolicySource *string `json:"PolicySource,omitempty" xml:"PolicySource,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAbacAuthorizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAbacAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *ListAbacAuthorizationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAbacAuthorizationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAbacAuthorizationsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListAbacAuthorizationsRequest) GetPolicySource() *string {
	return s.PolicySource
}

func (s *ListAbacAuthorizationsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAbacAuthorizationsRequest) SetPageNumber(v int64) *ListAbacAuthorizationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetPageSize(v int64) *ListAbacAuthorizationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetPolicyId(v string) *ListAbacAuthorizationsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetPolicySource(v string) *ListAbacAuthorizationsRequest {
	s.PolicySource = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetTid(v int64) *ListAbacAuthorizationsRequest {
	s.Tid = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) Validate() error {
	return dara.Validate(s)
}
