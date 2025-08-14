// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSource interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *Source
	GetArn() *string
	SetGmtCreate(v string) *Source
	GetGmtCreate() *string
	SetGmtModified(v string) *Source
	GetGmtModified() *string
	SetHttpPackageConfigurations(v string) *Source
	GetHttpPackageConfigurations() *string
	SetSourceLocationName(v string) *Source
	GetSourceLocationName() *string
	SetSourceName(v string) *Source
	GetSourceName() *string
	SetSourceType(v string) *Source
	GetSourceType() *string
	SetState(v int32) *Source
	GetState() *int32
}

type Source struct {
	Arn                       *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	GmtCreate                 *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified               *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HttpPackageConfigurations *string `json:"HttpPackageConfigurations,omitempty" xml:"HttpPackageConfigurations,omitempty"`
	SourceLocationName        *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName                *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceType                *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	State                     *int32  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s Source) String() string {
	return dara.Prettify(s)
}

func (s Source) GoString() string {
	return s.String()
}

func (s *Source) GetArn() *string {
	return s.Arn
}

func (s *Source) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Source) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Source) GetHttpPackageConfigurations() *string {
	return s.HttpPackageConfigurations
}

func (s *Source) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *Source) GetSourceName() *string {
	return s.SourceName
}

func (s *Source) GetSourceType() *string {
	return s.SourceType
}

func (s *Source) GetState() *int32 {
	return s.State
}

func (s *Source) SetArn(v string) *Source {
	s.Arn = &v
	return s
}

func (s *Source) SetGmtCreate(v string) *Source {
	s.GmtCreate = &v
	return s
}

func (s *Source) SetGmtModified(v string) *Source {
	s.GmtModified = &v
	return s
}

func (s *Source) SetHttpPackageConfigurations(v string) *Source {
	s.HttpPackageConfigurations = &v
	return s
}

func (s *Source) SetSourceLocationName(v string) *Source {
	s.SourceLocationName = &v
	return s
}

func (s *Source) SetSourceName(v string) *Source {
	s.SourceName = &v
	return s
}

func (s *Source) SetSourceType(v string) *Source {
	s.SourceType = &v
	return s
}

func (s *Source) SetState(v int32) *Source {
	s.State = &v
	return s
}

func (s *Source) Validate() error {
	return dara.Validate(s)
}
