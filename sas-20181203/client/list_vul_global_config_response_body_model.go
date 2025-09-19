// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulGlobalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVulGlobalConfigResponseBody
	GetRequestId() *string
	SetVulGlobalConfigList(v []*ListVulGlobalConfigResponseBodyVulGlobalConfigList) *ListVulGlobalConfigResponseBody
	GetVulGlobalConfigList() []*ListVulGlobalConfigResponseBodyVulGlobalConfigList
}

type ListVulGlobalConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations.
	VulGlobalConfigList []*ListVulGlobalConfigResponseBodyVulGlobalConfigList `json:"VulGlobalConfigList,omitempty" xml:"VulGlobalConfigList,omitempty" type:"Repeated"`
}

func (s ListVulGlobalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVulGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListVulGlobalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVulGlobalConfigResponseBody) GetVulGlobalConfigList() []*ListVulGlobalConfigResponseBodyVulGlobalConfigList {
	return s.VulGlobalConfigList
}

func (s *ListVulGlobalConfigResponseBody) SetRequestId(v string) *ListVulGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVulGlobalConfigResponseBody) SetVulGlobalConfigList(v []*ListVulGlobalConfigResponseBodyVulGlobalConfigList) *ListVulGlobalConfigResponseBody {
	s.VulGlobalConfigList = v
	return s
}

func (s *ListVulGlobalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVulGlobalConfigResponseBodyVulGlobalConfigList struct {
	// The key of the configuration item.
	//
	// example:
	//
	// vul_scan_ip_list
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// 127.0.*.*,127.0.*.*
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
}

func (s ListVulGlobalConfigResponseBodyVulGlobalConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListVulGlobalConfigResponseBodyVulGlobalConfigList) GoString() string {
	return s.String()
}

func (s *ListVulGlobalConfigResponseBodyVulGlobalConfigList) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ListVulGlobalConfigResponseBodyVulGlobalConfigList) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ListVulGlobalConfigResponseBodyVulGlobalConfigList) SetConfigKey(v string) *ListVulGlobalConfigResponseBodyVulGlobalConfigList {
	s.ConfigKey = &v
	return s
}

func (s *ListVulGlobalConfigResponseBodyVulGlobalConfigList) SetConfigValue(v string) *ListVulGlobalConfigResponseBodyVulGlobalConfigList {
	s.ConfigValue = &v
	return s
}

func (s *ListVulGlobalConfigResponseBodyVulGlobalConfigList) Validate() error {
	return dara.Validate(s)
}
