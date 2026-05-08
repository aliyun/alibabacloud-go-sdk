// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryIndividuationTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchQueryIndividuationTextResponseBody
	GetRequestId() *string
	SetTextList(v []*BatchQueryIndividuationTextResponseBodyTextList) *BatchQueryIndividuationTextResponseBody
	GetTextList() []*BatchQueryIndividuationTextResponseBodyTextList
}

type BatchQueryIndividuationTextResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 14878724-A835-578D-9DD5-4779ADCE9221
	RequestId *string                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TextList  []*BatchQueryIndividuationTextResponseBodyTextList `json:"textList,omitempty" xml:"textList,omitempty" type:"Repeated"`
}

func (s BatchQueryIndividuationTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryIndividuationTextResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryIndividuationTextResponseBody) GetTextList() []*BatchQueryIndividuationTextResponseBodyTextList {
	return s.TextList
}

func (s *BatchQueryIndividuationTextResponseBody) SetRequestId(v string) *BatchQueryIndividuationTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBody) SetTextList(v []*BatchQueryIndividuationTextResponseBodyTextList) *BatchQueryIndividuationTextResponseBody {
	s.TextList = v
	return s
}

func (s *BatchQueryIndividuationTextResponseBody) Validate() error {
	if s.TextList != nil {
		for _, item := range s.TextList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchQueryIndividuationTextResponseBodyTextList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 2849286
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 812884915104530432
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 837074737851613184
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 110825
	TextId *string `json:"textId,omitempty" xml:"textId,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 11
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryIndividuationTextResponseBodyTextList) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryIndividuationTextResponseBodyTextList) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetContent() *string {
	return s.Content
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetItemId() *string {
	return s.ItemId
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetProjectId() *string {
	return s.ProjectId
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetStatus() *string {
	return s.Status
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetTextId() *string {
	return s.TextId
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) GetUserId() *string {
	return s.UserId
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetContent(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.Content = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetCreateTime(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.CreateTime = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetErrorMsg(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.ErrorMsg = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetItemId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.ItemId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetProjectId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.ProjectId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetStatus(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.Status = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetTaskId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.TaskId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetTextId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.TextId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetUpdateTime(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.UpdateTime = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetUserId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.UserId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) Validate() error {
	return dara.Validate(s)
}
