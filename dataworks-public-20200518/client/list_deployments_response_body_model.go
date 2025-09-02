// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDeploymentsResponseBodyData) *ListDeploymentsResponseBody
	GetData() *ListDeploymentsResponseBodyData
	SetRequestId(v string) *ListDeploymentsResponseBody
	GetRequestId() *string
}

type ListDeploymentsResponseBody struct {
	// The data returned.
	Data *ListDeploymentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBody) GetData() *ListDeploymentsResponseBodyData {
	return s.Data
}

func (s *ListDeploymentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentsResponseBody) SetData(v *ListDeploymentsResponseBodyData) *ListDeploymentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentsResponseBody) SetRequestId(v string) *ListDeploymentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeploymentsResponseBodyData struct {
	// The deployment packages.
	Deployments []*ListDeploymentsResponseBodyDataDeployments `json:"Deployments,omitempty" xml:"Deployments,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeploymentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyData) GetDeployments() []*ListDeploymentsResponseBodyDataDeployments {
	return s.Deployments
}

func (s *ListDeploymentsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDeploymentsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDeploymentsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDeploymentsResponseBodyData) SetDeployments(v []*ListDeploymentsResponseBodyDataDeployments) *ListDeploymentsResponseBodyData {
	s.Deployments = v
	return s
}

func (s *ListDeploymentsResponseBodyData) SetPageNumber(v int64) *ListDeploymentsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsResponseBodyData) SetPageSize(v int64) *ListDeploymentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsResponseBodyData) SetTotalCount(v int64) *ListDeploymentsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDeploymentsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDeploymentsResponseBodyDataDeployments struct {
	// The time when the deployment package was created.
	//
	// example:
	//
	// 1593877765000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who created the deployment package.
	//
	// example:
	//
	// 2003****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The error message returned when the deployment package failed.
	//
	// example:
	//
	// OK
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the deployment package was run. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1593877765000
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who ran the deployment package.
	//
	// example:
	//
	// 2003****
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The deployment package ID. You can call the [GetDeployment](https://help.aliyun.com/document_detail/173950.html) operation to obtain the ID.
	//
	// example:
	//
	// 11111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// auto_created
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// 	- 0: The deployment package is ready.
	//
	// 	- 1: The deployment package is deployed.
	//
	// 	- 2: The deployment package fails to be deployed.
	//
	// 	- 6: The deployment package is rejected.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsResponseBodyDataDeployments) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsResponseBodyDataDeployments) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetCreator() *string {
	return s.Creator
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetExecutor() *string {
	return s.Executor
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetId() *int64 {
	return s.Id
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetName() *string {
	return s.Name
}

func (s *ListDeploymentsResponseBodyDataDeployments) GetStatus() *int32 {
	return s.Status
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetCreateTime(v int64) *ListDeploymentsResponseBodyDataDeployments {
	s.CreateTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetCreator(v string) *ListDeploymentsResponseBodyDataDeployments {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetErrorMessage(v string) *ListDeploymentsResponseBodyDataDeployments {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetExecuteTime(v int64) *ListDeploymentsResponseBodyDataDeployments {
	s.ExecuteTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetExecutor(v string) *ListDeploymentsResponseBodyDataDeployments {
	s.Executor = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetId(v int64) *ListDeploymentsResponseBodyDataDeployments {
	s.Id = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetName(v string) *ListDeploymentsResponseBodyDataDeployments {
	s.Name = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) SetStatus(v int32) *ListDeploymentsResponseBodyDataDeployments {
	s.Status = &v
	return s
}

func (s *ListDeploymentsResponseBodyDataDeployments) Validate() error {
	return dara.Validate(s)
}
