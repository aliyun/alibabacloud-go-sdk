// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayMaintenancePeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaintenancePeriod(v *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) *UpdateGatewayMaintenancePeriodRequest
	GetMaintenancePeriod() *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod
}

type UpdateGatewayMaintenancePeriodRequest struct {
	MaintenancePeriod *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod `json:"maintenancePeriod,omitempty" xml:"maintenancePeriod,omitempty" type:"Struct"`
}

func (s UpdateGatewayMaintenancePeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayMaintenancePeriodRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayMaintenancePeriodRequest) GetMaintenancePeriod() *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod {
	return s.MaintenancePeriod
}

func (s *UpdateGatewayMaintenancePeriodRequest) SetMaintenancePeriod(v *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) *UpdateGatewayMaintenancePeriodRequest {
	s.MaintenancePeriod = v
	return s
}

func (s *UpdateGatewayMaintenancePeriodRequest) Validate() error {
	if s.MaintenancePeriod != nil {
		if err := s.MaintenancePeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayMaintenancePeriodRequestMaintenancePeriod struct {
	// example:
	//
	// 02:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 02:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) GoString() string {
	return s.String()
}

func (s *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) SetEndTime(v string) *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod {
	s.EndTime = &v
	return s
}

func (s *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) SetStartTime(v string) *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod {
	s.StartTime = &v
	return s
}

func (s *UpdateGatewayMaintenancePeriodRequestMaintenancePeriod) Validate() error {
	return dara.Validate(s)
}
