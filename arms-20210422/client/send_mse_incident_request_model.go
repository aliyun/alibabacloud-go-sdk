// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMseIncidentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidents(v string) *SendMseIncidentRequest
	GetIncidents() *string
	SetRegionId(v string) *SendMseIncidentRequest
	GetRegionId() *string
}

type SendMseIncidentRequest struct {
	// This parameter is required.
	Incidents *string `json:"Incidents,omitempty" xml:"Incidents,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SendMseIncidentRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMseIncidentRequest) GoString() string {
	return s.String()
}

func (s *SendMseIncidentRequest) GetIncidents() *string {
	return s.Incidents
}

func (s *SendMseIncidentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SendMseIncidentRequest) SetIncidents(v string) *SendMseIncidentRequest {
	s.Incidents = &v
	return s
}

func (s *SendMseIncidentRequest) SetRegionId(v string) *SendMseIncidentRequest {
	s.RegionId = &v
	return s
}

func (s *SendMseIncidentRequest) Validate() error {
	return dara.Validate(s)
}
