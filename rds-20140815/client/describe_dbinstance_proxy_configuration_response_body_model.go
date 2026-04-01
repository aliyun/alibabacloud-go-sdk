// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceProxyConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttacksProtectionConfiguration(v string) *DescribeDBInstanceProxyConfigurationResponseBody
	GetAttacksProtectionConfiguration() *string
	SetPersistentConnectionsConfiguration(v string) *DescribeDBInstanceProxyConfigurationResponseBody
	GetPersistentConnectionsConfiguration() *string
	SetRequestId(v string) *DescribeDBInstanceProxyConfigurationResponseBody
	GetRequestId() *string
	SetTransparentSwitchConfiguration(v string) *DescribeDBInstanceProxyConfigurationResponseBody
	GetTransparentSwitchConfiguration() *string
}

type DescribeDBInstanceProxyConfigurationResponseBody struct {
	// Indicates whether the mechanism that is used to mitigate brute-force attacks is enabled:
	//
	// 	- **Enable**
	//
	// 	- **Disable**
	//
	// The return value is a JSON string. Example:
	//
	//     {"status":"Disable", "check_interval_seconds": 60,
	//
	//               "max_failed_login_attempts": 60, "blocking_seconds": 600}
	//
	// Description:
	//
	// 	- Each client allows {max_failed_login_attempts} logon attempts that fail due to incorrect passwords within {check_interval_seconds} seconds. If one more such attempt is conducted, the client must wait for {blocking_seconds} seconds before you can try again.
	//
	// 	- Valid values:
	//
	//     	- check_interval_seconds: **30 to 600**. Unit: seconds.
	//
	//     	- max_failed_login_attempts: **10 to 5000**. Unit: times.
	//
	//     	- blocking_seconds: **30 to 3600**. Unit: seconds.
	//
	// example:
	//
	// {\\"check_interval_seconds\\":\\"0\\",\\"max_failed_login_attempts\\":\\"0\\",\\"blocking_seconds\\":\\"0\\",\\"status\\":\\"Disable\\"}
	AttacksProtectionConfiguration *string `json:"AttacksProtectionConfiguration,omitempty" xml:"AttacksProtectionConfiguration,omitempty"`
	// Indicates whether the short-lived connection optimization feature is enabled.
	//
	// 	- **Enable**
	//
	// 	- **Disable**
	//
	// In this case, the return value is a JSON string. Examples:
	//
	//     {"status":"Disable"}.
	//
	// example:
	//
	// {\\"status\\":\\"Disable\\"}
	PersistentConnectionsConfiguration *string `json:"PersistentConnectionsConfiguration,omitempty" xml:"PersistentConnectionsConfiguration,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9DD55F4-1A5F-48CA-BA57-DFB3CA8C4C34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the transparent switchover feature is enabled.
	//
	// 	- **Enable**
	//
	// 	- **Disable**
	//
	// The return value is a JSON string. Example:
	//
	//     {"status":"Enable"}
	//
	// example:
	//
	// {\\"status\\":\\"Enable\\"}
	TransparentSwitchConfiguration *string `json:"TransparentSwitchConfiguration,omitempty" xml:"TransparentSwitchConfiguration,omitempty"`
}

func (s DescribeDBInstanceProxyConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceProxyConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) GetAttacksProtectionConfiguration() *string {
	return s.AttacksProtectionConfiguration
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) GetPersistentConnectionsConfiguration() *string {
	return s.PersistentConnectionsConfiguration
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) GetTransparentSwitchConfiguration() *string {
	return s.TransparentSwitchConfiguration
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) SetAttacksProtectionConfiguration(v string) *DescribeDBInstanceProxyConfigurationResponseBody {
	s.AttacksProtectionConfiguration = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) SetPersistentConnectionsConfiguration(v string) *DescribeDBInstanceProxyConfigurationResponseBody {
	s.PersistentConnectionsConfiguration = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) SetRequestId(v string) *DescribeDBInstanceProxyConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) SetTransparentSwitchConfiguration(v string) *DescribeDBInstanceProxyConfigurationResponseBody {
	s.TransparentSwitchConfiguration = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
