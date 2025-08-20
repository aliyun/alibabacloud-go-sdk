// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChartNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *GetChartNamespaceResponseBody
	GetAutoCreateRepo() *bool
	SetCode(v string) *GetChartNamespaceResponseBody
	GetCode() *string
	SetDefaultRepoType(v string) *GetChartNamespaceResponseBody
	GetDefaultRepoType() *string
	SetInstanceId(v string) *GetChartNamespaceResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetChartNamespaceResponseBody
	GetIsSuccess() *bool
	SetNamespaceId(v string) *GetChartNamespaceResponseBody
	GetNamespaceId() *string
	SetNamespaceName(v string) *GetChartNamespaceResponseBody
	GetNamespaceName() *string
	SetNamespaceStatus(v string) *GetChartNamespaceResponseBody
	GetNamespaceStatus() *string
	SetRequestId(v string) *GetChartNamespaceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetChartNamespaceResponseBody
	GetResourceGroupId() *string
}

type GetChartNamespaceResponseBody struct {
	// Indicates whether a repository was automatically created in the namespace. Valid values:
	//
	// 	- `true`: A repository was automatically created in the namespace.
	//
	// 	- `false`: No repository was automatically created in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository.
	//
	// 	- `PRIVATE`: a private repository.
	//
	// example:
	//
	// PRIVATE
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The ID of the namespace.
	//
	// example:
	//
	// crcn-43dhbjbyt2xl****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`: The namespace is normal.
	//
	// 	- `DELETING`: The namespace is being deleted.
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CD71CF13-93AA-4805-848B-69B2DD543A9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetChartNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetChartNamespaceResponseBody) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *GetChartNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChartNamespaceResponseBody) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *GetChartNamespaceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChartNamespaceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetChartNamespaceResponseBody) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetChartNamespaceResponseBody) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetChartNamespaceResponseBody) GetNamespaceStatus() *string {
	return s.NamespaceStatus
}

func (s *GetChartNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChartNamespaceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetChartNamespaceResponseBody) SetAutoCreateRepo(v bool) *GetChartNamespaceResponseBody {
	s.AutoCreateRepo = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetCode(v string) *GetChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetDefaultRepoType(v string) *GetChartNamespaceResponseBody {
	s.DefaultRepoType = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetInstanceId(v string) *GetChartNamespaceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetIsSuccess(v bool) *GetChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetNamespaceId(v string) *GetChartNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetNamespaceName(v string) *GetChartNamespaceResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetNamespaceStatus(v string) *GetChartNamespaceResponseBody {
	s.NamespaceStatus = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetRequestId(v string) *GetChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetResourceGroupId(v string) *GetChartNamespaceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetChartNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
