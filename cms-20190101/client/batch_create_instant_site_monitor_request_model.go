// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateInstantSiteMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *BatchCreateInstantSiteMonitorRequest
	GetRegionId() *string
	SetTaskList(v []*BatchCreateInstantSiteMonitorRequestTaskList) *BatchCreateInstantSiteMonitorRequest
	GetTaskList() []*BatchCreateInstantSiteMonitorRequestTaskList
}

type BatchCreateInstantSiteMonitorRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of site monitoring tasks.
	//
	// > You must create at least one site monitoring task. The `Address`, `TaskName`, and `TaskType` parameters are required.
	//
	// This parameter is required.
	TaskList []*BatchCreateInstantSiteMonitorRequestTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s BatchCreateInstantSiteMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateInstantSiteMonitorRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateInstantSiteMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchCreateInstantSiteMonitorRequest) GetTaskList() []*BatchCreateInstantSiteMonitorRequestTaskList {
	return s.TaskList
}

func (s *BatchCreateInstantSiteMonitorRequest) SetRegionId(v string) *BatchCreateInstantSiteMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorRequest) SetTaskList(v []*BatchCreateInstantSiteMonitorRequestTaskList) *BatchCreateInstantSiteMonitorRequest {
	s.TaskList = v
	return s
}

func (s *BatchCreateInstantSiteMonitorRequest) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateInstantSiteMonitorRequestTaskList struct {
	// The URL or IP address of the site monitoring task.
	//
	// > You must create at least one site monitoring task. The `Address`, `TaskName`, and `TaskType` parameters are required.
	//
	// example:
	//
	// https://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The information about the detection points. If you leave this parameter empty, the system randomly selects three detection points.
	//
	// The value must be a `JSONArray`. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` corresponds to Beijing, Hangzhou, and Qingdao.
	//
	// For more information about how to obtain detection point information, see [DescribeSiteMonitorISPCityList](https://help.aliyun.com/document_detail/115045.html).
	//
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCities *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	// The advanced extension options for the protocol type of the site monitoring task. Different protocol types correspond to different extension options.
	//
	// example:
	//
	// {"time_out":5000}
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	// The name of the site monitoring task.
	//
	// <props="china">
	//
	// The value must be 4 to 100 characters in length and can contain English letters, numbers, underscores (_), and Chinese characters.
	//
	//
	//
	// <props="intl">
	//
	// The name of the site monitoring task.
	//
	//
	//
	// <props="partner">
	//
	// The name must be 4 to 100 characters in length and can contain letters, digits, and underscores (_).
	//
	//
	//
	// > You must create at least one site monitoring task. The `Address`, `TaskName`, and `TaskType` parameters are required.
	//
	// example:
	//
	// HangZhou_ECS1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The protocol type of the monitoring task.
	//
	// Valid values: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	//
	// > You must create at least one site monitoring task. The `Address`, `TaskName`, and `TaskType` parameters are required.
	//
	// example:
	//
	// HTTP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s BatchCreateInstantSiteMonitorRequestTaskList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateInstantSiteMonitorRequestTaskList) GoString() string {
	return s.String()
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) GetAddress() *string {
	return s.Address
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) GetIspCities() *string {
	return s.IspCities
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) GetTaskName() *string {
	return s.TaskName
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) GetTaskType() *string {
	return s.TaskType
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) SetAddress(v string) *BatchCreateInstantSiteMonitorRequestTaskList {
	s.Address = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) SetIspCities(v string) *BatchCreateInstantSiteMonitorRequestTaskList {
	s.IspCities = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) SetOptionsJson(v string) *BatchCreateInstantSiteMonitorRequestTaskList {
	s.OptionsJson = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) SetTaskName(v string) *BatchCreateInstantSiteMonitorRequestTaskList {
	s.TaskName = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) SetTaskType(v string) *BatchCreateInstantSiteMonitorRequestTaskList {
	s.TaskType = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorRequestTaskList) Validate() error {
	return dara.Validate(s)
}
