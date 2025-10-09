// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUrlObservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSdkType(v string) *CreateUrlObservationRequest
	GetSdkType() *string
	SetSiteId(v int64) *CreateUrlObservationRequest
	GetSiteId() *int64
	SetUrl(v string) *CreateUrlObservationRequest
	GetUrl() *string
}

type CreateUrlObservationRequest struct {
	// SDK integration. Supported
	//
	// 	- **automatic**
	//
	// 	- **manual**
	//
	// This parameter is required.
	//
	// example:
	//
	// automatic
	SdkType *string `json:"SdkType,omitempty" xml:"SdkType,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The URL of the web page to monitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com/test
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateUrlObservationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUrlObservationRequest) GoString() string {
	return s.String()
}

func (s *CreateUrlObservationRequest) GetSdkType() *string {
	return s.SdkType
}

func (s *CreateUrlObservationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateUrlObservationRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateUrlObservationRequest) SetSdkType(v string) *CreateUrlObservationRequest {
	s.SdkType = &v
	return s
}

func (s *CreateUrlObservationRequest) SetSiteId(v int64) *CreateUrlObservationRequest {
	s.SiteId = &v
	return s
}

func (s *CreateUrlObservationRequest) SetUrl(v string) *CreateUrlObservationRequest {
	s.Url = &v
	return s
}

func (s *CreateUrlObservationRequest) Validate() error {
	return dara.Validate(s)
}
