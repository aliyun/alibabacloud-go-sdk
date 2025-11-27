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
	// The status of the automatic primary/secondary switchover feature. Valid values:
	//
	// 	- **Auto:*	- The automatic primary/secondary switchover feature is enabled. The system automatically switches your workloads over from the instance to its secondary instance in the event of a fault.
	//
	// 	- **Manual:*	- The automatic primary/secondary switchover feature is temporarily disabled.
	//
	// example:
	//
	// Manual
	HAConfig *string `json:"HAConfig,omitempty" xml:"HAConfig,omitempty"`
	// The time when the automatic primary/secondary switchover feature is enabled again. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-29T15:00:00Z
	ManualHATime *string `json:"ManualHATime,omitempty" xml:"ManualHATime,omitempty"`
	// The ID of the request.
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
