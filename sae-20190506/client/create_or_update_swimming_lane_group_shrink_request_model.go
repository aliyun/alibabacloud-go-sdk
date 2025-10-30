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
	// The ID of the baseline application.
	AppIdsShrink *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The unique ID of the corresponding gateway.
	//
	// example:
	//
	// mse_ingresspost-cn-axc49******
	EntryAppId *string `json:"EntryAppId,omitempty" xml:"EntryAppId,omitempty"`
	// The application entry type (gateway type).
	//
	// 	- **apig:*	- cloud-native API Gateway
	//
	// 	- **mse:*	- Java Services Gateway
	//
	// 	- **mse-gw:*	- MSE cloud-native Gateway
	//
	// example:
	//
	// mse-gw
	EntryAppType *string `json:"EntryAppType,omitempty" xml:"EntryAppType,omitempty"`
	// The ID of the lane group. This is required when you update a lane group.
	//
	// example:
	//
	// 2047
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// mse-test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of a namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The end-to-end grayscale version. Valid values: 0 and 2 (recommended).
	//
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
