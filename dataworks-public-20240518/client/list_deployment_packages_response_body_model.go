// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDeploymentPackagesResponseBodyData) *ListDeploymentPackagesResponseBody
	GetData() *ListDeploymentPackagesResponseBodyData
	SetRequestId(v string) *ListDeploymentPackagesResponseBody
	GetRequestId() *string
}

type ListDeploymentPackagesResponseBody struct {
	// The list of deployment packages that meet the query conditions.
	Data *ListDeploymentPackagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 952795279527ab****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackagesResponseBody) GetData() *ListDeploymentPackagesResponseBodyData {
	return s.Data
}

func (s *ListDeploymentPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentPackagesResponseBody) SetData(v *ListDeploymentPackagesResponseBodyData) *ListDeploymentPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentPackagesResponseBody) SetRequestId(v string) *ListDeploymentPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentPackagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeploymentPackagesResponseBodyData struct {
	// The returned list of deployment packages.
	Deployments []*ListDeploymentPackagesResponseBodyDataDeployments `json:"Deployments,omitempty" xml:"Deployments,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that meet the conditions.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeploymentPackagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackagesResponseBodyData) GetDeployments() []*ListDeploymentPackagesResponseBodyDataDeployments {
	return s.Deployments
}

func (s *ListDeploymentPackagesResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDeploymentPackagesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDeploymentPackagesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDeploymentPackagesResponseBodyData) SetDeployments(v []*ListDeploymentPackagesResponseBodyDataDeployments) *ListDeploymentPackagesResponseBodyData {
	s.Deployments = v
	return s
}

func (s *ListDeploymentPackagesResponseBodyData) SetPageNumber(v int64) *ListDeploymentPackagesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyData) SetPageSize(v int64) *ListDeploymentPackagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyData) SetTotalCount(v int64) *ListDeploymentPackagesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyData) Validate() error {
	if s.Deployments != nil {
		for _, item := range s.Deployments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeploymentPackagesResponseBodyDataDeployments struct {
	// The timestamp when the deployment package was created.
	//
	// example:
	//
	// 1593877765000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account ID of the deployment package creator.
	//
	// example:
	//
	// 2003****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// When the deployment package fails to execute, this parameter is used to record the error message.
	//
	// example:
	//
	// OK
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The timestamp when the deployment package was executed.
	//
	// example:
	//
	// 1593877765000
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The Alibaba Cloud account ID of the deployment package executor.
	//
	// example:
	//
	// 2003****
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The ID of the deployment package. You can use this ID to call the [GetDeployment](https://help.aliyun.com/document_detail/173950.html) operation to get the deployment package details.
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
	// 	- 0: It is ready.
	//
	// 	- 1: It was successfully deployed.
	//
	// 	- 2: It failed to be deployed.
	//
	// 	- 6: It was rejected.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentPackagesResponseBodyDataDeployments) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackagesResponseBodyDataDeployments) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetCreator() *string {
	return s.Creator
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetExecutor() *string {
	return s.Executor
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetId() *int64 {
	return s.Id
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetName() *string {
	return s.Name
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) GetStatus() *int32 {
	return s.Status
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetCreateTime(v int64) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.CreateTime = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetCreator(v string) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.Creator = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetErrorMessage(v string) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetExecuteTime(v int64) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.ExecuteTime = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetExecutor(v string) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.Executor = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetId(v int64) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.Id = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetName(v string) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.Name = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) SetStatus(v int32) *ListDeploymentPackagesResponseBodyDataDeployments {
	s.Status = &v
	return s
}

func (s *ListDeploymentPackagesResponseBodyDataDeployments) Validate() error {
	return dara.Validate(s)
}
