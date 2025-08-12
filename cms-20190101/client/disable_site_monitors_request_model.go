// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSiteMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DisableSiteMonitorsRequest
	GetRegionId() *string
	SetTaskIds(v string) *DisableSiteMonitorsRequest
	GetTaskIds() *string
}

type DisableSiteMonitorsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the site monitoring task. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 49f7b317-7645-4cc9-94fd-ea42e522****,49f7b317-7645-4cc9-94fd-ea42e522****
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DisableSiteMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSiteMonitorsRequest) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableSiteMonitorsRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *DisableSiteMonitorsRequest) SetRegionId(v string) *DisableSiteMonitorsRequest {
	s.RegionId = &v
	return s
}

func (s *DisableSiteMonitorsRequest) SetTaskIds(v string) *DisableSiteMonitorsRequest {
	s.TaskIds = &v
	return s
}

func (s *DisableSiteMonitorsRequest) Validate() error {
	return dara.Validate(s)
}
