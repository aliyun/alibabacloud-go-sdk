// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMemoryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMemoryListResponseBody
	GetCode() *string
	SetData(v *QueryMemoryListResponseBodyData) *QueryMemoryListResponseBody
	GetData() *QueryMemoryListResponseBodyData
	SetHttpStatusCode(v int32) *QueryMemoryListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryMemoryListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMemoryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMemoryListResponseBody
	GetSuccess() *bool
}

type QueryMemoryListResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryMemoryListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A42FFCBD-33A9-54AA-9351-86E3C3B316A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMemoryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMemoryListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMemoryListResponseBody) GetData() *QueryMemoryListResponseBodyData {
	return s.Data
}

func (s *QueryMemoryListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryMemoryListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMemoryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMemoryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMemoryListResponseBody) SetCode(v string) *QueryMemoryListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMemoryListResponseBody) SetData(v *QueryMemoryListResponseBodyData) *QueryMemoryListResponseBody {
	s.Data = v
	return s
}

func (s *QueryMemoryListResponseBody) SetHttpStatusCode(v int32) *QueryMemoryListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryMemoryListResponseBody) SetMessage(v string) *QueryMemoryListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMemoryListResponseBody) SetRequestId(v string) *QueryMemoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMemoryListResponseBody) SetSuccess(v bool) *QueryMemoryListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMemoryListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMemoryListResponseBodyData struct {
	MemoryNodes []*QueryMemoryListResponseBodyDataMemoryNodes `json:"MemoryNodes,omitempty" xml:"MemoryNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 60
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMemoryListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMemoryListResponseBodyData) GetMemoryNodes() []*QueryMemoryListResponseBodyDataMemoryNodes {
	return s.MemoryNodes
}

func (s *QueryMemoryListResponseBodyData) GetPageNum() *string {
	return s.PageNum
}

func (s *QueryMemoryListResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryMemoryListResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *QueryMemoryListResponseBodyData) SetMemoryNodes(v []*QueryMemoryListResponseBodyDataMemoryNodes) *QueryMemoryListResponseBodyData {
	s.MemoryNodes = v
	return s
}

func (s *QueryMemoryListResponseBodyData) SetPageNum(v string) *QueryMemoryListResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryMemoryListResponseBodyData) SetPageSize(v string) *QueryMemoryListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMemoryListResponseBodyData) SetTotal(v string) *QueryMemoryListResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryMemoryListResponseBodyData) Validate() error {
	if s.MemoryNodes != nil {
		for _, item := range s.MemoryNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMemoryListResponseBodyDataMemoryNodes struct {
	// example:
	//
	// []
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1743991502383
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ADD
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 384dc4786b9d4f5a8cab0d83112cd5a8
	MemoryNodeId *string                `json:"MemoryNodeId,omitempty" xml:"MemoryNodeId,omitempty"`
	MetaData     map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// []
	MetaDataJson *string `json:"MetaDataJson,omitempty" xml:"MetaDataJson,omitempty"`
	// example:
	//
	// []
	OldContent *string `json:"OldContent,omitempty" xml:"OldContent,omitempty"`
	// example:
	//
	// profile_project
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1743991502383
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 1743991502383
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryMemoryListResponseBodyDataMemoryNodes) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryListResponseBodyDataMemoryNodes) GoString() string {
	return s.String()
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetContent() *string {
	return s.Content
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetEvent() *string {
	return s.Event
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetMetaDataJson() *string {
	return s.MetaDataJson
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetOldContent() *string {
	return s.OldContent
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetContent(v string) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.Content = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetCreateTime(v int64) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.CreateTime = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetEvent(v string) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.Event = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetMemoryNodeId(v string) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.MemoryNodeId = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetMetaData(v map[string]interface{}) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.MetaData = v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetMetaDataJson(v string) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.MetaDataJson = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetOldContent(v string) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.OldContent = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetProjectId(v string) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.ProjectId = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetTimestamp(v int64) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.Timestamp = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) SetUpdateTime(v int64) *QueryMemoryListResponseBodyDataMemoryNodes {
	s.UpdateTime = &v
	return s
}

func (s *QueryMemoryListResponseBodyDataMemoryNodes) Validate() error {
	return dara.Validate(s)
}
