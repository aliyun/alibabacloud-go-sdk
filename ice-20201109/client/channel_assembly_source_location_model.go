// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelAssemblySourceLocation interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *ChannelAssemblySourceLocation
	GetArn() *string
	SetBaseUrl(v string) *ChannelAssemblySourceLocation
	GetBaseUrl() *string
	SetGmtCreate(v string) *ChannelAssemblySourceLocation
	GetGmtCreate() *string
	SetGmtModified(v string) *ChannelAssemblySourceLocation
	GetGmtModified() *string
	SetSegmentDeliveryConfigurations(v string) *ChannelAssemblySourceLocation
	GetSegmentDeliveryConfigurations() *string
	SetSourceLocationName(v string) *ChannelAssemblySourceLocation
	GetSourceLocationName() *string
	SetState(v int32) *ChannelAssemblySourceLocation
	GetState() *int32
}

type ChannelAssemblySourceLocation struct {
	// The ARN of the source location.
	//
	// example:
	//
	// acs:ims:mediaweaver:<regionId>:<UserId>:sourcelocation/MySourceLocation
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The base URL of the source location.
	//
	// example:
	//
	// http://xxxx.com
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The time when the source location was created.
	//
	// example:
	//
	// 2024-03-29T02:03:17Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the source location was last modified.
	//
	// example:
	//
	// 2024-03-29T02:03:17Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The segment delivery server configurations.
	//
	// example:
	//
	// http://segmentdelivery.com
	SegmentDeliveryConfigurations *string `json:"SegmentDeliveryConfigurations,omitempty" xml:"SegmentDeliveryConfigurations,omitempty"`
	// The name of the source location.
	//
	// example:
	//
	// MySourceLocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	// The status of the source location. 0: normal. 1: deleted.
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ChannelAssemblySourceLocation) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblySourceLocation) GoString() string {
	return s.String()
}

func (s *ChannelAssemblySourceLocation) GetArn() *string {
	return s.Arn
}

func (s *ChannelAssemblySourceLocation) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ChannelAssemblySourceLocation) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ChannelAssemblySourceLocation) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ChannelAssemblySourceLocation) GetSegmentDeliveryConfigurations() *string {
	return s.SegmentDeliveryConfigurations
}

func (s *ChannelAssemblySourceLocation) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ChannelAssemblySourceLocation) GetState() *int32 {
	return s.State
}

func (s *ChannelAssemblySourceLocation) SetArn(v string) *ChannelAssemblySourceLocation {
	s.Arn = &v
	return s
}

func (s *ChannelAssemblySourceLocation) SetBaseUrl(v string) *ChannelAssemblySourceLocation {
	s.BaseUrl = &v
	return s
}

func (s *ChannelAssemblySourceLocation) SetGmtCreate(v string) *ChannelAssemblySourceLocation {
	s.GmtCreate = &v
	return s
}

func (s *ChannelAssemblySourceLocation) SetGmtModified(v string) *ChannelAssemblySourceLocation {
	s.GmtModified = &v
	return s
}

func (s *ChannelAssemblySourceLocation) SetSegmentDeliveryConfigurations(v string) *ChannelAssemblySourceLocation {
	s.SegmentDeliveryConfigurations = &v
	return s
}

func (s *ChannelAssemblySourceLocation) SetSourceLocationName(v string) *ChannelAssemblySourceLocation {
	s.SourceLocationName = &v
	return s
}

func (s *ChannelAssemblySourceLocation) SetState(v int32) *ChannelAssemblySourceLocation {
	s.State = &v
	return s
}

func (s *ChannelAssemblySourceLocation) Validate() error {
	return dara.Validate(s)
}
