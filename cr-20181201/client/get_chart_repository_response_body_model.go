// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChartRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetChartRepositoryResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetChartRepositoryResponseBody
	GetCreateTime() *int64
	SetInstanceId(v string) *GetChartRepositoryResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetChartRepositoryResponseBody
	GetIsSuccess() *bool
	SetModifiedTime(v int64) *GetChartRepositoryResponseBody
	GetModifiedTime() *int64
	SetRepoId(v string) *GetChartRepositoryResponseBody
	GetRepoId() *string
	SetRepoName(v string) *GetChartRepositoryResponseBody
	GetRepoName() *string
	SetRepoNamespaceName(v string) *GetChartRepositoryResponseBody
	GetRepoNamespaceName() *string
	SetRepoStatus(v string) *GetChartRepositoryResponseBody
	GetRepoStatus() *string
	SetRepoType(v string) *GetChartRepositoryResponseBody
	GetRepoType() *string
	SetRequestId(v string) *GetChartRepositoryResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetChartRepositoryResponseBody
	GetResourceGroupId() *string
	SetSummary(v string) *GetChartRepositoryResponseBody
	GetSummary() *string
}

type GetChartRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the chart repository was created.
	//
	// example:
	//
	// 1563767620000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
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
	// The time when the chart repository was last modified.
	//
	// example:
	//
	// 1563767700000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the chart repository.
	//
	// example:
	//
	// crcr-c7letfwev5oq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the chart repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the chart repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the chart repository. Valid values:
	//
	// 	- `NORMAL`: The repository is normal.
	//
	// 	- `DELETING`: The repository is being deleted.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the chart repository. Valid values:
	//
	// 	- `PUBLIC`: a public repository
	//
	// 	- `PRIVATE`: a private repository
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3F6AB56-DEF4-4FF5-8DE4-680362C0E21F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary about the chart repository.
	//
	// example:
	//
	// test
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetChartRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetChartRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChartRepositoryResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetChartRepositoryResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChartRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetChartRepositoryResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetChartRepositoryResponseBody) GetRepoId() *string {
	return s.RepoId
}

func (s *GetChartRepositoryResponseBody) GetRepoName() *string {
	return s.RepoName
}

func (s *GetChartRepositoryResponseBody) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *GetChartRepositoryResponseBody) GetRepoStatus() *string {
	return s.RepoStatus
}

func (s *GetChartRepositoryResponseBody) GetRepoType() *string {
	return s.RepoType
}

func (s *GetChartRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChartRepositoryResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetChartRepositoryResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *GetChartRepositoryResponseBody) SetCode(v string) *GetChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetCreateTime(v int64) *GetChartRepositoryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetInstanceId(v string) *GetChartRepositoryResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetIsSuccess(v bool) *GetChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetModifiedTime(v int64) *GetChartRepositoryResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoId(v string) *GetChartRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoName(v string) *GetChartRepositoryResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoNamespaceName(v string) *GetChartRepositoryResponseBody {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoStatus(v string) *GetChartRepositoryResponseBody {
	s.RepoStatus = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoType(v string) *GetChartRepositoryResponseBody {
	s.RepoType = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRequestId(v string) *GetChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetResourceGroupId(v string) *GetChartRepositoryResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetSummary(v string) *GetChartRepositoryResponseBody {
	s.Summary = &v
	return s
}

func (s *GetChartRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
