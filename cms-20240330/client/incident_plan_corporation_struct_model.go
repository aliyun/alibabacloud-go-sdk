// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentPlanCorporationStruct interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *IncidentPlanCorporationStruct
	GetChannel() *string
	SetRobotId(v string) *IncidentPlanCorporationStruct
	GetRobotId() *string
}

type IncidentPlanCorporationStruct struct {
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	RobotId *string `json:"robotId,omitempty" xml:"robotId,omitempty"`
}

func (s IncidentPlanCorporationStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentPlanCorporationStruct) GoString() string {
	return s.String()
}

func (s *IncidentPlanCorporationStruct) GetChannel() *string {
	return s.Channel
}

func (s *IncidentPlanCorporationStruct) GetRobotId() *string {
	return s.RobotId
}

func (s *IncidentPlanCorporationStruct) SetChannel(v string) *IncidentPlanCorporationStruct {
	s.Channel = &v
	return s
}

func (s *IncidentPlanCorporationStruct) SetRobotId(v string) *IncidentPlanCorporationStruct {
	s.RobotId = &v
	return s
}

func (s *IncidentPlanCorporationStruct) Validate() error {
	return dara.Validate(s)
}
