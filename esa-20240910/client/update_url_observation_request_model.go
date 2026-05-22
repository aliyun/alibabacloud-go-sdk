// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUrlObservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateUrlObservationRequest
	GetConfigId() *int64
	SetSdkType(v string) *UpdateUrlObservationRequest
	GetSdkType() *string
	SetSiteId(v int64) *UpdateUrlObservationRequest
	GetSiteId() *int64
}

type UpdateUrlObservationRequest struct {
	// The ID of the configuration. You can call the [ListUrlObservations](~~ListUrlObservations~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// SDK integration. Valid values:
	//
	// 	- **automatic*	- (Recommended)
	//
	// 	- **manual**
	//
	// This parameter is required.
	//
	// example:
	//
	// manual
	SdkType *string `json:"SdkType,omitempty" xml:"SdkType,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateUrlObservationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUrlObservationRequest) GoString() string {
	return s.String()
}

func (s *UpdateUrlObservationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateUrlObservationRequest) GetSdkType() *string {
	return s.SdkType
}

func (s *UpdateUrlObservationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateUrlObservationRequest) SetConfigId(v int64) *UpdateUrlObservationRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateUrlObservationRequest) SetSdkType(v string) *UpdateUrlObservationRequest {
	s.SdkType = &v
	return s
}

func (s *UpdateUrlObservationRequest) SetSiteId(v int64) *UpdateUrlObservationRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateUrlObservationRequest) Validate() error {
	return dara.Validate(s)
}
