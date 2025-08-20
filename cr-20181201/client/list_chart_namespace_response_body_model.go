// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChartNamespaceResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListChartNamespaceResponseBody
	GetIsSuccess() *bool
	SetNamespaces(v []*ListChartNamespaceResponseBodyNamespaces) *ListChartNamespaceResponseBody
	GetNamespaces() []*ListChartNamespaceResponseBodyNamespaces
	SetPageNo(v int32) *ListChartNamespaceResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListChartNamespaceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListChartNamespaceResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListChartNamespaceResponseBody
	GetTotalCount() *string
}

type ListChartNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The namespaces.
	Namespaces []*ListChartNamespaceResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F56D589D-AF7F-4900-BA46-62C780AC2C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChartNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChartNamespaceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListChartNamespaceResponseBody) GetNamespaces() []*ListChartNamespaceResponseBodyNamespaces {
	return s.Namespaces
}

func (s *ListChartNamespaceResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChartNamespaceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChartNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChartNamespaceResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListChartNamespaceResponseBody) SetCode(v string) *ListChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetIsSuccess(v bool) *ListChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetNamespaces(v []*ListChartNamespaceResponseBodyNamespaces) *ListChartNamespaceResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListChartNamespaceResponseBody) SetPageNo(v int32) *ListChartNamespaceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetPageSize(v int32) *ListChartNamespaceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetRequestId(v string) *ListChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetTotalCount(v string) *ListChartNamespaceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChartNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListChartNamespaceResponseBodyNamespaces struct {
	// Indicates whether a repository was automatically created when a chart is pushed to the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository
	//
	// 	- `PRIVATE`: a private repository
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// null
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
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
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfm4n5kzyf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListChartNamespaceResponseBodyNamespaces) String() string {
	return dara.Prettify(s)
}

func (s ListChartNamespaceResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceResponseBodyNamespaces) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *ListChartNamespaceResponseBodyNamespaces) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *ListChartNamespaceResponseBodyNamespaces) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChartNamespaceResponseBodyNamespaces) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListChartNamespaceResponseBodyNamespaces) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListChartNamespaceResponseBodyNamespaces) GetNamespaceStatus() *string {
	return s.NamespaceStatus
}

func (s *ListChartNamespaceResponseBodyNamespaces) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetAutoCreateRepo(v bool) *ListChartNamespaceResponseBodyNamespaces {
	s.AutoCreateRepo = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetDefaultRepoType(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.DefaultRepoType = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetInstanceId(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.InstanceId = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetNamespaceId(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetNamespaceName(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetNamespaceStatus(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.NamespaceStatus = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetResourceGroupId(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.ResourceGroupId = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) Validate() error {
	return dara.Validate(s)
}
