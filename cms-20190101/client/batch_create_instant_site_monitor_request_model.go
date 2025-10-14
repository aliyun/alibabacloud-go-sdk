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
	// The site monitoring tasks.
	//
	// >  You must create at least one site monitoring task. You must specify all of the `Address`, `TaskName`, and `TaskType` parameters in each request.
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
	// The URL or IP address that is monitored by the task.
	//
	// >  You must create at least one site monitoring task. You must specify all of the `Address`, `TaskName`, and `TaskType` parameters in each request.
	//
	// example:
	//
	// https://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The detection points. If you leave this parameter empty, the system randomly selects three detection points.
	//
	// The value is a `JSON array`. Example: `{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}`. The values of the city field indicate Beijing, Hangzhou, and Qingdao.
	//
	// For information about how to obtain detection points, see [DescribeSiteMonitorISPCityList](https://help.aliyun.com/document_detail/115045.html).
	//
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCities *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	// The extended options of the protocol that is used by the site monitoring task. The options vary based on the protocol.
	//
	// example:
	//
	// {"time_out":5000}
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	// The name of the site monitoring task.
	//
	// The name must be 4 to 100 characters in length, and can contain letters, digits, and underscores (_).
	//
	// >  You must create at least one site monitoring task. You must specify all of the `Address`, `TaskName`, and `TaskType` parameters in each request.
	//
	// example:
	//
	// HangZhou_ECS1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the site monitoring task.
	//
	// Valid values: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	//
	// >  You must create at least one site monitoring task. You must specify all of the `Address`, `TaskName`, and `TaskType` parameters in each request.
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
