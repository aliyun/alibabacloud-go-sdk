// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallClusterAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*InstallClusterAddonsRequestBody) *InstallClusterAddonsRequest
	GetBody() []*InstallClusterAddonsRequestBody
}

type InstallClusterAddonsRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// ags-metrics-collector
	Body []*InstallClusterAddonsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s InstallClusterAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequest) GetBody() []*InstallClusterAddonsRequestBody {
	return s.Body
}

func (s *InstallClusterAddonsRequest) SetBody(v []*InstallClusterAddonsRequestBody) *InstallClusterAddonsRequest {
	s.Body = v
	return s
}

func (s *InstallClusterAddonsRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallClusterAddonsRequestBody struct {
	// The custom parameters of the component, encoded as a JSON string.
	//
	// example:
	//
	// {\\"IngressDashboardEnabled\\":\\"true\\",\\"sls_project_name\\":\\"your_sls_project_name\\"}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The component name. You can call the [ListAddons](https://help.aliyun.com/document_detail/2667939.html) operation to query information about available components, including component names and versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// storage-operato
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The component version. You can call the [ListAddons](https://help.aliyun.com/document_detail/2667939.html) operation to query information about available components, including component names and versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1.32.9
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s InstallClusterAddonsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsRequestBody) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequestBody) GetConfig() *string {
	return s.Config
}

func (s *InstallClusterAddonsRequestBody) GetName() *string {
	return s.Name
}

func (s *InstallClusterAddonsRequestBody) GetVersion() *string {
	return s.Version
}

func (s *InstallClusterAddonsRequestBody) SetConfig(v string) *InstallClusterAddonsRequestBody {
	s.Config = &v
	return s
}

func (s *InstallClusterAddonsRequestBody) SetName(v string) *InstallClusterAddonsRequestBody {
	s.Name = &v
	return s
}

func (s *InstallClusterAddonsRequestBody) SetVersion(v string) *InstallClusterAddonsRequestBody {
	s.Version = &v
	return s
}

func (s *InstallClusterAddonsRequestBody) Validate() error {
	return dara.Validate(s)
}
