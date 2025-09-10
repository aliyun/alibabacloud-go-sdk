// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusDashboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *ListPrometheusDashboardsRequest
	GetAliyunLang() *string
	SetResourceGroupId(v string) *ListPrometheusDashboardsRequest
	GetResourceGroupId() *string
}

type ListPrometheusDashboardsRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListPrometheusDashboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusDashboardsRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *ListPrometheusDashboardsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusDashboardsRequest) SetAliyunLang(v string) *ListPrometheusDashboardsRequest {
	s.AliyunLang = &v
	return s
}

func (s *ListPrometheusDashboardsRequest) SetResourceGroupId(v string) *ListPrometheusDashboardsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusDashboardsRequest) Validate() error {
	return dara.Validate(s)
}
