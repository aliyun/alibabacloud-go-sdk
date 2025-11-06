// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMdsCubeTasksResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMdsCubeTasksResponseBody
	GetResultCode() *string
	SetResultContent(v *ListMdsCubeTasksResponseBodyResultContent) *ListMdsCubeTasksResponseBody
	GetResultContent() *ListMdsCubeTasksResponseBodyResultContent
	SetResultMessage(v string) *ListMdsCubeTasksResponseBody
	GetResultMessage() *string
}

type ListMdsCubeTasksResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ListMdsCubeTasksResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMdsCubeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeTasksResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMdsCubeTasksResponseBody) GetResultContent() *ListMdsCubeTasksResponseBodyResultContent {
	return s.ResultContent
}

func (s *ListMdsCubeTasksResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMdsCubeTasksResponseBody) SetRequestId(v string) *ListMdsCubeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeTasksResponseBody) SetResultCode(v string) *ListMdsCubeTasksResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMdsCubeTasksResponseBody) SetResultContent(v *ListMdsCubeTasksResponseBodyResultContent) *ListMdsCubeTasksResponseBody {
	s.ResultContent = v
	return s
}

func (s *ListMdsCubeTasksResponseBody) SetResultMessage(v string) *ListMdsCubeTasksResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMdsCubeTasksResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeTasksResponseBodyResultContent struct {
	Data *ListMdsCubeTasksResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMdsCubeTasksResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTasksResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTasksResponseBodyResultContent) GetData() *ListMdsCubeTasksResponseBodyResultContentData {
	return s.Data
}

func (s *ListMdsCubeTasksResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeTasksResponseBodyResultContent) SetData(v *ListMdsCubeTasksResponseBodyResultContentData) *ListMdsCubeTasksResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContent) SetRequestId(v string) *ListMdsCubeTasksResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeTasksResponseBodyResultContentData struct {
	Content []*ListMdsCubeTasksResponseBodyResultContentDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// NONE
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMdsCubeTasksResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTasksResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) GetContent() []*ListMdsCubeTasksResponseBodyResultContentDataContent {
	return s.Content
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) SetContent(v []*ListMdsCubeTasksResponseBodyResultContentDataContent) *ListMdsCubeTasksResponseBodyResultContentData {
	s.Content = v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) SetErrorCode(v string) *ListMdsCubeTasksResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) SetRequestId(v string) *ListMdsCubeTasksResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) SetResultMsg(v string) *ListMdsCubeTasksResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) SetSuccess(v bool) *ListMdsCubeTasksResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMdsCubeTasksResponseBodyResultContentDataContent struct {
	// example:
	//
	// 570DA89281533-default
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// example:
	//
	// 2025-10-24 15:17:15
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-10-24 15:17:15
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// {}
	GreyConfigInfo *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	// example:
	//
	// 2025-10-24 15:17:15
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	// example:
	//
	// 1
	GreyNum *int32 `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 2
	PublishMode *int32 `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	// example:
	//
	// 3
	PublishType *int32 `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// example:
	//
	// 0.0.0.1
	ResourceVersion *string `json:"ResourceVersion,omitempty" xml:"ResourceVersion,omitempty"`
	// example:
	//
	// TEST
	TaskDesc *string `json:"TaskDesc,omitempty" xml:"TaskDesc,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 1
	TemplateResourceId *int64 `json:"TemplateResourceId,omitempty" xml:"TemplateResourceId,omitempty"`
	// example:
	//
	// TEST
	WhitelistIds *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
}

func (s ListMdsCubeTasksResponseBodyResultContentDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTasksResponseBodyResultContentDataContent) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetId() *int64 {
	return s.Id
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetOperator() *string {
	return s.Operator
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetPublishType() *int32 {
	return s.PublishType
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetResourceVersion() *string {
	return s.ResourceVersion
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetTaskDesc() *string {
	return s.TaskDesc
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetTemplateResourceId() *int64 {
	return s.TemplateResourceId
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetAppCode(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.AppCode = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetGmtCreate(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.GmtCreate = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetGmtModified(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetGreyConfigInfo(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.GreyConfigInfo = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetGreyEndtimeData(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.GreyEndtimeData = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetGreyNum(v int32) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.GreyNum = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetId(v int64) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.Id = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetOperator(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.Operator = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetPublishMode(v int32) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.PublishMode = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetPublishType(v int32) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.PublishType = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetResourceVersion(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.ResourceVersion = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetTaskDesc(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.TaskDesc = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetTaskStatus(v int32) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.TaskStatus = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetTemplateId(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.TemplateId = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetTemplateResourceId(v int64) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.TemplateResourceId = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) SetWhitelistIds(v string) *ListMdsCubeTasksResponseBodyResultContentDataContent {
	s.WhitelistIds = &v
	return s
}

func (s *ListMdsCubeTasksResponseBodyResultContentDataContent) Validate() error {
	return dara.Validate(s)
}
