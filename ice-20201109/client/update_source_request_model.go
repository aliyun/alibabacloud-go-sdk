// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHttpPackageConfigurations(v string) *UpdateSourceRequest
	GetHttpPackageConfigurations() *string
	SetSourceLocationName(v string) *UpdateSourceRequest
	GetSourceLocationName() *string
	SetSourceName(v string) *UpdateSourceRequest
	GetSourceName() *string
	SetSourceType(v string) *UpdateSourceRequest
	GetSourceType() *string
}

type UpdateSourceRequest struct {
	// The source configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{
	//
	// 	"sourceGroupName": "mySourceGroup-1",
	//
	// 	"relativePath": "group1/hls.m3u8",
	//
	// 	"packageType": "hls"
	//
	// }]
	HttpPackageConfigurations *string `json:"HttpPackageConfigurations,omitempty" xml:"HttpPackageConfigurations,omitempty"`
	// The name of the source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySourcelocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	// The name of the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySource
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

func (s UpdateSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateSourceRequest) GetHttpPackageConfigurations() *string {
	return s.HttpPackageConfigurations
}

func (s *UpdateSourceRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *UpdateSourceRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *UpdateSourceRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateSourceRequest) SetHttpPackageConfigurations(v string) *UpdateSourceRequest {
	s.HttpPackageConfigurations = &v
	return s
}

func (s *UpdateSourceRequest) SetSourceLocationName(v string) *UpdateSourceRequest {
	s.SourceLocationName = &v
	return s
}

func (s *UpdateSourceRequest) SetSourceName(v string) *UpdateSourceRequest {
	s.SourceName = &v
	return s
}

func (s *UpdateSourceRequest) SetSourceType(v string) *UpdateSourceRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateSourceRequest) Validate() error {
	return dara.Validate(s)
}
