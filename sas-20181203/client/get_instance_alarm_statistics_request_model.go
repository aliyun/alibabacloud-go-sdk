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
	SetResourceDirectoryAccountId(v int64) *GetInstanceAlarmStatisticsRequest
	GetResourceDirectoryAccountId() *int64
	SetUuid(v string) *GetInstanceAlarmStatisticsRequest
	GetUuid() *string
}

type GetInstanceAlarmStatisticsRequest struct {
	// The data source from which instance alert statistics are collected. Default value: aqs.
	//
	// Valid values:
	//
	// - **sas**: Threat Detection Service data source.
	//
	// - **aqs**: alert event data.
	//
	// - **honeypot**: cloud honeypot.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource folder.
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The UUID of the server to query.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
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

func (s *GetInstanceAlarmStatisticsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetInstanceAlarmStatisticsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetInstanceAlarmStatisticsRequest) SetFrom(v string) *GetInstanceAlarmStatisticsRequest {
	s.From = &v
	return s
}

func (s *GetInstanceAlarmStatisticsRequest) SetResourceDirectoryAccountId(v int64) *GetInstanceAlarmStatisticsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetInstanceAlarmStatisticsRequest) SetUuid(v string) *GetInstanceAlarmStatisticsRequest {
	s.Uuid = &v
	return s
}

func (s *GetInstanceAlarmStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
