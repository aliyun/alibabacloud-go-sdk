// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityCheckScheduleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDaysOfWeek(v string) *ModifySecurityCheckScheduleConfigRequest
	GetDaysOfWeek() *string
	SetEndTime(v int32) *ModifySecurityCheckScheduleConfigRequest
	GetEndTime() *int32
	SetLang(v string) *ModifySecurityCheckScheduleConfigRequest
	GetLang() *string
	SetResourceOwnerId(v int64) *ModifySecurityCheckScheduleConfigRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *ModifySecurityCheckScheduleConfigRequest
	GetSourceIp() *string
	SetStartTime(v int32) *ModifySecurityCheckScheduleConfigRequest
	GetStartTime() *int32
}

type ModifySecurityCheckScheduleConfigRequest struct {
	// The days on which the automatic configuration check runs. You can specify multiple days. Separate multiple days with commas (,). Valid values:
	//
	// 	- **1**: Monday
	//
	// 	- **2**: Tuesday
	//
	// 	- **3**: Wednesday
	//
	// 	- **4**: Thursday
	//
	// 	- **5**: Friday
	//
	// 	- **6**: Saturday
	//
	// 	- **7**: Sunday
	//
	// This parameter is required.
	//
	// example:
	//
	// 4,5,6
	DaysOfWeek *string `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty"`
	// The time period during which the automatic configuration check ends. Valid values:
	//
	// 	- **0**: 00:00 to 06:00
	//
	// 	- **6**: 06:00 to 12:00
	//
	// 	- **12**: 12:00 to 18:00
	//
	// 	- **18**: 18:00 to 24:00
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The time period during which the automatic configuration check starts. Valid values:
	//
	// 	- **0**: 00:00 to 06:00
	//
	// 	- **6**: 06:00 to 12:00
	//
	// 	- **12**: 12:00 to 18:00
	//
	// 	- **18**: 18:00 to 24:00
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifySecurityCheckScheduleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityCheckScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityCheckScheduleConfigRequest) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *ModifySecurityCheckScheduleConfigRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ModifySecurityCheckScheduleConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifySecurityCheckScheduleConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityCheckScheduleConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifySecurityCheckScheduleConfigRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetDaysOfWeek(v string) *ModifySecurityCheckScheduleConfigRequest {
	s.DaysOfWeek = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetEndTime(v int32) *ModifySecurityCheckScheduleConfigRequest {
	s.EndTime = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetLang(v string) *ModifySecurityCheckScheduleConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetResourceOwnerId(v int64) *ModifySecurityCheckScheduleConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetSourceIp(v string) *ModifySecurityCheckScheduleConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetStartTime(v int32) *ModifySecurityCheckScheduleConfigRequest {
	s.StartTime = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) Validate() error {
	return dara.Validate(s)
}
