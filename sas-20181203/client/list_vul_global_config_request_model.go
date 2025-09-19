// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulGlobalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *ListVulGlobalConfigRequest
	GetConfigKey() *string
}

type ListVulGlobalConfigRequest struct {
	// The key of the configuration item. Valid values:
	//
	// 	- **vul_scan_ip_list**: The IP addresses that are detected.
	//
	// example:
	//
	// vul_scan_ip_list
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
}

func (s ListVulGlobalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVulGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *ListVulGlobalConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ListVulGlobalConfigRequest) SetConfigKey(v string) *ListVulGlobalConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *ListVulGlobalConfigRequest) Validate() error {
	return dara.Validate(s)
}
