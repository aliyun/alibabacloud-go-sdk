// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomIncidentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidents(v string) *SendCustomIncidentsRequest
	GetIncidents() *string
	SetProductType(v string) *SendCustomIncidentsRequest
	GetProductType() *string
	SetRegionId(v string) *SendCustomIncidentsRequest
	GetRegionId() *string
}

type SendCustomIncidentsRequest struct {
	// This parameter is required.
	Incidents *string `json:"Incidents,omitempty" xml:"Incidents,omitempty"`
	// This parameter is required.
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SendCustomIncidentsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendCustomIncidentsRequest) GoString() string {
	return s.String()
}

func (s *SendCustomIncidentsRequest) GetIncidents() *string {
	return s.Incidents
}

func (s *SendCustomIncidentsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *SendCustomIncidentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SendCustomIncidentsRequest) SetIncidents(v string) *SendCustomIncidentsRequest {
	s.Incidents = &v
	return s
}

func (s *SendCustomIncidentsRequest) SetProductType(v string) *SendCustomIncidentsRequest {
	s.ProductType = &v
	return s
}

func (s *SendCustomIncidentsRequest) SetRegionId(v string) *SendCustomIncidentsRequest {
	s.RegionId = &v
	return s
}

func (s *SendCustomIncidentsRequest) Validate() error {
	return dara.Validate(s)
}
