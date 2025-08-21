// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallLogstashSystemPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallLogstashSystemPluginResponseBody
	GetRequestId() *string
	SetResult(v []*string) *InstallLogstashSystemPluginResponseBody
	GetResult() []*string
}

type InstallLogstashSystemPluginResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC4****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallLogstashSystemPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallLogstashSystemPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallLogstashSystemPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallLogstashSystemPluginResponseBody) GetResult() []*string {
	return s.Result
}

func (s *InstallLogstashSystemPluginResponseBody) SetRequestId(v string) *InstallLogstashSystemPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallLogstashSystemPluginResponseBody) SetResult(v []*string) *InstallLogstashSystemPluginResponseBody {
	s.Result = v
	return s
}

func (s *InstallLogstashSystemPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
