// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSystemPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallSystemPluginResponseBody
	GetRequestId() *string
	SetResult(v []*string) *InstallSystemPluginResponseBody
	GetResult() []*string
}

type InstallSystemPluginResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of plug-ins to be installed.
	Result []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallSystemPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallSystemPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSystemPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallSystemPluginResponseBody) GetResult() []*string {
	return s.Result
}

func (s *InstallSystemPluginResponseBody) SetRequestId(v string) *InstallSystemPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallSystemPluginResponseBody) SetResult(v []*string) *InstallSystemPluginResponseBody {
	s.Result = v
	return s
}

func (s *InstallSystemPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
