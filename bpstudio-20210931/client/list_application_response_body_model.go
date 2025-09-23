// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApplicationResponseBody
	GetCode() *string
	SetData(v []*ListApplicationResponseBodyData) *ListApplicationResponseBody
	GetData() []*ListApplicationResponseBodyData
	SetMessage(v string) *ListApplicationResponseBody
	GetMessage() *string
	SetNextToken(v int32) *ListApplicationResponseBody
	GetNextToken() *int32
	SetRequestId(v string) *ListApplicationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationResponseBody
	GetTotalCount() *int32
}

type ListApplicationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// App listing information
	Data []*ListApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The interface returns information
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The query token returned in this call.
	//
	// example:
	//
	// 2
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 123
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApplicationResponseBody) GetData() []*ListApplicationResponseBodyData {
	return s.Data
}

func (s *ListApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationResponseBody) GetNextToken() *int32 {
	return s.NextToken
}

func (s *ListApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationResponseBody) SetCode(v string) *ListApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationResponseBody) SetData(v []*ListApplicationResponseBodyData) *ListApplicationResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationResponseBody) SetMessage(v string) *ListApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationResponseBody) SetNextToken(v int32) *ListApplicationResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationResponseBody) SetRequestId(v string) *ListApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationResponseBody) SetTotalCount(v int32) *ListApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// JIX9NEZUALGS46UI
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-15  08:30:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The URL of the application architecture image.
	//
	// example:
	//
	// https://bp-studio-daily.oss-cn-beijing.aliyuncs.com/1411182597819805/sr-Y3KR7ZSQZR2F0YX3.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// cadt-appName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the application belongs.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the application.
	//
	// example:
	//
	// Deployed_Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApplicationResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *ListApplicationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListApplicationResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApplicationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationResponseBodyData) SetApplicationId(v string) *ListApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetCreateTime(v string) *ListApplicationResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetImageURL(v string) *ListApplicationResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetName(v string) *ListApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetResourceGroupId(v string) *ListApplicationResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetStatus(v string) *ListApplicationResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
