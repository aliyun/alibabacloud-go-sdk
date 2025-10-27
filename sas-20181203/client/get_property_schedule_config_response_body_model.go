// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPropertyScheduleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyScheduleConfig(v *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) *GetPropertyScheduleConfigResponseBody
	GetPropertyScheduleConfig() *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig
	SetRequestId(v string) *GetPropertyScheduleConfigResponseBody
	GetRequestId() *string
}

type GetPropertyScheduleConfigResponseBody struct {
	// The configurations for the collection frequency of asset fingerprints.
	PropertyScheduleConfig *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig `json:"PropertyScheduleConfig,omitempty" xml:"PropertyScheduleConfig,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B37C9052-A73E-4707-A024-92477028****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPropertyScheduleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPropertyScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPropertyScheduleConfigResponseBody) GetPropertyScheduleConfig() *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig {
	return s.PropertyScheduleConfig
}

func (s *GetPropertyScheduleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPropertyScheduleConfigResponseBody) SetPropertyScheduleConfig(v *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) *GetPropertyScheduleConfigResponseBody {
	s.PropertyScheduleConfig = v
	return s
}

func (s *GetPropertyScheduleConfigResponseBody) SetRequestId(v string) *GetPropertyScheduleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPropertyScheduleConfigResponseBody) Validate() error {
	if s.PropertyScheduleConfig != nil {
		if err := s.PropertyScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig struct {
	// The timestamp when the next collection of asset fingerprints starts. Unit: milliseconds.
	//
	// example:
	//
	// 1671630647018
	NextScheduleTime *int64 `json:"NextScheduleTime,omitempty" xml:"NextScheduleTime,omitempty"`
	// The collection frequency of asset fingerprints. Valid values:
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
	// example:
	//
	// 3
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The type of the asset fingerprints. Valid values:
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
	// 	- **scheduler_sca_period**: middleware
	//
	// 	- **scheduler_autorun_period**: startup item
	//
	// 	- **scheduler_lkm_period**: kernel module
	//
	// 	- **scheduler_sca_proxy_period**: website
	//
	// example:
	//
	// scheduler_account_period
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) GoString() string {
	return s.String()
}

func (s *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) GetNextScheduleTime() *int64 {
	return s.NextScheduleTime
}

func (s *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) GetType() *string {
	return s.Type
}

func (s *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) SetNextScheduleTime(v int64) *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig {
	s.NextScheduleTime = &v
	return s
}

func (s *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) SetScheduleTime(v string) *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig {
	s.ScheduleTime = &v
	return s
}

func (s *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) SetType(v string) *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig {
	s.Type = &v
	return s
}

func (s *GetPropertyScheduleConfigResponseBodyPropertyScheduleConfig) Validate() error {
	return dara.Validate(s)
}
