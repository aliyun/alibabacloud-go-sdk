// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmEventStackInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventName(v string) *DescribeAlarmEventStackInfoRequest
	GetEventName() *string
	SetLang(v string) *DescribeAlarmEventStackInfoRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v int64) *DescribeAlarmEventStackInfoRequest
	GetResourceDirectoryAccountId() *int64
	SetSourceIp(v string) *DescribeAlarmEventStackInfoRequest
	GetSourceIp() *string
	SetUniqueInfo(v string) *DescribeAlarmEventStackInfoRequest
	GetUniqueInfo() *string
	SetUuid(v string) *DescribeAlarmEventStackInfoRequest
	GetUuid() *string
}

type DescribeAlarmEventStackInfoRequest struct {
	// The name of the event.
	//
	// >  You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to query the names of events.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mining program
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the alert event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1fbe8d16727f61d1478a674d6fa0****
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	// The UUID of the server to query.
	//
	// >  You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18b7336e-d469-473b-af83-8e5420f9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAlarmEventStackInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventStackInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAlarmEventStackInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAlarmEventStackInfoRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeAlarmEventStackInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAlarmEventStackInfoRequest) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *DescribeAlarmEventStackInfoRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAlarmEventStackInfoRequest) SetEventName(v string) *DescribeAlarmEventStackInfoRequest {
	s.EventName = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetLang(v string) *DescribeAlarmEventStackInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetResourceDirectoryAccountId(v int64) *DescribeAlarmEventStackInfoRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetSourceIp(v string) *DescribeAlarmEventStackInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetUniqueInfo(v string) *DescribeAlarmEventStackInfoRequest {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetUuid(v string) *DescribeAlarmEventStackInfoRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) Validate() error {
	return dara.Validate(s)
}
