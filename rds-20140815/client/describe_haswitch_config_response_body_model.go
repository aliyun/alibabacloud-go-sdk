// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHASwitchConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHAConfig(v string) *DescribeHASwitchConfigResponseBody
	GetHAConfig() *string
	SetManualHATime(v string) *DescribeHASwitchConfigResponseBody
	GetManualHATime() *string
	SetRequestId(v string) *DescribeHASwitchConfigResponseBody
	GetRequestId() *string
}

type DescribeHASwitchConfigResponseBody struct {
	// The automatic primary/secondary switchover setting. Valid values:
	//
	// 	- **Auto**: The system automatically switches over between the primary and secondary instances upon a fault.
	//
	// 	- **Manual**: Automatic switchover has been temporarily disabled.
	//
	// example:
	//
	// Manual
	HAConfig *string `json:"HAConfig,omitempty" xml:"HAConfig,omitempty"`
	// The deadline for the temporary disabling of automatic switchover. The time follows the ISO 8601 standard in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-29T15:00:00Z
	ManualHATime *string `json:"ManualHATime,omitempty" xml:"ManualHATime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4FDF4B79-2741-4C5F-8C76-4B953FC5C2B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHASwitchConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHASwitchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHASwitchConfigResponseBody) GetHAConfig() *string {
	return s.HAConfig
}

func (s *DescribeHASwitchConfigResponseBody) GetManualHATime() *string {
	return s.ManualHATime
}

func (s *DescribeHASwitchConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHASwitchConfigResponseBody) SetHAConfig(v string) *DescribeHASwitchConfigResponseBody {
	s.HAConfig = &v
	return s
}

func (s *DescribeHASwitchConfigResponseBody) SetManualHATime(v string) *DescribeHASwitchConfigResponseBody {
	s.ManualHATime = &v
	return s
}

func (s *DescribeHASwitchConfigResponseBody) SetRequestId(v string) *DescribeHASwitchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHASwitchConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
