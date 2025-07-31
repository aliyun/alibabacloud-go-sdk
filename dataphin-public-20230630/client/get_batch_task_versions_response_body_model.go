// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchTaskVersionsResponseBody
	GetCode() *string
	SetData(v *GetBatchTaskVersionsResponseBodyData) *GetBatchTaskVersionsResponseBody
	GetData() *GetBatchTaskVersionsResponseBodyData
	SetHttpStatusCode(v int32) *GetBatchTaskVersionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBatchTaskVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTaskVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBatchTaskVersionsResponseBody
	GetSuccess() *bool
}

type GetBatchTaskVersionsResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBatchTaskVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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

func (s GetBatchTaskVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTaskVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchTaskVersionsResponseBody) GetData() *GetBatchTaskVersionsResponseBodyData {
	return s.Data
}

func (s *GetBatchTaskVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBatchTaskVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTaskVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTaskVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBatchTaskVersionsResponseBody) SetCode(v string) *GetBatchTaskVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBody) SetData(v *GetBatchTaskVersionsResponseBodyData) *GetBatchTaskVersionsResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchTaskVersionsResponseBody) SetHttpStatusCode(v int32) *GetBatchTaskVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBody) SetMessage(v string) *GetBatchTaskVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBody) SetRequestId(v string) *GetBatchTaskVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBody) SetSuccess(v bool) *GetBatchTaskVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskVersionsResponseBodyData struct {
	BatchTaskVersionList []*GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList `json:"BatchTaskVersionList,omitempty" xml:"BatchTaskVersionList,omitempty" type:"Repeated"`
}

func (s GetBatchTaskVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBatchTaskVersionsResponseBodyData) GetBatchTaskVersionList() []*GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	return s.BatchTaskVersionList
}

func (s *GetBatchTaskVersionsResponseBodyData) SetBatchTaskVersionList(v []*GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) *GetBatchTaskVersionsResponseBodyData {
	s.BatchTaskVersionList = v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList struct {
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2024-10-10 10:10:10
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-10-10 10:10:10
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// n_10231001
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 10232111011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Published *bool  `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// 20110110
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetComment() *string {
	return s.Comment
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetPublished() *bool {
	return s.Published
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetUserId() *string {
	return s.UserId
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetUserName() *string {
	return s.UserName
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) GetVersion() *string {
	return s.Version
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetComment(v string) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.Comment = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetGmtCreate(v string) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.GmtCreate = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetGmtModified(v string) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.GmtModified = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetNodeId(v string) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.NodeId = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetProjectId(v int64) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetPublished(v bool) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.Published = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetUserId(v string) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.UserId = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetUserName(v string) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.UserName = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) SetVersion(v string) *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList {
	s.Version = &v
	return s
}

func (s *GetBatchTaskVersionsResponseBodyDataBatchTaskVersionList) Validate() error {
	return dara.Validate(s)
}
