// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAlarmStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *GetInstanceAlarmStatisticsRequest
	GetFrom() *string
	SetUuid(v string) *GetInstanceAlarmStatisticsRequest
	GetUuid() *string
}

type GetInstanceAlarmStatisticsRequest struct {
	// The data source for statistics on instance alarms, with a default value of aqs:
	//
	// - *sas*: Situation Awareness data source
	//
	// - *aqs*: Alarm event data
	//
	// - *honeypot*: Honeypot
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The UUID of the server to be queried.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) API to obtain this parameter.
	//
	// example:
	//
	// 00fea5a1-9792-4373-ab1e-bb6536ba****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetInstanceAlarmStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAlarmStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceAlarmStatisticsRequest) GetFrom() *string {
	return s.From
}

func (s *GetInstanceAlarmStatisticsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetInstanceAlarmStatisticsRequest) SetFrom(v string) *GetInstanceAlarmStatisticsRequest {
	s.From = &v
	return s
}

func (s *GetInstanceAlarmStatisticsRequest) SetUuid(v string) *GetInstanceAlarmStatisticsRequest {
	s.Uuid = &v
	return s
}

func (s *GetInstanceAlarmStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
