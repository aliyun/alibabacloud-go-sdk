// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeListModelResponseBody
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeListModelResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeListModelResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeListModelResponseBodyResultObject) *DescribeListModelResponseBody
	GetResultObject() []*DescribeListModelResponseBodyResultObject
	SetTotalItem(v string) *DescribeListModelResponseBody
	GetTotalItem() *string
	SetTotalPage(v string) *DescribeListModelResponseBody
	GetTotalPage() *string
}

type DescribeListModelResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data.
	ResultObject []*DescribeListModelResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// Total number of records.
	//
	// example:
	//
	// 1
	TotalItem *string `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *string `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeListModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeListModelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListModelResponseBody) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeListModelResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeListModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeListModelResponseBody) GetResultObject() []*DescribeListModelResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeListModelResponseBody) GetTotalItem() *string {
	return s.TotalItem
}

func (s *DescribeListModelResponseBody) GetTotalPage() *string {
	return s.TotalPage
}

func (s *DescribeListModelResponseBody) SetCurrentPage(v string) *DescribeListModelResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeListModelResponseBody) SetPageSize(v string) *DescribeListModelResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeListModelResponseBody) SetRequestId(v string) *DescribeListModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListModelResponseBody) SetResultObject(v []*DescribeListModelResponseBodyResultObject) *DescribeListModelResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeListModelResponseBody) SetTotalItem(v string) *DescribeListModelResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeListModelResponseBody) SetTotalPage(v string) *DescribeListModelResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeListModelResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeListModelResponseBodyResultObject struct {
	// Uploader ID.
	//
	// example:
	//
	// 011786
	BucId *string `json:"bucId,omitempty" xml:"bucId,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2025-05-07T02:26:01Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Model code.
	//
	// example:
	//
	// code
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// Unique identifier of the model in use.
	//
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// Model name.
	//
	// example:
	//
	// custom_model
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// Model scenario.
	//
	// example:
	//
	// custom_scene
	ModelScene *string `json:"modelScene,omitempty" xml:"modelScene,omitempty"`
	// Model status, values: -**ENABLED**: Enabled-**DISABLED**: Disabled
	//
	// example:
	//
	// ENABLED
	ModelStatus *string `json:"modelStatus,omitempty" xml:"modelStatus,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 0c8a3799-5ac8-434b-b255-e06edb35c05f
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// Last update time of the model.
	//
	// example:
	//
	// 2024-11-18T02:07:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// User ID.
	//
	// example:
	//
	// 1806507582222226
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// File name.
	//
	// example:
	//
	// saf-ai-1662519103706.csv
	UserLocalFileName *string `json:"userLocalFileName,omitempty" xml:"userLocalFileName,omitempty"`
}

func (s DescribeListModelResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeListModelResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeListModelResponseBodyResultObject) GetBucId() *string {
	return s.BucId
}

func (s *DescribeListModelResponseBodyResultObject) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeListModelResponseBodyResultObject) GetModelCode() *string {
	return s.ModelCode
}

func (s *DescribeListModelResponseBodyResultObject) GetModelId() *string {
	return s.ModelId
}

func (s *DescribeListModelResponseBodyResultObject) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeListModelResponseBodyResultObject) GetModelScene() *string {
	return s.ModelScene
}

func (s *DescribeListModelResponseBodyResultObject) GetModelStatus() *string {
	return s.ModelStatus
}

func (s *DescribeListModelResponseBodyResultObject) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeListModelResponseBodyResultObject) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeListModelResponseBodyResultObject) GetUserId() *string {
	return s.UserId
}

func (s *DescribeListModelResponseBodyResultObject) GetUserLocalFileName() *string {
	return s.UserLocalFileName
}

func (s *DescribeListModelResponseBodyResultObject) SetBucId(v string) *DescribeListModelResponseBodyResultObject {
	s.BucId = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetCreateTime(v string) *DescribeListModelResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetModelCode(v string) *DescribeListModelResponseBodyResultObject {
	s.ModelCode = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetModelId(v string) *DescribeListModelResponseBodyResultObject {
	s.ModelId = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetModelName(v string) *DescribeListModelResponseBodyResultObject {
	s.ModelName = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetModelScene(v string) *DescribeListModelResponseBodyResultObject {
	s.ModelScene = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetModelStatus(v string) *DescribeListModelResponseBodyResultObject {
	s.ModelStatus = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetTaskId(v string) *DescribeListModelResponseBodyResultObject {
	s.TaskId = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetUpdateTime(v string) *DescribeListModelResponseBodyResultObject {
	s.UpdateTime = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetUserId(v string) *DescribeListModelResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) SetUserLocalFileName(v string) *DescribeListModelResponseBodyResultObject {
	s.UserLocalFileName = &v
	return s
}

func (s *DescribeListModelResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
