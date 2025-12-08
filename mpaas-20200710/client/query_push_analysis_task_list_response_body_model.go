// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPushAnalysisTaskListResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryPushAnalysisTaskListResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryPushAnalysisTaskListResponseBodyResultContent) *QueryPushAnalysisTaskListResponseBody
	GetResultContent() *QueryPushAnalysisTaskListResponseBodyResultContent
	SetResultMessage(v string) *QueryPushAnalysisTaskListResponseBody
	GetResultMessage() *string
}

type QueryPushAnalysisTaskListResponseBody struct {
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryPushAnalysisTaskListResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryPushAnalysisTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPushAnalysisTaskListResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryPushAnalysisTaskListResponseBody) GetResultContent() *QueryPushAnalysisTaskListResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryPushAnalysisTaskListResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryPushAnalysisTaskListResponseBody) SetRequestId(v string) *QueryPushAnalysisTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBody) SetResultCode(v string) *QueryPushAnalysisTaskListResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBody) SetResultContent(v *QueryPushAnalysisTaskListResponseBodyResultContent) *QueryPushAnalysisTaskListResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBody) SetResultMessage(v string) *QueryPushAnalysisTaskListResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPushAnalysisTaskListResponseBodyResultContent struct {
	Data []*QueryPushAnalysisTaskListResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryPushAnalysisTaskListResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskListResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContent) GetData() []*QueryPushAnalysisTaskListResponseBodyResultContentData {
	return s.Data
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContent) SetData(v []*QueryPushAnalysisTaskListResponseBodyResultContentData) *QueryPushAnalysisTaskListResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContent) Validate() error {
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

type QueryPushAnalysisTaskListResponseBodyResultContentData struct {
	GmtCreate    *int64                                                        `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	List         []*QueryPushAnalysisTaskListResponseBodyResultContentDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TaskId       *string                                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName     *string                                                       `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateId   *string                                                       `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string                                                       `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Type         *int64                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryPushAnalysisTaskListResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskListResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) GetList() []*QueryPushAnalysisTaskListResponseBodyResultContentDataList {
	return s.List
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) GetType() *int64 {
	return s.Type
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) SetGmtCreate(v int64) *QueryPushAnalysisTaskListResponseBodyResultContentData {
	s.GmtCreate = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) SetList(v []*QueryPushAnalysisTaskListResponseBodyResultContentDataList) *QueryPushAnalysisTaskListResponseBodyResultContentData {
	s.List = v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) SetTaskId(v string) *QueryPushAnalysisTaskListResponseBodyResultContentData {
	s.TaskId = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) SetTaskName(v string) *QueryPushAnalysisTaskListResponseBodyResultContentData {
	s.TaskName = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) SetTemplateId(v string) *QueryPushAnalysisTaskListResponseBodyResultContentData {
	s.TemplateId = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) SetTemplateName(v string) *QueryPushAnalysisTaskListResponseBodyResultContentData {
	s.TemplateName = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) SetType(v int64) *QueryPushAnalysisTaskListResponseBodyResultContentData {
	s.Type = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPushAnalysisTaskListResponseBodyResultContentDataList struct {
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName     *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Type         *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryPushAnalysisTaskListResponseBodyResultContentDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskListResponseBodyResultContentDataList) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) GetType() *int64 {
	return s.Type
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) SetGmtCreate(v int64) *QueryPushAnalysisTaskListResponseBodyResultContentDataList {
	s.GmtCreate = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) SetTaskId(v string) *QueryPushAnalysisTaskListResponseBodyResultContentDataList {
	s.TaskId = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) SetTaskName(v string) *QueryPushAnalysisTaskListResponseBodyResultContentDataList {
	s.TaskName = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) SetTemplateId(v string) *QueryPushAnalysisTaskListResponseBodyResultContentDataList {
	s.TemplateId = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) SetTemplateName(v string) *QueryPushAnalysisTaskListResponseBodyResultContentDataList {
	s.TemplateName = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) SetType(v int64) *QueryPushAnalysisTaskListResponseBodyResultContentDataList {
	s.Type = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponseBodyResultContentDataList) Validate() error {
	return dara.Validate(s)
}
