// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetPrometheusInstanceRequest
	GetAliyunLang() *string
	SetResourceGroupId(v string) *GetPrometheusInstanceRequest
	GetResourceGroupId() *string
}

type GetPrometheusInstanceRequest struct {
	// The language of the response. Valid values: \\`zh\\` (Chinese) and \\`en\\` (English). Default value: \\`zh\\`.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bhocin5e2na
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s GetPrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetPrometheusInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetPrometheusInstanceRequest) SetAliyunLang(v string) *GetPrometheusInstanceRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetPrometheusInstanceRequest) SetResourceGroupId(v string) *GetPrometheusInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetPrometheusInstanceRequest) Validate() error {
	return dara.Validate(s)
}
