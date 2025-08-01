// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetApplicationListResponseBodyData) *GetApplicationListResponseBody
	GetData() *GetApplicationListResponseBodyData
	SetMessage(v string) *GetApplicationListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplicationListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApplicationListResponseBody
	GetSuccess() *bool
}

type GetApplicationListResponseBody struct {
	// The details of the data.
	Data *GetApplicationListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F55E608F-7C15-****-9CFD-DF832EBC4A0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetApplicationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationListResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationListResponseBody) GetData() *GetApplicationListResponseBodyData {
	return s.Data
}

func (s *GetApplicationListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApplicationListResponseBody) SetData(v *GetApplicationListResponseBodyData) *GetApplicationListResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationListResponseBody) SetMessage(v string) *GetApplicationListResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationListResponseBody) SetRequestId(v string) *GetApplicationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationListResponseBody) SetSuccess(v bool) *GetApplicationListResponseBody {
	s.Success = &v
	return s
}

func (s *GetApplicationListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetApplicationListResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data returned.
	Result []*GetApplicationListResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 11
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetApplicationListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetApplicationListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetApplicationListResponseBodyData) GetResult() []*GetApplicationListResponseBodyDataResult {
	return s.Result
}

func (s *GetApplicationListResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetApplicationListResponseBodyData) SetPageNumber(v int32) *GetApplicationListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetApplicationListResponseBodyData) SetPageSize(v int32) *GetApplicationListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetApplicationListResponseBodyData) SetResult(v []*GetApplicationListResponseBodyDataResult) *GetApplicationListResponseBodyData {
	s.Result = v
	return s
}

func (s *GetApplicationListResponseBodyData) SetTotalSize(v int32) *GetApplicationListResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *GetApplicationListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetApplicationListResponseBodyDataResult struct {
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@5f1b08becb*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// finance
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The additional information.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstancesNumber *int32 `json:"InstancesNumber,omitempty" xml:"InstancesNumber,omitempty"`
	// The programming language of the application.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The microservice namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the application.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status.
	//
	// example:
	//
	// 0
	Status *int64                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1234567890
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetApplicationListResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationListResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetApplicationListResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationListResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationListResponseBodyDataResult) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetApplicationListResponseBodyDataResult) GetInstancesNumber() *int32 {
	return s.InstancesNumber
}

func (s *GetApplicationListResponseBodyDataResult) GetLanguage() *string {
	return s.Language
}

func (s *GetApplicationListResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *GetApplicationListResponseBodyDataResult) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApplicationListResponseBodyDataResult) GetSource() *string {
	return s.Source
}

func (s *GetApplicationListResponseBodyDataResult) GetStatus() *int64 {
	return s.Status
}

func (s *GetApplicationListResponseBodyDataResult) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetApplicationListResponseBodyDataResult) GetUserId() *string {
	return s.UserId
}

func (s *GetApplicationListResponseBodyDataResult) SetAppId(v string) *GetApplicationListResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetAppName(v string) *GetApplicationListResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetExtraInfo(v string) *GetApplicationListResponseBodyDataResult {
	s.ExtraInfo = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetInstancesNumber(v int32) *GetApplicationListResponseBodyDataResult {
	s.InstancesNumber = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetLanguage(v string) *GetApplicationListResponseBodyDataResult {
	s.Language = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetNamespace(v string) *GetApplicationListResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetRegionId(v string) *GetApplicationListResponseBodyDataResult {
	s.RegionId = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetSource(v string) *GetApplicationListResponseBodyDataResult {
	s.Source = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetStatus(v int64) *GetApplicationListResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetTags(v map[string]interface{}) *GetApplicationListResponseBodyDataResult {
	s.Tags = v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) SetUserId(v string) *GetApplicationListResponseBodyDataResult {
	s.UserId = &v
	return s
}

func (s *GetApplicationListResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
