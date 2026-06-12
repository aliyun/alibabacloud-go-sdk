// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*string) *LaunchServiceRequest
	GetCategories() []*string
	SetClientToken(v string) *LaunchServiceRequest
	GetClientToken() *string
	SetRecommend(v bool) *LaunchServiceRequest
	GetRecommend() *bool
	SetRegionId(v string) *LaunchServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *LaunchServiceRequest
	GetServiceId() *string
	SetServiceVersion(v string) *LaunchServiceRequest
	GetServiceVersion() *string
}

type LaunchServiceRequest struct {
	// The service categories.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// A client token used to ensure the idempotence of the request. Generate a unique value from your client for each request. The ClientToken parameter supports only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to recommend the service when publishing it to the Service Catalog.
	//
	// example:
	//
	// false
	Recommend *bool `json:"Recommend,omitempty" xml:"Recommend,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s LaunchServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s LaunchServiceRequest) GoString() string {
	return s.String()
}

func (s *LaunchServiceRequest) GetCategories() []*string {
	return s.Categories
}

func (s *LaunchServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *LaunchServiceRequest) GetRecommend() *bool {
	return s.Recommend
}

func (s *LaunchServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LaunchServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *LaunchServiceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *LaunchServiceRequest) SetCategories(v []*string) *LaunchServiceRequest {
	s.Categories = v
	return s
}

func (s *LaunchServiceRequest) SetClientToken(v string) *LaunchServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *LaunchServiceRequest) SetRecommend(v bool) *LaunchServiceRequest {
	s.Recommend = &v
	return s
}

func (s *LaunchServiceRequest) SetRegionId(v string) *LaunchServiceRequest {
	s.RegionId = &v
	return s
}

func (s *LaunchServiceRequest) SetServiceId(v string) *LaunchServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *LaunchServiceRequest) SetServiceVersion(v string) *LaunchServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *LaunchServiceRequest) Validate() error {
	return dara.Validate(s)
}
