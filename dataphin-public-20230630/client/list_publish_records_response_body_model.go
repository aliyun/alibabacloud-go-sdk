// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPublishRecordsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListPublishRecordsResponseBody
	GetHttpStatusCode() *int32
	SetListResult(v *ListPublishRecordsResponseBodyListResult) *ListPublishRecordsResponseBody
	GetListResult() *ListPublishRecordsResponseBodyListResult
	SetMessage(v string) *ListPublishRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPublishRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPublishRecordsResponseBody
	GetSuccess() *bool
}

type ListPublishRecordsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ListResult     *ListPublishRecordsResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPublishRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPublishRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPublishRecordsResponseBody) GetListResult() *ListPublishRecordsResponseBodyListResult {
	return s.ListResult
}

func (s *ListPublishRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPublishRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPublishRecordsResponseBody) SetCode(v string) *ListPublishRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishRecordsResponseBody) SetHttpStatusCode(v int32) *ListPublishRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPublishRecordsResponseBody) SetListResult(v *ListPublishRecordsResponseBodyListResult) *ListPublishRecordsResponseBody {
	s.ListResult = v
	return s
}

func (s *ListPublishRecordsResponseBody) SetMessage(v string) *ListPublishRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishRecordsResponseBody) SetRequestId(v string) *ListPublishRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishRecordsResponseBody) SetSuccess(v bool) *ListPublishRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPublishRecordsResponseBody) Validate() error {
	if s.ListResult != nil {
		if err := s.ListResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishRecordsResponseBodyListResult struct {
	Data []*ListPublishRecordsResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishRecordsResponseBodyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsResponseBodyListResult) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsResponseBodyListResult) GetData() []*ListPublishRecordsResponseBodyListResultData {
	return s.Data
}

func (s *ListPublishRecordsResponseBodyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublishRecordsResponseBodyListResult) SetData(v []*ListPublishRecordsResponseBodyListResultData) *ListPublishRecordsResponseBodyListResult {
	s.Data = v
	return s
}

func (s *ListPublishRecordsResponseBodyListResult) SetTotalCount(v int32) *ListPublishRecordsResponseBodyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublishRecordsResponseBodyListResultData struct {
	// example:
	//
	// 0
	ChangeType *int32 `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// example:
	//
	// error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 2024-10-10 10:10:10
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// example:
	//
	// 1241844456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// n_123456
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1234567
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// 对象A
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 1
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
	// example:
	//
	// 1241844456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test_lc__20241118171502
	PublishName *string `json:"PublishName,omitempty" xml:"PublishName,omitempty"`
	// example:
	//
	// 1
	PublishStatus *int32 `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// example:
	//
	// 307999999
	Publisher *string `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	// example:
	//
	// 张三
	PublisherName *string `json:"PublisherName,omitempty" xml:"PublisherName,omitempty"`
}

func (s ListPublishRecordsResponseBodyListResultData) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsResponseBodyListResultData) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsResponseBodyListResultData) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *ListPublishRecordsResponseBodyListResultData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPublishRecordsResponseBodyListResultData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListPublishRecordsResponseBodyListResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListPublishRecordsResponseBodyListResultData) GetGmtModify() *string {
	return s.GmtModify
}

func (s *ListPublishRecordsResponseBodyListResultData) GetId() *int64 {
	return s.Id
}

func (s *ListPublishRecordsResponseBodyListResultData) GetNodeId() *string {
	return s.NodeId
}

func (s *ListPublishRecordsResponseBodyListResultData) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListPublishRecordsResponseBodyListResultData) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListPublishRecordsResponseBodyListResultData) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListPublishRecordsResponseBodyListResultData) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *ListPublishRecordsResponseBodyListResultData) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListPublishRecordsResponseBodyListResultData) GetPublishName() *string {
	return s.PublishName
}

func (s *ListPublishRecordsResponseBodyListResultData) GetPublishStatus() *int32 {
	return s.PublishStatus
}

func (s *ListPublishRecordsResponseBodyListResultData) GetPublisher() *string {
	return s.Publisher
}

func (s *ListPublishRecordsResponseBodyListResultData) GetPublisherName() *string {
	return s.PublisherName
}

func (s *ListPublishRecordsResponseBodyListResultData) SetChangeType(v int32) *ListPublishRecordsResponseBodyListResultData {
	s.ChangeType = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetErrorMessage(v string) *ListPublishRecordsResponseBodyListResultData {
	s.ErrorMessage = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetFinishTime(v string) *ListPublishRecordsResponseBodyListResultData {
	s.FinishTime = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetGmtCreate(v string) *ListPublishRecordsResponseBodyListResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetGmtModify(v string) *ListPublishRecordsResponseBodyListResultData {
	s.GmtModify = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetId(v int64) *ListPublishRecordsResponseBodyListResultData {
	s.Id = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetNodeId(v string) *ListPublishRecordsResponseBodyListResultData {
	s.NodeId = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetObjectId(v string) *ListPublishRecordsResponseBodyListResultData {
	s.ObjectId = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetObjectName(v string) *ListPublishRecordsResponseBodyListResultData {
	s.ObjectName = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetObjectType(v string) *ListPublishRecordsResponseBodyListResultData {
	s.ObjectType = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetObjectVersion(v string) *ListPublishRecordsResponseBodyListResultData {
	s.ObjectVersion = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetProjectId(v string) *ListPublishRecordsResponseBodyListResultData {
	s.ProjectId = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetPublishName(v string) *ListPublishRecordsResponseBodyListResultData {
	s.PublishName = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetPublishStatus(v int32) *ListPublishRecordsResponseBodyListResultData {
	s.PublishStatus = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetPublisher(v string) *ListPublishRecordsResponseBodyListResultData {
	s.Publisher = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) SetPublisherName(v string) *ListPublishRecordsResponseBodyListResultData {
	s.PublisherName = &v
	return s
}

func (s *ListPublishRecordsResponseBodyListResultData) Validate() error {
	return dara.Validate(s)
}
