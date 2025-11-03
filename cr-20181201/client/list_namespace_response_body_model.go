// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNamespaceResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListNamespaceResponseBody
	GetIsSuccess() *bool
	SetNamespaces(v []*ListNamespaceResponseBodyNamespaces) *ListNamespaceResponseBody
	GetNamespaces() []*ListNamespaceResponseBodyNamespaces
	SetPageNo(v int32) *ListNamespaceResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListNamespaceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNamespaceResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListNamespaceResponseBody
	GetTotalCount() *string
}

type ListNamespaceResponseBody struct {
	// The HTTP status code.
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
	// The queried namespaces.
	Namespaces []*ListNamespaceResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B7E5FCA5-55ED-451C-9649-0BB2B93387D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the queried namespaces.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNamespaceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListNamespaceResponseBody) GetNamespaces() []*ListNamespaceResponseBodyNamespaces {
	return s.Namespaces
}

func (s *ListNamespaceResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListNamespaceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespaceResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListNamespaceResponseBody) SetCode(v string) *ListNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespaceResponseBody) SetIsSuccess(v bool) *ListNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListNamespaceResponseBody) SetNamespaces(v []*ListNamespaceResponseBodyNamespaces) *ListNamespaceResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListNamespaceResponseBody) SetPageNo(v int32) *ListNamespaceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListNamespaceResponseBody) SetPageSize(v int32) *ListNamespaceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNamespaceResponseBody) SetRequestId(v string) *ListNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespaceResponseBody) SetTotalCount(v string) *ListNamespaceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNamespaceResponseBody) Validate() error {
	if s.Namespaces != nil {
		for _, item := range s.Namespaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespaceResponseBodyNamespaces struct {
	// Indicates whether the automatically creating repositories feature is enabled for the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo           *bool              `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of repositories in the namespace. Valid values:
	//
	// 	- `PUBLIC`: public repositories.
	//
	// 	- `PRIVATE`: private repositories.
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-94klsruryslx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// crn-tiw8t3f8i5lt****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The namespace name.
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
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm4n5kzyf2fbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNamespaceResponseBodyNamespaces) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListNamespaceResponseBodyNamespaces) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *ListNamespaceResponseBodyNamespaces) GetDefaultRepoConfiguration() *RepoConfiguration {
	return s.DefaultRepoConfiguration
}

func (s *ListNamespaceResponseBodyNamespaces) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *ListNamespaceResponseBodyNamespaces) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNamespaceResponseBodyNamespaces) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamespaceResponseBodyNamespaces) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListNamespaceResponseBodyNamespaces) GetNamespaceStatus() *string {
	return s.NamespaceStatus
}

func (s *ListNamespaceResponseBodyNamespaces) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNamespaceResponseBodyNamespaces) SetAutoCreateRepo(v bool) *ListNamespaceResponseBodyNamespaces {
	s.AutoCreateRepo = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetDefaultRepoConfiguration(v *RepoConfiguration) *ListNamespaceResponseBodyNamespaces {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetDefaultRepoType(v string) *ListNamespaceResponseBodyNamespaces {
	s.DefaultRepoType = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetInstanceId(v string) *ListNamespaceResponseBodyNamespaces {
	s.InstanceId = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetNamespaceId(v string) *ListNamespaceResponseBodyNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetNamespaceName(v string) *ListNamespaceResponseBodyNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetNamespaceStatus(v string) *ListNamespaceResponseBodyNamespaces {
	s.NamespaceStatus = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetResourceGroupId(v string) *ListNamespaceResponseBodyNamespaces {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) Validate() error {
	if s.DefaultRepoConfiguration != nil {
		if err := s.DefaultRepoConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
