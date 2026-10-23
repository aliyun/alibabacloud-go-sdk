// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UninstallPluginRequest
	GetBody() *string
	SetClientToken(v string) *UninstallPluginRequest
	GetClientToken() *string
	SetForce(v bool) *UninstallPluginRequest
	GetForce() *bool
}

type UninstallPluginRequest struct {
	// The names of the system plug-ins to uninstall.
	//
	// example:
	//
	// ["aliyun-sql","codec-compression"]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to forcibly uninstall the plug-ins. You can use this parameter to cancel the installation of plug-ins that are being installed. This parameter is only applicable to instances that use the new architecture. Default value: false.
	//
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s UninstallPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallPluginRequest) GoString() string {
	return s.String()
}

func (s *UninstallPluginRequest) GetBody() *string {
	return s.Body
}

func (s *UninstallPluginRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UninstallPluginRequest) GetForce() *bool {
	return s.Force
}

func (s *UninstallPluginRequest) SetBody(v string) *UninstallPluginRequest {
	s.Body = &v
	return s
}

func (s *UninstallPluginRequest) SetClientToken(v string) *UninstallPluginRequest {
	s.ClientToken = &v
	return s
}

func (s *UninstallPluginRequest) SetForce(v bool) *UninstallPluginRequest {
	s.Force = &v
	return s
}

func (s *UninstallPluginRequest) Validate() error {
	return dara.Validate(s)
}
