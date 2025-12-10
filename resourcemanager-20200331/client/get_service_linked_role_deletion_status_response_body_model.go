// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleDeletionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v *GetServiceLinkedRoleDeletionStatusResponseBodyReason) *GetServiceLinkedRoleDeletionStatusResponseBody
	GetReason() *GetServiceLinkedRoleDeletionStatusResponseBodyReason
	SetRequestId(v string) *GetServiceLinkedRoleDeletionStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetServiceLinkedRoleDeletionStatusResponseBody
	GetStatus() *string
}

type GetServiceLinkedRoleDeletionStatusResponseBody struct {
	// The cause for the failure of the deletion task.
	Reason *GetServiceLinkedRoleDeletionStatusResponseBodyReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 07194EB1-DB50-4513-A51D-99B30D635AEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the task.
	//
	// 	- SUCCEEDED
	//
	// 	- IN_PROGRESS
	//
	// 	- FAILED
	//
	// 	- NOT_STARTED
	//
	// 	- INTERNAL_ERROR
	//
	// example:
	//
	// FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) GetReason() *GetServiceLinkedRoleDeletionStatusResponseBodyReason {
	return s.Reason
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) SetReason(v *GetServiceLinkedRoleDeletionStatusResponseBodyReason) *GetServiceLinkedRoleDeletionStatusResponseBody {
	s.Reason = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) SetRequestId(v string) *GetServiceLinkedRoleDeletionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) SetStatus(v string) *GetServiceLinkedRoleDeletionStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) Validate() error {
	if s.Reason != nil {
		if err := s.Reason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReason struct {
	// The failure information.
	//
	// example:
	//
	// Service-Linked Role acs:ram::196813227629****:role/aliyunserviceroleforhdr cannot be deleted as it is in use by hdr.aliyuncs.com.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about the resources that the service-linked role can use.
	RoleUsages *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" type:"Struct"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReason) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReason) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReason) GetMessage() *string {
	return s.Message
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReason) GetRoleUsages() *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages {
	return s.RoleUsages
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReason) SetMessage(v string) *GetServiceLinkedRoleDeletionStatusResponseBodyReason {
	s.Message = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReason) SetRoleUsages(v *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) *GetServiceLinkedRoleDeletionStatusResponseBodyReason {
	s.RoleUsages = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReason) Validate() error {
	if s.RoleUsages != nil {
		if err := s.RoleUsages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages struct {
	RoleUsage []*GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage `json:"RoleUsage,omitempty" xml:"RoleUsage,omitempty" type:"Repeated"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) GetRoleUsage() []*GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage {
	return s.RoleUsage
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) SetRoleUsage(v []*GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages {
	s.RoleUsage = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) Validate() error {
	if s.RoleUsage != nil {
		for _, item := range s.RoleUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage struct {
	// The region.
	//
	// example:
	//
	// global
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The information about resources.
	Resources *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) GetRegion() *string {
	return s.Region
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) GetResources() *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources {
	return s.Resources
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) SetRegion(v string) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage {
	s.Region = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) SetResources(v *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage {
	s.Resources = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources struct {
	Resource []*string `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) GetResource() []*string {
	return s.Resource
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) SetResource(v []*string) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources {
	s.Resource = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) Validate() error {
	return dara.Validate(s)
}
