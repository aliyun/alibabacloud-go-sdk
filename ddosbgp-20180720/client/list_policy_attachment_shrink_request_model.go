// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyAttachmentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpPortProtocolListShrink(v string) *ListPolicyAttachmentShrinkRequest
	GetIpPortProtocolListShrink() *string
	SetPageNo(v int64) *ListPolicyAttachmentShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListPolicyAttachmentShrinkRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListPolicyAttachmentShrinkRequest
	GetPolicyId() *string
	SetPolicyType(v string) *ListPolicyAttachmentShrinkRequest
	GetPolicyType() *string
	SetPortVersion(v string) *ListPolicyAttachmentShrinkRequest
	GetPortVersion() *string
}

type ListPolicyAttachmentShrinkRequest struct {
	// The list of protected objects.
	IpPortProtocolListShrink *string `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty"`
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of rows per page in a paging query. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// f38f6520-92b7-451e-b520-9ab3********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The policy type. Valid values:
	//
	// - **default**: default mitigation policy.
	//
	// - **l3**: IP-specific mitigation policy.
	//
	// - **l4**: port-specific mitigation policy.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The version of the port-specific mitigation policy. Valid values:
	//
	// - **Not specified**: queries the policy associations that attach to the default surf DPI engine.
	//
	// - **2**: queries the policy associations that attach to the new stream DPI engine.
	//
	// example:
	//
	// 2
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s ListPolicyAttachmentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentShrinkRequest) GetIpPortProtocolListShrink() *string {
	return s.IpPortProtocolListShrink
}

func (s *ListPolicyAttachmentShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListPolicyAttachmentShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPolicyAttachmentShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPolicyAttachmentShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicyAttachmentShrinkRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ListPolicyAttachmentShrinkRequest) SetIpPortProtocolListShrink(v string) *ListPolicyAttachmentShrinkRequest {
	s.IpPortProtocolListShrink = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPageNo(v int64) *ListPolicyAttachmentShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPageSize(v int64) *ListPolicyAttachmentShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPolicyId(v string) *ListPolicyAttachmentShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPolicyType(v string) *ListPolicyAttachmentShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPortVersion(v string) *ListPolicyAttachmentShrinkRequest {
	s.PortVersion = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
