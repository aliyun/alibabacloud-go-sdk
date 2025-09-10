// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetPrometheusViewRequest
	GetAliyunLang() *string
	SetResourceGroupId(v string) *GetPrometheusViewRequest
	GetResourceGroupId() *string
}

type GetPrometheusViewRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// example:
	//
	// rg-aek2bhocin5e2na
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s GetPrometheusViewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusViewRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusViewRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetPrometheusViewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetPrometheusViewRequest) SetAliyunLang(v string) *GetPrometheusViewRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetPrometheusViewRequest) SetResourceGroupId(v string) *GetPrometheusViewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetPrometheusViewRequest) Validate() error {
	return dara.Validate(s)
}
