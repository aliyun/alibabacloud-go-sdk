// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityCheckScheduleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityCheckScheduleConfigResponseBody
	GetRequestId() *string
	SetRiskCheckJobConfig(v *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) *DescribeSecurityCheckScheduleConfigResponseBody
	GetRiskCheckJobConfig() *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig
}

type DescribeSecurityCheckScheduleConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 48D2E9A9-A1B0-4295-B727-0995757C47E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of custom check tasks.
	RiskCheckJobConfig *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig `json:"RiskCheckJobConfig,omitempty" xml:"RiskCheckJobConfig,omitempty" type:"Struct"`
}

func (s DescribeSecurityCheckScheduleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityCheckScheduleConfigResponseBody) GetRiskCheckJobConfig() *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig {
	return s.RiskCheckJobConfig
}

func (s *DescribeSecurityCheckScheduleConfigResponseBody) SetRequestId(v string) *DescribeSecurityCheckScheduleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBody) SetRiskCheckJobConfig(v *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) *DescribeSecurityCheckScheduleConfigResponseBody {
	s.RiskCheckJobConfig = v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig struct {
	// The day of the week when the check tasks are performed. Multiple days can be specified. Multiple days are separated by commas (,).
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
	// example:
	//
	// 1,2,3
	DaysOfWeek *string `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty"`
	// The time range during which check tasks end. Valid values:
	//
	// 	- **6**: 00:00 to 06:00
	//
	// 	- **12**: 06:00 to 12:00
	//
	// 	- **18**: 12:00 to 18:00
	//
	// 	- **24**: 18:00 to 24:00
	//
	// example:
	//
	// 12
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time range during which check tasks start. Valid values:
	//
	// 	- **0**: 00:00 to 06:00
	//
	// 	- **6**: 06:00 to 12:00
	//
	// 	- **12**: 12:00 to 18:00
	//
	// 	- **18**: 18:00 to 24:00
	//
	// example:
	//
	// 6
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) GetEndTime() *int32 {
	return s.EndTime
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) GetStartTime() *int32 {
	return s.StartTime
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) SetDaysOfWeek(v string) *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig {
	s.DaysOfWeek = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) SetEndTime(v int32) *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) SetStartTime(v int32) *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) Validate() error {
	return dara.Validate(s)
}
