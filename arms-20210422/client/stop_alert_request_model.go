// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v string) *StopAlertRequest
	GetAlertId() *string
	SetRegionId(v string) *StopAlertRequest
	GetRegionId() *string
}

type StopAlertRequest struct {
	// This parameter is required.
	AlertId *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAlertRequest) GoString() string {
	return s.String()
}

func (s *StopAlertRequest) GetAlertId() *string {
	return s.AlertId
}

func (s *StopAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopAlertRequest) SetAlertId(v string) *StopAlertRequest {
	s.AlertId = &v
	return s
}

func (s *StopAlertRequest) SetRegionId(v string) *StopAlertRequest {
	s.RegionId = &v
	return s
}

func (s *StopAlertRequest) Validate() error {
	return dara.Validate(s)
}
