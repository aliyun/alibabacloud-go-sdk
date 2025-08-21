// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallLogstashSystemPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *InstallLogstashSystemPluginRequest
	GetBody() *string
	SetClientToken(v string) *InstallLogstashSystemPluginRequest
	GetClientToken() *string
}

type InstallLogstashSystemPluginRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InstallLogstashSystemPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallLogstashSystemPluginRequest) GoString() string {
	return s.String()
}

func (s *InstallLogstashSystemPluginRequest) GetBody() *string {
	return s.Body
}

func (s *InstallLogstashSystemPluginRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InstallLogstashSystemPluginRequest) SetBody(v string) *InstallLogstashSystemPluginRequest {
	s.Body = &v
	return s
}

func (s *InstallLogstashSystemPluginRequest) SetClientToken(v string) *InstallLogstashSystemPluginRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallLogstashSystemPluginRequest) Validate() error {
	return dara.Validate(s)
}
