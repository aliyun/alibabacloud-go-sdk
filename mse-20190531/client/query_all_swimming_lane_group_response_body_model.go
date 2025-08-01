// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryAllSwimmingLaneGroupResponseBody
	GetCode() *int32
	SetData(v []*QueryAllSwimmingLaneGroupResponseBodyData) *QueryAllSwimmingLaneGroupResponseBody
	GetData() []*QueryAllSwimmingLaneGroupResponseBodyData
	SetDynamicMessage(v string) *QueryAllSwimmingLaneGroupResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *QueryAllSwimmingLaneGroupResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *QueryAllSwimmingLaneGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryAllSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAllSwimmingLaneGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAllSwimmingLaneGroupResponseBody
	GetSuccess() *bool
}

type QueryAllSwimmingLaneGroupResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	//
	// example:
	//
	// [{id:100,name:"test"}]
	Data []*QueryAllSwimmingLaneGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54973C90-F379-4372-9AA5-053A3F7****
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

func (s QueryAllSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetData() []*QueryAllSwimmingLaneGroupResponseBodyData {
	return s.Data
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAllSwimmingLaneGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetCode(v int32) *QueryAllSwimmingLaneGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetData(v []*QueryAllSwimmingLaneGroupResponseBodyData) *QueryAllSwimmingLaneGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetDynamicMessage(v string) *QueryAllSwimmingLaneGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetErrorCode(v string) *QueryAllSwimmingLaneGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetHttpStatusCode(v int32) *QueryAllSwimmingLaneGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetMessage(v string) *QueryAllSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetRequestId(v string) *QueryAllSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) SetSuccess(v bool) *QueryAllSwimmingLaneGroupResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAllSwimmingLaneGroupResponseBodyData struct {
	// example:
	//
	// abcde@abcde,abcde@abcde
	AppIds      *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	CanaryModel *int32  `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// example:
	//
	// mse:abcde@abcde
	EntryApp *string `json:"EntryApp,omitempty" xml:"EntryApp,omitempty"`
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Client
	MessageQueueFilterSide *string `json:"MessageQueueFilterSide,omitempty" xml:"MessageQueueFilterSide,omitempty"`
	MessageQueueGrayEnable *bool   `json:"MessageQueueGrayEnable,omitempty" xml:"MessageQueueGrayEnable,omitempty"`
	// example:
	//
	// swimmingGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// default
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Paths              *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
	RecordCanaryDetail *bool   `json:"RecordCanaryDetail,omitempty" xml:"RecordCanaryDetail,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SwimVersion *int32  `json:"SwimVersion,omitempty" xml:"SwimVersion,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryAllSwimmingLaneGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetAppIds() *string {
	return s.AppIds
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetEntryApp() *string {
	return s.EntryApp
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetMessageQueueFilterSide() *string {
	return s.MessageQueueFilterSide
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetMessageQueueGrayEnable() *bool {
	return s.MessageQueueGrayEnable
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetPaths() *string {
	return s.Paths
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetRecordCanaryDetail() *bool {
	return s.RecordCanaryDetail
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetSwimVersion() *int32 {
	return s.SwimVersion
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetAppIds(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.AppIds = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetCanaryModel(v int32) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.CanaryModel = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetEntryApp(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.EntryApp = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetId(v int64) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetMessageQueueFilterSide(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.MessageQueueFilterSide = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetMessageQueueGrayEnable(v bool) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.MessageQueueGrayEnable = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetName(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetNamespace(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetPaths(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.Paths = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetRecordCanaryDetail(v bool) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.RecordCanaryDetail = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetRegion(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.Region = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetSwimVersion(v int32) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.SwimVersion = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) SetUserId(v string) *QueryAllSwimmingLaneGroupResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
