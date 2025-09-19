// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPropertyScheduleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScheduleTime(v string) *ModifyPropertyScheduleConfigRequest
	GetScheduleTime() *string
	SetType(v string) *ModifyPropertyScheduleConfigRequest
	GetType() *string
}

type ModifyPropertyScheduleConfigRequest struct {
	// The new collection frequency of asset fingerprints. Valid values:
	//
	// 	- **0**: disabled, which indicates that the asset fingerprints are not automatically or periodically collected.
	//
	// 	- **1**: collects asset fingerprints once an hour.
	//
	// 	- **3**: collects asset fingerprints once every 3 hours.
	//
	// 	- **12**: collects asset fingerprints once every 12 hours.
	//
	// 	- **24**: collects asset fingerprints once a day.
	//
	// 	- **168**: collects asset fingerprints once every 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The type of the asset fingerprints for which you want to modify the collection frequency. Valid values:
	//
	// 	- **scheduler_port_period**: listening port
	//
	// 	- **scheduler_process_period**: running process
	//
	// 	- **scheduler_account_period**: account
	//
	// 	- **scheduler_software_period**: software
	//
	// 	- **scheduler_cron_period**: scheduled task
	//
	// 	- **scheduler_sca_period**: middleware, database, or web service
	//
	// 	- **scheduler_autorun_period**: startup item
	//
	// 	- **scheduler_lkm_period**: kernel module
	//
	// 	- **scheduler_sca_proxy_period**: website
	//
	// This parameter is required.
	//
	// example:
	//
	// scheduler_port_period
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPropertyScheduleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPropertyScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyPropertyScheduleConfigRequest) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ModifyPropertyScheduleConfigRequest) GetType() *string {
	return s.Type
}

func (s *ModifyPropertyScheduleConfigRequest) SetScheduleTime(v string) *ModifyPropertyScheduleConfigRequest {
	s.ScheduleTime = &v
	return s
}

func (s *ModifyPropertyScheduleConfigRequest) SetType(v string) *ModifyPropertyScheduleConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyPropertyScheduleConfigRequest) Validate() error {
	return dara.Validate(s)
}
