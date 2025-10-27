// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpaClusterPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstallStatus(v []*CreateOpaClusterPluginResponseBodyInstallStatus) *CreateOpaClusterPluginResponseBody
	GetInstallStatus() []*CreateOpaClusterPluginResponseBodyInstallStatus
	SetRequestId(v string) *CreateOpaClusterPluginResponseBody
	GetRequestId() *string
}

type CreateOpaClusterPluginResponseBody struct {
	// The installation status of the components.
	InstallStatus []*CreateOpaClusterPluginResponseBodyInstallStatus `json:"InstallStatus,omitempty" xml:"InstallStatus,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0DC1F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOpaClusterPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaClusterPluginResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpaClusterPluginResponseBody) GetInstallStatus() []*CreateOpaClusterPluginResponseBodyInstallStatus {
	return s.InstallStatus
}

func (s *CreateOpaClusterPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpaClusterPluginResponseBody) SetInstallStatus(v []*CreateOpaClusterPluginResponseBodyInstallStatus) *CreateOpaClusterPluginResponseBody {
	s.InstallStatus = v
	return s
}

func (s *CreateOpaClusterPluginResponseBody) SetRequestId(v string) *CreateOpaClusterPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpaClusterPluginResponseBody) Validate() error {
	if s.InstallStatus != nil {
		for _, item := range s.InstallStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOpaClusterPluginResponseBodyInstallStatus struct {
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

func (s CreateOpaClusterPluginResponseBodyInstallStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaClusterPluginResponseBodyInstallStatus) GoString() string {
	return s.String()
}

func (s *CreateOpaClusterPluginResponseBodyInstallStatus) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateOpaClusterPluginResponseBodyInstallStatus) GetInstallStatus() *bool {
	return s.InstallStatus
}

func (s *CreateOpaClusterPluginResponseBodyInstallStatus) SetClusterId(v string) *CreateOpaClusterPluginResponseBodyInstallStatus {
	s.ClusterId = &v
	return s
}

func (s *CreateOpaClusterPluginResponseBodyInstallStatus) SetInstallStatus(v bool) *CreateOpaClusterPluginResponseBodyInstallStatus {
	s.InstallStatus = &v
	return s
}

func (s *CreateOpaClusterPluginResponseBodyInstallStatus) Validate() error {
	return dara.Validate(s)
}
