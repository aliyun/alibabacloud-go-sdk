// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRepo(v bool) *GetNamespaceResponseBody
	GetAutoCreateRepo() *bool
	SetCode(v string) *GetNamespaceResponseBody
	GetCode() *string
	SetDefaultRepoConfiguration(v *RepoConfiguration) *GetNamespaceResponseBody
	GetDefaultRepoConfiguration() *RepoConfiguration
	SetDefaultRepoType(v string) *GetNamespaceResponseBody
	GetDefaultRepoType() *string
	SetInstanceId(v string) *GetNamespaceResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetNamespaceResponseBody
	GetIsSuccess() *bool
	SetNamespaceId(v string) *GetNamespaceResponseBody
	GetNamespaceId() *string
	SetNamespaceName(v string) *GetNamespaceResponseBody
	GetNamespaceName() *string
	SetNamespaceStatus(v string) *GetNamespaceResponseBody
	GetNamespaceStatus() *string
	SetRequestId(v string) *GetNamespaceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetNamespaceResponseBody
	GetResourceGroupId() *string
}

type GetNamespaceResponseBody struct {
	// Indicates whether a repository is automatically created when an image is pushed to the namespace.
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
	Code                     *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of repositories in the namespace. Valid values:
	//
	// 	- PUBLIC: public repositories.
	//
	// 	- PRIVATE: private repositories.
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// crn-tiw8t3f8i5lt****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace.
	//
	// 	- NORMAL
	//
	// 	- DELETING
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E4BC9E21-8AA5-4582-83C1-C1209AB8196F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the namespace belongs.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNamespaceResponseBody) GetAutoCreateRepo() *bool {
	return s.AutoCreateRepo
}

func (s *GetNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNamespaceResponseBody) GetDefaultRepoConfiguration() *RepoConfiguration {
	return s.DefaultRepoConfiguration
}

func (s *GetNamespaceResponseBody) GetDefaultRepoType() *string {
	return s.DefaultRepoType
}

func (s *GetNamespaceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNamespaceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetNamespaceResponseBody) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNamespaceResponseBody) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetNamespaceResponseBody) GetNamespaceStatus() *string {
	return s.NamespaceStatus
}

func (s *GetNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNamespaceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetNamespaceResponseBody) SetAutoCreateRepo(v bool) *GetNamespaceResponseBody {
	s.AutoCreateRepo = &v
	return s
}

func (s *GetNamespaceResponseBody) SetCode(v string) *GetNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetNamespaceResponseBody) SetDefaultRepoConfiguration(v *RepoConfiguration) *GetNamespaceResponseBody {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *GetNamespaceResponseBody) SetDefaultRepoType(v string) *GetNamespaceResponseBody {
	s.DefaultRepoType = &v
	return s
}

func (s *GetNamespaceResponseBody) SetInstanceId(v string) *GetNamespaceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetNamespaceResponseBody) SetIsSuccess(v bool) *GetNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetNamespaceResponseBody) SetNamespaceId(v string) *GetNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *GetNamespaceResponseBody) SetNamespaceName(v string) *GetNamespaceResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetNamespaceResponseBody) SetNamespaceStatus(v string) *GetNamespaceResponseBody {
	s.NamespaceStatus = &v
	return s
}

func (s *GetNamespaceResponseBody) SetRequestId(v string) *GetNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNamespaceResponseBody) SetResourceGroupId(v string) *GetNamespaceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetNamespaceResponseBody) Validate() error {
	if s.DefaultRepoConfiguration != nil {
		if err := s.DefaultRepoConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
