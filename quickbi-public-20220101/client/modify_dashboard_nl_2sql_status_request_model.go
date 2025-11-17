// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDashboardNl2sqlStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDashboardIds(v string) *ModifyDashboardNl2sqlStatusRequest
	GetDashboardIds() *string
	SetStatus(v int32) *ModifyDashboardNl2sqlStatusRequest
	GetStatus() *int32
}

type ModifyDashboardNl2sqlStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// asda,sadaf
	DashboardIds *string `json:"DashboardIds,omitempty" xml:"DashboardIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDashboardNl2sqlStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDashboardNl2sqlStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDashboardNl2sqlStatusRequest) GetDashboardIds() *string {
	return s.DashboardIds
}

func (s *ModifyDashboardNl2sqlStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyDashboardNl2sqlStatusRequest) SetDashboardIds(v string) *ModifyDashboardNl2sqlStatusRequest {
	s.DashboardIds = &v
	return s
}

func (s *ModifyDashboardNl2sqlStatusRequest) SetStatus(v int32) *ModifyDashboardNl2sqlStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyDashboardNl2sqlStatusRequest) Validate() error {
	return dara.Validate(s)
}
