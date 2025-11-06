// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateOrUpdateSwimmingLaneGroupResponseBodyData) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetData() *CreateOrUpdateSwimmingLaneGroupResponseBodyData
	SetErrorCode(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetSuccess() *bool
}

type CreateOrUpdateSwimmingLaneGroupResponseBody struct {
	// The response parameters.
	//
	// example:
	//
	// {}
	Data *CreateOrUpdateSwimmingLaneGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetData() *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	return s.Data
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetData(v *CreateOrUpdateSwimmingLaneGroupResponseBodyData) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetErrorCode(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetMessage(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetRequestId(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetSuccess(v bool) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneGroupResponseBodyData struct {
	// example:
	//
	// abcd1@abcde123,abcd1@abcde124
	AppIds      *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	CanaryModel *int32  `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// example:
	//
	// true
	DbGrayEnable *string `json:"DbGrayEnable,omitempty" xml:"DbGrayEnable,omitempty"`
	// example:
	//
	// mse:abcd1@a2345
	EntryApp *string `json:"EntryApp,omitempty" xml:"EntryApp,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Client
	MessageQueueFilterSide *string `json:"MessageQueueFilterSide,omitempty" xml:"MessageQueueFilterSide,omitempty"`
	MessageQueueGrayEnable *bool   `json:"MessageQueueGrayEnable,omitempty" xml:"MessageQueueGrayEnable,omitempty"`
	// example:
	//
	// example-app
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// prod
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Paths              *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
	RecordCanaryDetail *bool   `json:"RecordCanaryDetail,omitempty" xml:"RecordCanaryDetail,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetAppIds() *string {
	return s.AppIds
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetDbGrayEnable() *string {
	return s.DbGrayEnable
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetEntryApp() *string {
	return s.EntryApp
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetMessageQueueFilterSide() *string {
	return s.MessageQueueFilterSide
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetMessageQueueGrayEnable() *bool {
	return s.MessageQueueGrayEnable
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetPaths() *string {
	return s.Paths
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetRecordCanaryDetail() *bool {
	return s.RecordCanaryDetail
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetAppIds(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.AppIds = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.CanaryModel = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetDbGrayEnable(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.DbGrayEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetEntryApp(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.EntryApp = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetId(v int64) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetMessageQueueFilterSide(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.MessageQueueFilterSide = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetMessageQueueGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.MessageQueueGrayEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetName(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetNamespace(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetPaths(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.Paths = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetRecordCanaryDetail(v bool) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.RecordCanaryDetail = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetRegion(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.Region = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetUserId(v string) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
