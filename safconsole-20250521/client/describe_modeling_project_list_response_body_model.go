// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelingProjectListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeModelingProjectListResponseBody
	GetCode() *int64
	SetCurrentPage(v int64) *DescribeModelingProjectListResponseBody
	GetCurrentPage() *int64
	SetPageSize(v int64) *DescribeModelingProjectListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeModelingProjectListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeModelingProjectListResponseBodyResultObject) *DescribeModelingProjectListResponseBody
	GetResultObject() []*DescribeModelingProjectListResponseBodyResultObject
	SetSuccess(v bool) *DescribeModelingProjectListResponseBody
	GetSuccess() *bool
	SetTotalItem(v int64) *DescribeModelingProjectListResponseBody
	GetTotalItem() *int64
	SetTotalPage(v int64) *DescribeModelingProjectListResponseBody
	GetTotalPage() *int64
}

type DescribeModelingProjectListResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Pagination parameter: number of items per page, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result.
	ResultObject []*DescribeModelingProjectListResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 1
	TotalItem *int64 `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeModelingProjectListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeModelingProjectListResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeModelingProjectListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeModelingProjectListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelingProjectListResponseBody) GetResultObject() []*DescribeModelingProjectListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeModelingProjectListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModelingProjectListResponseBody) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeModelingProjectListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeModelingProjectListResponseBody) SetCode(v int64) *DescribeModelingProjectListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeModelingProjectListResponseBody) SetCurrentPage(v int64) *DescribeModelingProjectListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeModelingProjectListResponseBody) SetPageSize(v int64) *DescribeModelingProjectListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeModelingProjectListResponseBody) SetRequestId(v string) *DescribeModelingProjectListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelingProjectListResponseBody) SetResultObject(v []*DescribeModelingProjectListResponseBodyResultObject) *DescribeModelingProjectListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeModelingProjectListResponseBody) SetSuccess(v bool) *DescribeModelingProjectListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModelingProjectListResponseBody) SetTotalItem(v int64) *DescribeModelingProjectListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeModelingProjectListResponseBody) SetTotalPage(v int64) *DescribeModelingProjectListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeModelingProjectListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelingProjectListResponseBodyResultObject struct {
	// End time of the secure modeling project.
	//
	// example:
	//
	// 2025-12-29T02:15:10Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Secure environment status.
	//
	// example:
	//
	// PREPARING
	EnvStatus *string `json:"EnvStatus,omitempty" xml:"EnvStatus,omitempty"`
	// Login account.
	//
	// example:
	//
	// xxx
	LoginAccount *string `json:"LoginAccount,omitempty" xml:"LoginAccount,omitempty"`
	// Modeling project status.
	//
	// example:
	//
	// PREPARING
	ModelingStatus *string `json:"ModelingStatus,omitempty" xml:"ModelingStatus,omitempty"`
	// Project ID.
	//
	// example:
	//
	// 01
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Project name.
	//
	// example:
	//
	// project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Start time of the secure modeling project.
	//
	// example:
	//
	// 2025-03-23T01:20:00Z
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModelingProjectListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectListResponseBodyResultObject) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeModelingProjectListResponseBodyResultObject) GetEnvStatus() *string {
	return s.EnvStatus
}

func (s *DescribeModelingProjectListResponseBodyResultObject) GetLoginAccount() *string {
	return s.LoginAccount
}

func (s *DescribeModelingProjectListResponseBodyResultObject) GetModelingStatus() *string {
	return s.ModelingStatus
}

func (s *DescribeModelingProjectListResponseBodyResultObject) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeModelingProjectListResponseBodyResultObject) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeModelingProjectListResponseBodyResultObject) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeModelingProjectListResponseBodyResultObject) SetEndTime(v int64) *DescribeModelingProjectListResponseBodyResultObject {
	s.EndTime = &v
	return s
}

func (s *DescribeModelingProjectListResponseBodyResultObject) SetEnvStatus(v string) *DescribeModelingProjectListResponseBodyResultObject {
	s.EnvStatus = &v
	return s
}

func (s *DescribeModelingProjectListResponseBodyResultObject) SetLoginAccount(v string) *DescribeModelingProjectListResponseBodyResultObject {
	s.LoginAccount = &v
	return s
}

func (s *DescribeModelingProjectListResponseBodyResultObject) SetModelingStatus(v string) *DescribeModelingProjectListResponseBodyResultObject {
	s.ModelingStatus = &v
	return s
}

func (s *DescribeModelingProjectListResponseBodyResultObject) SetProjectId(v string) *DescribeModelingProjectListResponseBodyResultObject {
	s.ProjectId = &v
	return s
}

func (s *DescribeModelingProjectListResponseBodyResultObject) SetProjectName(v string) *DescribeModelingProjectListResponseBodyResultObject {
	s.ProjectName = &v
	return s
}

func (s *DescribeModelingProjectListResponseBodyResultObject) SetStartTime(v int64) *DescribeModelingProjectListResponseBodyResultObject {
	s.StartTime = &v
	return s
}

func (s *DescribeModelingProjectListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
