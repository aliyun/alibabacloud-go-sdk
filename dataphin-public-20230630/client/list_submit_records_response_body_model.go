// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubmitRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSubmitRecordsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSubmitRecordsResponseBody
	GetHttpStatusCode() *int32
	SetListResult(v *ListSubmitRecordsResponseBodyListResult) *ListSubmitRecordsResponseBody
	GetListResult() *ListSubmitRecordsResponseBodyListResult
	SetMessage(v string) *ListSubmitRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubmitRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSubmitRecordsResponseBody
	GetSuccess() *bool
}

type ListSubmitRecordsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ListResult     *ListSubmitRecordsResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
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

func (s ListSubmitRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSubmitRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSubmitRecordsResponseBody) GetListResult() *ListSubmitRecordsResponseBodyListResult {
	return s.ListResult
}

func (s *ListSubmitRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubmitRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubmitRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubmitRecordsResponseBody) SetCode(v string) *ListSubmitRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetHttpStatusCode(v int32) *ListSubmitRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetListResult(v *ListSubmitRecordsResponseBodyListResult) *ListSubmitRecordsResponseBody {
	s.ListResult = v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetMessage(v string) *ListSubmitRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetRequestId(v string) *ListSubmitRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetSuccess(v bool) *ListSubmitRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSubmitRecordsResponseBodyListResult struct {
	Data []*ListSubmitRecordsResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubmitRecordsResponseBodyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsResponseBodyListResult) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsResponseBodyListResult) GetData() []*ListSubmitRecordsResponseBodyListResultData {
	return s.Data
}

func (s *ListSubmitRecordsResponseBodyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSubmitRecordsResponseBodyListResult) SetData(v []*ListSubmitRecordsResponseBodyListResultData) *ListSubmitRecordsResponseBodyListResult {
	s.Data = v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResult) SetTotalCount(v int32) *ListSubmitRecordsResponseBodyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResult) Validate() error {
	return dara.Validate(s)
}

type ListSubmitRecordsResponseBodyListResultData struct {
	// example:
	//
	// 0
	ChangeType *int32 `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
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
	// 提交信息
	SubmitComment *string `json:"SubmitComment,omitempty" xml:"SubmitComment,omitempty"`
	// example:
	//
	// 307999999
	Submitter *string `json:"Submitter,omitempty" xml:"Submitter,omitempty"`
	// example:
	//
	// 张三
	SubmitterName *string `json:"SubmitterName,omitempty" xml:"SubmitterName,omitempty"`
}

func (s ListSubmitRecordsResponseBodyListResultData) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsResponseBodyListResultData) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetGmtModify() *string {
	return s.GmtModify
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetId() *int64 {
	return s.Id
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetNodeId() *string {
	return s.NodeId
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetSubmitComment() *string {
	return s.SubmitComment
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetSubmitter() *string {
	return s.Submitter
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetSubmitterName() *string {
	return s.SubmitterName
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetChangeType(v int32) *ListSubmitRecordsResponseBodyListResultData {
	s.ChangeType = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetGmtCreate(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetGmtModify(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.GmtModify = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetId(v int64) *ListSubmitRecordsResponseBodyListResultData {
	s.Id = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetNodeId(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.NodeId = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectId(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectId = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectName(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectName = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectType(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectType = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectVersion(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectVersion = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetProjectId(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ProjectId = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetSubmitComment(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.SubmitComment = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetSubmitter(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.Submitter = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetSubmitterName(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.SubmitterName = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) Validate() error {
	return dara.Validate(s)
}
