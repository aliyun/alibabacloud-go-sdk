// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallKibanaSystemPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallKibanaSystemPluginResponseBody
	GetRequestId() *string
	SetResult(v []*string) *InstallKibanaSystemPluginResponseBody
	GetResult() []*string
}

type InstallKibanaSystemPluginResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of plug-ins to be installed.
	Result []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallKibanaSystemPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallKibanaSystemPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallKibanaSystemPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallKibanaSystemPluginResponseBody) GetResult() []*string {
	return s.Result
}

func (s *InstallKibanaSystemPluginResponseBody) SetRequestId(v string) *InstallKibanaSystemPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallKibanaSystemPluginResponseBody) SetResult(v []*string) *InstallKibanaSystemPluginResponseBody {
	s.Result = v
	return s
}

func (s *InstallKibanaSystemPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
