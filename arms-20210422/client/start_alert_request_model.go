// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v string) *StartAlertRequest
	GetAlertId() *string
	SetRegionId(v string) *StartAlertRequest
	GetRegionId() *string
}

type StartAlertRequest struct {
	// This parameter is required.
	AlertId *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAlertRequest) GoString() string {
	return s.String()
}

func (s *StartAlertRequest) GetAlertId() *string {
	return s.AlertId
}

func (s *StartAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartAlertRequest) SetAlertId(v string) *StartAlertRequest {
	s.AlertId = &v
	return s
}

func (s *StartAlertRequest) SetRegionId(v string) *StartAlertRequest {
	s.RegionId = &v
	return s
}

func (s *StartAlertRequest) Validate() error {
	return dara.Validate(s)
}
