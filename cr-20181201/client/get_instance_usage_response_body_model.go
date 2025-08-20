// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChartNamespaceQuota(v string) *GetInstanceUsageResponseBody
	GetChartNamespaceQuota() *string
	SetChartNamespaceUsage(v string) *GetInstanceUsageResponseBody
	GetChartNamespaceUsage() *string
	SetChartRepoQuota(v string) *GetInstanceUsageResponseBody
	GetChartRepoQuota() *string
	SetChartRepoUsage(v string) *GetInstanceUsageResponseBody
	GetChartRepoUsage() *string
	SetCode(v string) *GetInstanceUsageResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *GetInstanceUsageResponseBody
	GetIsSuccess() *bool
	SetNamespaceQuota(v string) *GetInstanceUsageResponseBody
	GetNamespaceQuota() *string
	SetNamespaceUsage(v string) *GetInstanceUsageResponseBody
	GetNamespaceUsage() *string
	SetRepoQuota(v string) *GetInstanceUsageResponseBody
	GetRepoQuota() *string
	SetRepoUsage(v string) *GetInstanceUsageResponseBody
	GetRepoUsage() *string
	SetRequestId(v string) *GetInstanceUsageResponseBody
	GetRequestId() *string
	SetVpcQuota(v string) *GetInstanceUsageResponseBody
	GetVpcQuota() *string
	SetVpcUsage(v string) *GetInstanceUsageResponseBody
	GetVpcUsage() *string
}

type GetInstanceUsageResponseBody struct {
	// The quota of chart namespaces.
	//
	// example:
	//
	// 50
	ChartNamespaceQuota *string `json:"ChartNamespaceQuota,omitempty" xml:"ChartNamespaceQuota,omitempty"`
	// The number of chart namespaces that are created in the instance.
	//
	// example:
	//
	// 2
	ChartNamespaceUsage *string `json:"ChartNamespaceUsage,omitempty" xml:"ChartNamespaceUsage,omitempty"`
	// The quota of chart repositories for the instance.
	//
	// example:
	//
	// 5000
	ChartRepoQuota *string `json:"ChartRepoQuota,omitempty" xml:"ChartRepoQuota,omitempty"`
	// The number of chart repositories that are created.
	//
	// example:
	//
	// 5
	ChartRepoUsage *string `json:"ChartRepoUsage,omitempty" xml:"ChartRepoUsage,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The quota of image namespaces for the instance.
	//
	// example:
	//
	// 100
	NamespaceQuota *string `json:"NamespaceQuota,omitempty" xml:"NamespaceQuota,omitempty"`
	// The number of image namespaces that are created in the instance.
	//
	// example:
	//
	// 4
	NamespaceUsage *string `json:"NamespaceUsage,omitempty" xml:"NamespaceUsage,omitempty"`
	// The quota of image repositories for the instance.
	//
	// example:
	//
	// 1000
	RepoQuota *string `json:"RepoQuota,omitempty" xml:"RepoQuota,omitempty"`
	// The number of image repositories that are created in the instance.
	//
	// example:
	//
	// 2
	RepoUsage *string `json:"RepoUsage,omitempty" xml:"RepoUsage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A726E801-7FCF-43F9-AF1C-51B3E65D3E7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// VPC quota
	//
	// example:
	//
	// 5
	VpcQuota *string `json:"VpcQuota,omitempty" xml:"VpcQuota,omitempty"`
	// Number of bound VPCs
	//
	// example:
	//
	// 2
	VpcUsage *string `json:"VpcUsage,omitempty" xml:"VpcUsage,omitempty"`
}

func (s GetInstanceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceUsageResponseBody) GetChartNamespaceQuota() *string {
	return s.ChartNamespaceQuota
}

func (s *GetInstanceUsageResponseBody) GetChartNamespaceUsage() *string {
	return s.ChartNamespaceUsage
}

func (s *GetInstanceUsageResponseBody) GetChartRepoQuota() *string {
	return s.ChartRepoQuota
}

func (s *GetInstanceUsageResponseBody) GetChartRepoUsage() *string {
	return s.ChartRepoUsage
}

func (s *GetInstanceUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceUsageResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetInstanceUsageResponseBody) GetNamespaceQuota() *string {
	return s.NamespaceQuota
}

func (s *GetInstanceUsageResponseBody) GetNamespaceUsage() *string {
	return s.NamespaceUsage
}

func (s *GetInstanceUsageResponseBody) GetRepoQuota() *string {
	return s.RepoQuota
}

func (s *GetInstanceUsageResponseBody) GetRepoUsage() *string {
	return s.RepoUsage
}

func (s *GetInstanceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceUsageResponseBody) GetVpcQuota() *string {
	return s.VpcQuota
}

func (s *GetInstanceUsageResponseBody) GetVpcUsage() *string {
	return s.VpcUsage
}

func (s *GetInstanceUsageResponseBody) SetChartNamespaceQuota(v string) *GetInstanceUsageResponseBody {
	s.ChartNamespaceQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetChartNamespaceUsage(v string) *GetInstanceUsageResponseBody {
	s.ChartNamespaceUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetChartRepoQuota(v string) *GetInstanceUsageResponseBody {
	s.ChartRepoQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetChartRepoUsage(v string) *GetInstanceUsageResponseBody {
	s.ChartRepoUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetCode(v string) *GetInstanceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetIsSuccess(v bool) *GetInstanceUsageResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetNamespaceQuota(v string) *GetInstanceUsageResponseBody {
	s.NamespaceQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetNamespaceUsage(v string) *GetInstanceUsageResponseBody {
	s.NamespaceUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetRepoQuota(v string) *GetInstanceUsageResponseBody {
	s.RepoQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetRepoUsage(v string) *GetInstanceUsageResponseBody {
	s.RepoUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetRequestId(v string) *GetInstanceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetVpcQuota(v string) *GetInstanceUsageResponseBody {
	s.VpcQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetVpcUsage(v string) *GetInstanceUsageResponseBody {
	s.VpcUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
