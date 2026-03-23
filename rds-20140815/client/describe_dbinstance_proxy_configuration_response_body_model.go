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
	AttacksProtectionConfiguration     *string `json:"AttacksProtectionConfiguration,omitempty" xml:"AttacksProtectionConfiguration,omitempty"`
	PersistentConnectionsConfiguration *string `json:"PersistentConnectionsConfiguration,omitempty" xml:"PersistentConnectionsConfiguration,omitempty"`
	RequestId                          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransparentSwitchConfiguration     *string `json:"TransparentSwitchConfiguration,omitempty" xml:"TransparentSwitchConfiguration,omitempty"`
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
