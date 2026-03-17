// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceAutoUpgradePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetCronExpression() *string
	SetDuration(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetDuration() *string
	SetJitter(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetJitter() *string
	SetRequestId(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetSerialNumber() *string
	SetSmartAGId(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetSmartAGId() *string
	SetTimeZone(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetTimeZone() *string
	SetUpgradeType(v string) *DescribeDeviceAutoUpgradePolicyResponseBody
	GetUpgradeType() *string
}

type DescribeDeviceAutoUpgradePolicyResponseBody struct {
	// The time when upgrades start. The time was configured by using a cron expression.
	//
	// Example value: `0 0 4 1/1 	- ?`. The example value indicates that upgrades start at 04:00:00 (UTC+8) on the first day of each month.
	//
	// example:
	//
	// 0 0 4 1/1 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The time period during which upgrades are performed.
	//
	// Valid values: **30 to 120**.
	//
	// Unit: minutes.
	//
	// example:
	//
	// 60
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time differences between upgrades. Unit: minutes.
	//
	// example:
	//
	// 5
	Jitter *string `json:"Jitter,omitempty" xml:"Jitter,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0E20FBB8-BCFC-4F5E-BD94-77FF6A2133D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the SAG instance.
	//
	// example:
	//
	// sage62x022502****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-kxe2cv7hot7qrv****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The time zone in which the trigger period takes effect.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The update type. Valid values:
	//
	// 	- **scheduled**: scheduled upgrade.
	//
	// 	- **boot**: automatic upgrade upon instance startup.
	//
	// 	- **manual**: manual upgrade.
	//
	// example:
	//
	// scheduled
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s DescribeDeviceAutoUpgradePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceAutoUpgradePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetDuration() *string {
	return s.Duration
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetJitter() *string {
	return s.Jitter
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetCronExpression(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.CronExpression = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetDuration(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetJitter(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.Jitter = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetRequestId(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetSerialNumber(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetSmartAGId(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.SmartAGId = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetTimeZone(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.TimeZone = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) SetUpgradeType(v string) *DescribeDeviceAutoUpgradePolicyResponseBody {
	s.UpgradeType = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
