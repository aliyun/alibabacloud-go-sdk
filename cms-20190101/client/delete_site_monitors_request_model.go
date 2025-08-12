// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsDeleteAlarms(v bool) *DeleteSiteMonitorsRequest
	GetIsDeleteAlarms() *bool
	SetRegionId(v string) *DeleteSiteMonitorsRequest
	GetRegionId() *string
	SetTaskIds(v string) *DeleteSiteMonitorsRequest
	GetTaskIds() *string
}

type DeleteSiteMonitorsRequest struct {
	// Specifies whether to delete the alert rules configured for the site monitoring tasks. Valid values:
	//
	// 	- true (default value)
	//
	// 	- false
	//
	// example:
	//
	// true
	IsDeleteAlarms *bool   `json:"IsDeleteAlarms,omitempty" xml:"IsDeleteAlarms,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the site monitoring tasks that you want to delete. Separate multiple task IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 01adacc2-ece5-41b6-afa2-3143ab5d****,43bd1ead-514f-4524-813e-228ce091****
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DeleteSiteMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteMonitorsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsRequest) GetIsDeleteAlarms() *bool {
	return s.IsDeleteAlarms
}

func (s *DeleteSiteMonitorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSiteMonitorsRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *DeleteSiteMonitorsRequest) SetIsDeleteAlarms(v bool) *DeleteSiteMonitorsRequest {
	s.IsDeleteAlarms = &v
	return s
}

func (s *DeleteSiteMonitorsRequest) SetRegionId(v string) *DeleteSiteMonitorsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSiteMonitorsRequest) SetTaskIds(v string) *DeleteSiteMonitorsRequest {
	s.TaskIds = &v
	return s
}

func (s *DeleteSiteMonitorsRequest) Validate() error {
	return dara.Validate(s)
}
