// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdsShrink(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetAppIdsShrink() *string
	SetEntryAppId(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetEntryAppId() *string
	SetEntryAppType(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetEntryAppType() *string
	SetGroupId(v int64) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetGroupId() *int64
	SetGroupName(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetGroupName() *string
	SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetNamespaceId() *string
	SetSwimVersion(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetSwimVersion() *string
}

type CreateOrUpdateSwimmingLaneGroupShrinkRequest struct {
	AppIdsShrink *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
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

func (s CreateOrUpdateSwimmingLaneGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetAppIdsShrink() *string {
	return s.AppIdsShrink
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetEntryAppId() *string {
	return s.EntryAppId
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetEntryAppType() *string {
	return s.EntryAppType
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetSwimVersion() *string {
	return s.SwimVersion
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetAppIdsShrink(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.AppIdsShrink = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetEntryAppId(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.EntryAppId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetEntryAppType(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.EntryAppType = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetGroupName(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetSwimVersion(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.SwimVersion = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
