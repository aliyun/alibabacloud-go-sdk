// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaPluginStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstallStatus(v []*GetOpaPluginStatusResponseBodyInstallStatus) *GetOpaPluginStatusResponseBody
	GetInstallStatus() []*GetOpaPluginStatusResponseBodyInstallStatus
	SetRequestId(v string) *GetOpaPluginStatusResponseBody
	GetRequestId() *string
}

type GetOpaPluginStatusResponseBody struct {
	// The installation status of the components that are required for clusters protected by proactive defense for containers.
	InstallStatus []*GetOpaPluginStatusResponseBodyInstallStatus `json:"InstallStatus,omitempty" xml:"InstallStatus,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8BF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpaPluginStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpaPluginStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpaPluginStatusResponseBody) GetInstallStatus() []*GetOpaPluginStatusResponseBodyInstallStatus {
	return s.InstallStatus
}

func (s *GetOpaPluginStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpaPluginStatusResponseBody) SetInstallStatus(v []*GetOpaPluginStatusResponseBodyInstallStatus) *GetOpaPluginStatusResponseBody {
	s.InstallStatus = v
	return s
}

func (s *GetOpaPluginStatusResponseBody) SetRequestId(v string) *GetOpaPluginStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpaPluginStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOpaPluginStatusResponseBodyInstallStatus struct {
	// The cluster ID.
	//
	// example:
	//
	// c60b77fe62093480db6164a3c2fa****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Indicates whether the component is installed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	InstallStatus *bool `json:"InstallStatus,omitempty" xml:"InstallStatus,omitempty"`
}

func (s GetOpaPluginStatusResponseBodyInstallStatus) String() string {
	return dara.Prettify(s)
}

func (s GetOpaPluginStatusResponseBodyInstallStatus) GoString() string {
	return s.String()
}

func (s *GetOpaPluginStatusResponseBodyInstallStatus) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetOpaPluginStatusResponseBodyInstallStatus) GetInstallStatus() *bool {
	return s.InstallStatus
}

func (s *GetOpaPluginStatusResponseBodyInstallStatus) SetClusterId(v string) *GetOpaPluginStatusResponseBodyInstallStatus {
	s.ClusterId = &v
	return s
}

func (s *GetOpaPluginStatusResponseBodyInstallStatus) SetInstallStatus(v bool) *GetOpaPluginStatusResponseBodyInstallStatus {
	s.InstallStatus = &v
	return s
}

func (s *GetOpaPluginStatusResponseBodyInstallStatus) Validate() error {
	return dara.Validate(s)
}
