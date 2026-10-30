// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticSecurityPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeArchived(v bool) *ListAgenticSecurityPoliciesRequest
	GetIncludeArchived() *bool
	SetPage(v int32) *ListAgenticSecurityPoliciesRequest
	GetPage() *int32
	SetPolicyStatus(v string) *ListAgenticSecurityPoliciesRequest
	GetPolicyStatus() *string
	SetSize(v int32) *ListAgenticSecurityPoliciesRequest
	GetSize() *int32
}

type ListAgenticSecurityPoliciesRequest struct {
	// Specifies whether to include archived policies in the query scope of user policies. Valid values:
	//
	// - true: Includes archived policies.
	//
	// - false: Does not include archived policies.
	//
	// Default value: false. When querying archived policies, do not specify PolicyStatus. If ON or OFF is specified, only user policies in the corresponding state are returned. This parameter does not affect system default policies.
	IncludeArchived *bool `json:"IncludeArchived,omitempty" xml:"IncludeArchived,omitempty"`
	// The page number of the user policies to query. Pages start from 1. Default value: 1. Only user policies are paginated. System default policies are returned separately in each call.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The filter condition for the user policy status. Valid values: ON (enabled) and OFF (disabled). If this parameter is not specified or contains only whitespace characters, policies are not filtered by status. You cannot specify ARCHIVED. To query archived user policies, do not specify this parameter and set IncludeArchived to true. This parameter does not affect system default policies.
	//
	// example:
	//
	// ON
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// The number of user policies to query on each page. Valid values: 1 to 100. Default value: 20. In addition, each call returns up to Size enabled system default policies. Therefore, the number of items in the returned Data may be greater than Size.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListAgenticSecurityPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSecurityPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListAgenticSecurityPoliciesRequest) GetIncludeArchived() *bool {
	return s.IncludeArchived
}

func (s *ListAgenticSecurityPoliciesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListAgenticSecurityPoliciesRequest) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *ListAgenticSecurityPoliciesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAgenticSecurityPoliciesRequest) SetIncludeArchived(v bool) *ListAgenticSecurityPoliciesRequest {
	s.IncludeArchived = &v
	return s
}

func (s *ListAgenticSecurityPoliciesRequest) SetPage(v int32) *ListAgenticSecurityPoliciesRequest {
	s.Page = &v
	return s
}

func (s *ListAgenticSecurityPoliciesRequest) SetPolicyStatus(v string) *ListAgenticSecurityPoliciesRequest {
	s.PolicyStatus = &v
	return s
}

func (s *ListAgenticSecurityPoliciesRequest) SetSize(v int32) *ListAgenticSecurityPoliciesRequest {
	s.Size = &v
	return s
}

func (s *ListAgenticSecurityPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
