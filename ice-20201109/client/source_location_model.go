// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceLocation interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *SourceLocation
	GetArn() *string
	SetBaseUrl(v string) *SourceLocation
	GetBaseUrl() *string
	SetGmtCreate(v string) *SourceLocation
	GetGmtCreate() *string
	SetGmtModified(v string) *SourceLocation
	GetGmtModified() *string
	SetSegmentDeliveryConfigurations(v string) *SourceLocation
	GetSegmentDeliveryConfigurations() *string
	SetSourceLocationName(v string) *SourceLocation
	GetSourceLocationName() *string
	SetState(v int32) *SourceLocation
	GetState() *int32
}

type SourceLocation struct {
	Arn                           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	BaseUrl                       *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	GmtCreate                     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified                   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SegmentDeliveryConfigurations *string `json:"SegmentDeliveryConfigurations,omitempty" xml:"SegmentDeliveryConfigurations,omitempty"`
	SourceLocationName            *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	State                         *int32  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SourceLocation) String() string {
	return dara.Prettify(s)
}

func (s SourceLocation) GoString() string {
	return s.String()
}

func (s *SourceLocation) GetArn() *string {
	return s.Arn
}

func (s *SourceLocation) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *SourceLocation) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *SourceLocation) GetGmtModified() *string {
	return s.GmtModified
}

func (s *SourceLocation) GetSegmentDeliveryConfigurations() *string {
	return s.SegmentDeliveryConfigurations
}

func (s *SourceLocation) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *SourceLocation) GetState() *int32 {
	return s.State
}

func (s *SourceLocation) SetArn(v string) *SourceLocation {
	s.Arn = &v
	return s
}

func (s *SourceLocation) SetBaseUrl(v string) *SourceLocation {
	s.BaseUrl = &v
	return s
}

func (s *SourceLocation) SetGmtCreate(v string) *SourceLocation {
	s.GmtCreate = &v
	return s
}

func (s *SourceLocation) SetGmtModified(v string) *SourceLocation {
	s.GmtModified = &v
	return s
}

func (s *SourceLocation) SetSegmentDeliveryConfigurations(v string) *SourceLocation {
	s.SegmentDeliveryConfigurations = &v
	return s
}

func (s *SourceLocation) SetSourceLocationName(v string) *SourceLocation {
	s.SourceLocationName = &v
	return s
}

func (s *SourceLocation) SetState(v int32) *SourceLocation {
	s.State = &v
	return s
}

func (s *SourceLocation) Validate() error {
	return dara.Validate(s)
}
