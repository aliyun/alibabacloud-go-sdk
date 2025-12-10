// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPolicyAttachmentsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPolicyAttachmentsResponseBody
	GetPageSize() *int32
	SetPolicyAttachments(v *ListPolicyAttachmentsResponseBodyPolicyAttachments) *ListPolicyAttachmentsResponseBody
	GetPolicyAttachments() *ListPolicyAttachmentsResponseBodyPolicyAttachments
	SetRequestId(v string) *ListPolicyAttachmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPolicyAttachmentsResponseBody
	GetTotalCount() *int32
}

type ListPolicyAttachmentsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the permission policies.
	PolicyAttachments *ListPolicyAttachmentsResponseBodyPolicyAttachments `json:"PolicyAttachments,omitempty" xml:"PolicyAttachments,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPolicyAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPolicyAttachmentsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPolicyAttachmentsResponseBody) GetPolicyAttachments() *ListPolicyAttachmentsResponseBodyPolicyAttachments {
	return s.PolicyAttachments
}

func (s *ListPolicyAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicyAttachmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPolicyAttachmentsResponseBody) SetPageNumber(v int32) *ListPolicyAttachmentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetPageSize(v int32) *ListPolicyAttachmentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetPolicyAttachments(v *ListPolicyAttachmentsResponseBodyPolicyAttachments) *ListPolicyAttachmentsResponseBody {
	s.PolicyAttachments = v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetRequestId(v string) *ListPolicyAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetTotalCount(v int32) *ListPolicyAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) Validate() error {
	if s.PolicyAttachments != nil {
		if err := s.PolicyAttachments.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPolicyAttachmentsResponseBodyPolicyAttachments struct {
	PolicyAttachment []*ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment `json:"PolicyAttachment,omitempty" xml:"PolicyAttachment,omitempty" type:"Repeated"`
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachments) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachments) GetPolicyAttachment() []*ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	return s.PolicyAttachment
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachments) SetPolicyAttachment(v []*ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) *ListPolicyAttachmentsResponseBodyPolicyAttachments {
	s.PolicyAttachment = v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachments) Validate() error {
	if s.PolicyAttachment != nil {
		for _, item := range s.PolicyAttachment {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment struct {
	// The time when the permission policy is attached.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The description of the permission policy.
	//
	// example:
	//
	// The description of the policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// AdministratorAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the permission policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the object to which the permission policy is attached.
	//
	// example:
	//
	// alice@demo.onaliyun.com
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the object to which the permission policy is attached. Valid values:
	//
	// 	- IMSUser: RAM user
	//
	// 	- IMSGroup: RAM user group
	//
	// 	- ServiceRole: RAM role
	//
	// example:
	//
	// IMSUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GetDescription() *string {
	return s.Description
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetAttachDate(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetDescription(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.Description = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPolicyName(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPolicyType(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPrincipalName(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PrincipalName = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPrincipalType(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PrincipalType = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetResourceGroupId(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) Validate() error {
	return dara.Validate(s)
}
