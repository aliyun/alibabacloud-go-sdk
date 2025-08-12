// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeAlert(v bool) *DescribeSiteMonitorAttributeRequest
	GetIncludeAlert() *bool
	SetRegionId(v string) *DescribeSiteMonitorAttributeRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeSiteMonitorAttributeRequest
	GetTaskId() *string
}

type DescribeSiteMonitorAttributeRequest struct {
	// Specifies whether to return the information of the alert rules that are configured for the site monitoring task. Valid values:
	//
	// 	- true: The system returns the information of the alert rules that are configured for the site monitoring task.
	//
	// 	- false (default): The system does not return the information of the alert rules that are configured for the site monitoring task.
	//
	// example:
	//
	// false
	IncludeAlert *bool   `json:"IncludeAlert,omitempty" xml:"IncludeAlert,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the site monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc641dff-c19d-45f3-ad0a-818a0c4f****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSiteMonitorAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeRequest) GetIncludeAlert() *bool {
	return s.IncludeAlert
}

func (s *DescribeSiteMonitorAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorAttributeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSiteMonitorAttributeRequest) SetIncludeAlert(v bool) *DescribeSiteMonitorAttributeRequest {
	s.IncludeAlert = &v
	return s
}

func (s *DescribeSiteMonitorAttributeRequest) SetRegionId(v string) *DescribeSiteMonitorAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeRequest) SetTaskId(v string) *DescribeSiteMonitorAttributeRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeRequest) Validate() error {
	return dara.Validate(s)
}
