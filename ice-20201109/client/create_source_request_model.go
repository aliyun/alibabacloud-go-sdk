// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHttpPackageConfigurations(v string) *CreateSourceRequest
	GetHttpPackageConfigurations() *string
	SetSourceLocationName(v string) *CreateSourceRequest
	GetSourceLocationName() *string
	SetSourceName(v string) *CreateSourceRequest
	GetSourceName() *string
	SetSourceType(v string) *CreateSourceRequest
	GetSourceType() *string
}

type CreateSourceRequest struct {
	// The source configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// “[{
	//
	// 	"sourceGroupName": "mySourceGroup-1",
	//
	// 	"relativePath": "group1/hls.m3u8",
	//
	// 	"type": "hls"
	//
	// }]”
	HttpPackageConfigurations *string `json:"HttpPackageConfigurations,omitempty" xml:"HttpPackageConfigurations,omitempty"`
	// The name of the source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySourceLocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	// The name of the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyVodSource
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// The source type. Valid values: vodSource and liveSource.
	//
	// This parameter is required.
	//
	// example:
	//
	// vodSource
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateSourceRequest) GetHttpPackageConfigurations() *string {
	return s.HttpPackageConfigurations
}

func (s *CreateSourceRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *CreateSourceRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *CreateSourceRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateSourceRequest) SetHttpPackageConfigurations(v string) *CreateSourceRequest {
	s.HttpPackageConfigurations = &v
	return s
}

func (s *CreateSourceRequest) SetSourceLocationName(v string) *CreateSourceRequest {
	s.SourceLocationName = &v
	return s
}

func (s *CreateSourceRequest) SetSourceName(v string) *CreateSourceRequest {
	s.SourceName = &v
	return s
}

func (s *CreateSourceRequest) SetSourceType(v string) *CreateSourceRequest {
	s.SourceType = &v
	return s
}

func (s *CreateSourceRequest) Validate() error {
	return dara.Validate(s)
}
