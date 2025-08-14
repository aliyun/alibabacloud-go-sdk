// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelAssemblySource interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *ChannelAssemblySource
	GetArn() *string
	SetGmtCreate(v string) *ChannelAssemblySource
	GetGmtCreate() *string
	SetGmtModified(v string) *ChannelAssemblySource
	GetGmtModified() *string
	SetHttpPackageConfigurations(v string) *ChannelAssemblySource
	GetHttpPackageConfigurations() *string
	SetSourceLocationName(v string) *ChannelAssemblySource
	GetSourceLocationName() *string
	SetSourceName(v string) *ChannelAssemblySource
	GetSourceName() *string
	SetSourceType(v string) *ChannelAssemblySource
	GetSourceType() *string
	SetState(v int32) *ChannelAssemblySource
	GetState() *int32
}

type ChannelAssemblySource struct {
	Arn                       *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	GmtCreate                 *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified               *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HttpPackageConfigurations *string `json:"HttpPackageConfigurations,omitempty" xml:"HttpPackageConfigurations,omitempty"`
	SourceLocationName        *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName                *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceType                *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	State                     *int32  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ChannelAssemblySource) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblySource) GoString() string {
	return s.String()
}

func (s *ChannelAssemblySource) GetArn() *string {
	return s.Arn
}

func (s *ChannelAssemblySource) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ChannelAssemblySource) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ChannelAssemblySource) GetHttpPackageConfigurations() *string {
	return s.HttpPackageConfigurations
}

func (s *ChannelAssemblySource) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ChannelAssemblySource) GetSourceName() *string {
	return s.SourceName
}

func (s *ChannelAssemblySource) GetSourceType() *string {
	return s.SourceType
}

func (s *ChannelAssemblySource) GetState() *int32 {
	return s.State
}

func (s *ChannelAssemblySource) SetArn(v string) *ChannelAssemblySource {
	s.Arn = &v
	return s
}

func (s *ChannelAssemblySource) SetGmtCreate(v string) *ChannelAssemblySource {
	s.GmtCreate = &v
	return s
}

func (s *ChannelAssemblySource) SetGmtModified(v string) *ChannelAssemblySource {
	s.GmtModified = &v
	return s
}

func (s *ChannelAssemblySource) SetHttpPackageConfigurations(v string) *ChannelAssemblySource {
	s.HttpPackageConfigurations = &v
	return s
}

func (s *ChannelAssemblySource) SetSourceLocationName(v string) *ChannelAssemblySource {
	s.SourceLocationName = &v
	return s
}

func (s *ChannelAssemblySource) SetSourceName(v string) *ChannelAssemblySource {
	s.SourceName = &v
	return s
}

func (s *ChannelAssemblySource) SetSourceType(v string) *ChannelAssemblySource {
	s.SourceType = &v
	return s
}

func (s *ChannelAssemblySource) SetState(v int32) *ChannelAssemblySource {
	s.State = &v
	return s
}

func (s *ChannelAssemblySource) Validate() error {
	return dara.Validate(s)
}
