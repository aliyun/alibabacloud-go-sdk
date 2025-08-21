// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallLogstashPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*string) *UninstallLogstashPluginRequest
	GetBody() []*string
	SetClientToken(v string) *UninstallLogstashPluginRequest
	GetClientToken() *string
}

type UninstallLogstashPluginRequest struct {
	// example:
	//
	// ["logstash-input-datahub", "logstash-input-maxcompute" ]
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UninstallLogstashPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallLogstashPluginRequest) GoString() string {
	return s.String()
}

func (s *UninstallLogstashPluginRequest) GetBody() []*string {
	return s.Body
}

func (s *UninstallLogstashPluginRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UninstallLogstashPluginRequest) SetBody(v []*string) *UninstallLogstashPluginRequest {
	s.Body = v
	return s
}

func (s *UninstallLogstashPluginRequest) SetClientToken(v string) *UninstallLogstashPluginRequest {
	s.ClientToken = &v
	return s
}

func (s *UninstallLogstashPluginRequest) Validate() error {
	return dara.Validate(s)
}
