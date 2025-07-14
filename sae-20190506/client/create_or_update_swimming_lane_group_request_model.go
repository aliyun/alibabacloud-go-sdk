// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v []*string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetAppIds() []*string
	SetEntryAppId(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetEntryAppId() *string
	SetEntryAppType(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetEntryAppType() *string
	SetGroupId(v int64) *CreateOrUpdateSwimmingLaneGroupRequest
	GetGroupId() *int64
	SetGroupName(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetGroupName() *string
	SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetNamespaceId() *string
	SetSwimVersion(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetSwimVersion() *string
}

type CreateOrUpdateSwimmingLaneGroupRequest struct {
	AppIds []*string `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	// example:
	//
	// mse_ingresspost-cn-axc49******
	EntryAppId *string `json:"EntryAppId,omitempty" xml:"EntryAppId,omitempty"`
	// example:
	//
	// mse-gw
	EntryAppType *string `json:"EntryAppType,omitempty" xml:"EntryAppType,omitempty"`
	// example:
	//
	// 2047
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// mse-test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// 2
	SwimVersion *string `json:"SwimVersion,omitempty" xml:"SwimVersion,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetAppIds() []*string {
	return s.AppIds
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetEntryAppId() *string {
	return s.EntryAppId
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetEntryAppType() *string {
	return s.EntryAppType
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetSwimVersion() *string {
	return s.SwimVersion
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetAppIds(v []*string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.AppIds = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetEntryAppId(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.EntryAppId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetEntryAppType(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.EntryAppType = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetGroupName(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetSwimVersion(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.SwimVersion = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
