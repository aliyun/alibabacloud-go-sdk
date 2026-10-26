// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppServiceAggregate interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *AppServiceAggregate
	GetBizId() *string
	SetDeleted(v int32) *AppServiceAggregate
	GetDeleted() *int32
	SetEndTime(v string) *AppServiceAggregate
	GetEndTime() *string
	SetEspBizId(v string) *AppServiceAggregate
	GetEspBizId() *string
	SetGmtCreate(v string) *AppServiceAggregate
	GetGmtCreate() *string
	SetGmtModified(v string) *AppServiceAggregate
	GetGmtModified() *string
	SetGroup(v *AppServiceGroup) *AppServiceAggregate
	GetGroup() *AppServiceGroup
	SetInstanceBizId(v string) *AppServiceAggregate
	GetInstanceBizId() *string
	SetName(v string) *AppServiceAggregate
	GetName() *string
	SetNodeList(v []*TreeNode) *AppServiceAggregate
	GetNodeList() []*TreeNode
	SetOperationAddress(v *AppOperationAddress) *AppServiceAggregate
	GetOperationAddress() *AppOperationAddress
	SetProfile(v *AppServiceProfile) *AppServiceAggregate
	GetProfile() *AppServiceProfile
	SetServiceType(v string) *AppServiceAggregate
	GetServiceType() *string
	SetServiceTypeText(v string) *AppServiceAggregate
	GetServiceTypeText() *string
	SetSlug(v string) *AppServiceAggregate
	GetSlug() *string
	SetStartTime(v string) *AppServiceAggregate
	GetStartTime() *string
	SetStatus(v string) *AppServiceAggregate
	GetStatus() *string
	SetUserId(v string) *AppServiceAggregate
	GetUserId() *string
}

type AppServiceAggregate struct {
	// The business ID of the application instance.
	//
	// example:
	//
	// WS20251211153330000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The deletion flag.
	//
	// example:
	//
	// 0
	Deleted *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The end time. It is in ISO 8601 format and uses the UTC time zone (Z).
	//
	// example:
	//
	// 2025-07-18T02:23:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// esp bizId
	//
	// example:
	//
	// p20202933455
	EspBizId *string `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-07-30T02:08:40Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// Tue Sep 09 10:27:49 CST 2025
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The operation address of the application.
	Group *AppServiceGroup `json:"Group,omitempty" xml:"Group,omitempty"`
	// The associated business ID.
	//
	// example:
	//
	// WS20250915163734000001
	InstanceBizId *string `json:"InstanceBizId,omitempty" xml:"InstanceBizId,omitempty"`
	// The service name.
	//
	// example:
	//
	// Website design service
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of flow nodes.
	NodeList []*TreeNode `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
	// Uniquely identifies an operation.
	OperationAddress *AppOperationAddress `json:"OperationAddress,omitempty" xml:"OperationAddress,omitempty"`
	// The configuration information of the application service.
	Profile *AppServiceProfile `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The service type.
	//
	// example:
	//
	// WEBSITE_DESIGN
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The description of the service type.
	ServiceTypeText *string `json:"ServiceTypeText,omitempty" xml:"ServiceTypeText,omitempty"`
	// The application identifier.
	//
	// example:
	//
	// inst-001
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The start time. It is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1754447102000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123241414
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AppServiceAggregate) String() string {
	return dara.Prettify(s)
}

func (s AppServiceAggregate) GoString() string {
	return s.String()
}

func (s *AppServiceAggregate) GetBizId() *string {
	return s.BizId
}

func (s *AppServiceAggregate) GetDeleted() *int32 {
	return s.Deleted
}

func (s *AppServiceAggregate) GetEndTime() *string {
	return s.EndTime
}

func (s *AppServiceAggregate) GetEspBizId() *string {
	return s.EspBizId
}

func (s *AppServiceAggregate) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AppServiceAggregate) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AppServiceAggregate) GetGroup() *AppServiceGroup {
	return s.Group
}

func (s *AppServiceAggregate) GetInstanceBizId() *string {
	return s.InstanceBizId
}

func (s *AppServiceAggregate) GetName() *string {
	return s.Name
}

func (s *AppServiceAggregate) GetNodeList() []*TreeNode {
	return s.NodeList
}

func (s *AppServiceAggregate) GetOperationAddress() *AppOperationAddress {
	return s.OperationAddress
}

func (s *AppServiceAggregate) GetProfile() *AppServiceProfile {
	return s.Profile
}

func (s *AppServiceAggregate) GetServiceType() *string {
	return s.ServiceType
}

func (s *AppServiceAggregate) GetServiceTypeText() *string {
	return s.ServiceTypeText
}

func (s *AppServiceAggregate) GetSlug() *string {
	return s.Slug
}

func (s *AppServiceAggregate) GetStartTime() *string {
	return s.StartTime
}

func (s *AppServiceAggregate) GetStatus() *string {
	return s.Status
}

func (s *AppServiceAggregate) GetUserId() *string {
	return s.UserId
}

func (s *AppServiceAggregate) SetBizId(v string) *AppServiceAggregate {
	s.BizId = &v
	return s
}

func (s *AppServiceAggregate) SetDeleted(v int32) *AppServiceAggregate {
	s.Deleted = &v
	return s
}

func (s *AppServiceAggregate) SetEndTime(v string) *AppServiceAggregate {
	s.EndTime = &v
	return s
}

func (s *AppServiceAggregate) SetEspBizId(v string) *AppServiceAggregate {
	s.EspBizId = &v
	return s
}

func (s *AppServiceAggregate) SetGmtCreate(v string) *AppServiceAggregate {
	s.GmtCreate = &v
	return s
}

func (s *AppServiceAggregate) SetGmtModified(v string) *AppServiceAggregate {
	s.GmtModified = &v
	return s
}

func (s *AppServiceAggregate) SetGroup(v *AppServiceGroup) *AppServiceAggregate {
	s.Group = v
	return s
}

func (s *AppServiceAggregate) SetInstanceBizId(v string) *AppServiceAggregate {
	s.InstanceBizId = &v
	return s
}

func (s *AppServiceAggregate) SetName(v string) *AppServiceAggregate {
	s.Name = &v
	return s
}

func (s *AppServiceAggregate) SetNodeList(v []*TreeNode) *AppServiceAggregate {
	s.NodeList = v
	return s
}

func (s *AppServiceAggregate) SetOperationAddress(v *AppOperationAddress) *AppServiceAggregate {
	s.OperationAddress = v
	return s
}

func (s *AppServiceAggregate) SetProfile(v *AppServiceProfile) *AppServiceAggregate {
	s.Profile = v
	return s
}

func (s *AppServiceAggregate) SetServiceType(v string) *AppServiceAggregate {
	s.ServiceType = &v
	return s
}

func (s *AppServiceAggregate) SetServiceTypeText(v string) *AppServiceAggregate {
	s.ServiceTypeText = &v
	return s
}

func (s *AppServiceAggregate) SetSlug(v string) *AppServiceAggregate {
	s.Slug = &v
	return s
}

func (s *AppServiceAggregate) SetStartTime(v string) *AppServiceAggregate {
	s.StartTime = &v
	return s
}

func (s *AppServiceAggregate) SetStatus(v string) *AppServiceAggregate {
	s.Status = &v
	return s
}

func (s *AppServiceAggregate) SetUserId(v string) *AppServiceAggregate {
	s.UserId = &v
	return s
}

func (s *AppServiceAggregate) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	if s.NodeList != nil {
		for _, item := range s.NodeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OperationAddress != nil {
		if err := s.OperationAddress.Validate(); err != nil {
			return err
		}
	}
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
