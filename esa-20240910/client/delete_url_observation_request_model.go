// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrlObservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteUrlObservationRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteUrlObservationRequest
	GetSiteId() *int64
}

type DeleteUrlObservationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteUrlObservationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrlObservationRequest) GoString() string {
	return s.String()
}

func (s *DeleteUrlObservationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteUrlObservationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteUrlObservationRequest) SetConfigId(v int64) *DeleteUrlObservationRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteUrlObservationRequest) SetSiteId(v int64) *DeleteUrlObservationRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteUrlObservationRequest) Validate() error {
	return dara.Validate(s)
}
