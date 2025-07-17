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
	Data *ListDeploymentPackagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type ListDeploymentPackagesResponseBodyData struct {
	Deployments []*ListDeploymentPackagesResponseBodyDataDeployments `json:"Deployments,omitempty" xml:"Deployments,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	return dara.Validate(s)
}

type ListDeploymentPackagesResponseBodyDataDeployments struct {
	// example:
	//
	// 1593877765000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2003****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// OK
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1593877765000
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// example:
	//
	// 2003****
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// example:
	//
	// 11111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// auto_created
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
