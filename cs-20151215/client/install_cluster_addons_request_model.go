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
	// 请求体参数。
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
	return dara.Validate(s)
}

type InstallClusterAddonsRequestBody struct {
	// 组件自定义参数，使用JSON字符串编码。
	//
	// example:
	//
	// {\\"IngressDashboardEnabled\\":\\"true\\",\\"sls_project_name\\":\\"your_sls_project_name\\"}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// 组件名称。您可以通过[ListAddons](https://help.aliyun.com/document_detail/2667939.html)接口查询可用组件的信息，包括组件名称及版本等。
	//
	// This parameter is required.
	//
	// example:
	//
	// logtail-ds
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 组件版本。您可以通过[ListAddons](https://help.aliyun.com/document_detail/2667939.html)接口查询可用组件的信息，包括组件名称及版本等。
	//
	// This parameter is required.
	//
	// example:
	//
	// v1.7.3.0-aliyun
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
