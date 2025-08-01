// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetServiceListPageResponseBodyData) *GetServiceListPageResponseBody
	GetData() *GetServiceListPageResponseBodyData
	SetMessage(v string) *GetServiceListPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceListPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceListPageResponseBody
	GetSuccess() *bool
}

type GetServiceListPageResponseBody struct {
	// The response to the request.
	Data *GetServiceListPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request information.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A2F946FB-F2E3-5BF4-8CBE-xxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceListPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceListPageResponseBody) GetData() *GetServiceListPageResponseBodyData {
	return s.Data
}

func (s *GetServiceListPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceListPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceListPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceListPageResponseBody) SetData(v *GetServiceListPageResponseBodyData) *GetServiceListPageResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceListPageResponseBody) SetMessage(v string) *GetServiceListPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceListPageResponseBody) SetRequestId(v string) *GetServiceListPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceListPageResponseBody) SetSuccess(v bool) *GetServiceListPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceListPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceListPageResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data on the current page.
	Result []*GetServiceListPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalSize *string `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetServiceListPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceListPageResponseBodyData) GetPageNumber() *string {
	return s.PageNumber
}

func (s *GetServiceListPageResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *GetServiceListPageResponseBodyData) GetResult() []*GetServiceListPageResponseBodyDataResult {
	return s.Result
}

func (s *GetServiceListPageResponseBodyData) GetTotalSize() *string {
	return s.TotalSize
}

func (s *GetServiceListPageResponseBodyData) SetPageNumber(v string) *GetServiceListPageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetServiceListPageResponseBodyData) SetPageSize(v string) *GetServiceListPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetServiceListPageResponseBodyData) SetResult(v []*GetServiceListPageResponseBodyDataResult) *GetServiceListPageResponseBodyData {
	s.Result = v
	return s
}

func (s *GetServiceListPageResponseBodyData) SetTotalSize(v string) *GetServiceListPageResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *GetServiceListPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetServiceListPageResponseBodyDataResult struct {
	// The application ID.
	//
	// example:
	//
	// dez4xxxxx@f3f75ed8ffxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the service was last updated.
	//
	// example:
	//
	// 123456
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// The group to which the service belongs.
	//
	// example:
	//
	// DEFAULT_GROUP
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The number of service nodes.
	//
	// example:
	//
	// 1
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The service name.
	//
	// example:
	//
	// sc-A
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetServiceListPageResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetServiceListPageResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *GetServiceListPageResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *GetServiceListPageResponseBodyDataResult) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *GetServiceListPageResponseBodyDataResult) GetGroup() *string {
	return s.Group
}

func (s *GetServiceListPageResponseBodyDataResult) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *GetServiceListPageResponseBodyDataResult) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceListPageResponseBodyDataResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceListPageResponseBodyDataResult) GetVersion() *string {
	return s.Version
}

func (s *GetServiceListPageResponseBodyDataResult) SetAppId(v string) *GetServiceListPageResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) SetAppName(v string) *GetServiceListPageResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) SetGmtModifyTime(v string) *GetServiceListPageResponseBodyDataResult {
	s.GmtModifyTime = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) SetGroup(v string) *GetServiceListPageResponseBodyDataResult {
	s.Group = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) SetInstanceNum(v int32) *GetServiceListPageResponseBodyDataResult {
	s.InstanceNum = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) SetServiceName(v string) *GetServiceListPageResponseBodyDataResult {
	s.ServiceName = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) SetServiceType(v string) *GetServiceListPageResponseBodyDataResult {
	s.ServiceType = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) SetVersion(v string) *GetServiceListPageResponseBodyDataResult {
	s.Version = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
