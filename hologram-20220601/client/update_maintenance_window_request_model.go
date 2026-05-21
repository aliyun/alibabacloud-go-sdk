// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaintenanceWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *UpdateMaintenanceWindowRequest
	GetEndTime() *string
	SetStartTime(v string) *UpdateMaintenanceWindowRequest
	GetStartTime() *string
}

type UpdateMaintenanceWindowRequest struct {
	// example:
	//
	// 02:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateMaintenanceWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaintenanceWindowRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaintenanceWindowRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateMaintenanceWindowRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateMaintenanceWindowRequest) SetEndTime(v string) *UpdateMaintenanceWindowRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateMaintenanceWindowRequest) SetStartTime(v string) *UpdateMaintenanceWindowRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateMaintenanceWindowRequest) Validate() error {
	return dara.Validate(s)
}
