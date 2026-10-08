// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBatchTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListBatchTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBatchTasksResponseBody
	GetMessage() *string
	SetPageResult(v *ListBatchTasksResponseBodyPageResult) *ListBatchTasksResponseBody
	GetPageResult() *ListBatchTasksResponseBodyPageResult
	SetRequestId(v string) *ListBatchTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBatchTasksResponseBody
	GetSuccess() *bool
}

type ListBatchTasksResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListBatchTasksResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBatchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBatchTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBatchTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBatchTasksResponseBody) GetPageResult() *ListBatchTasksResponseBodyPageResult {
	return s.PageResult
}

func (s *ListBatchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBatchTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBatchTasksResponseBody) SetCode(v string) *ListBatchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListBatchTasksResponseBody) SetHttpStatusCode(v int32) *ListBatchTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBatchTasksResponseBody) SetMessage(v string) *ListBatchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListBatchTasksResponseBody) SetPageResult(v *ListBatchTasksResponseBodyPageResult) *ListBatchTasksResponseBody {
	s.PageResult = v
	return s
}

func (s *ListBatchTasksResponseBody) SetRequestId(v string) *ListBatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchTasksResponseBody) SetSuccess(v bool) *ListBatchTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListBatchTasksResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBatchTasksResponseBodyPageResult struct {
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResultData []*ListBatchTasksResponseBodyPageResultResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Repeated"`
}

func (s ListBatchTasksResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTasksResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponseBodyPageResult) GetCount() *int32 {
	return s.Count
}

func (s *ListBatchTasksResponseBodyPageResult) GetPage() *int32 {
	return s.Page
}

func (s *ListBatchTasksResponseBodyPageResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBatchTasksResponseBodyPageResult) GetResultData() []*ListBatchTasksResponseBodyPageResultResultData {
	return s.ResultData
}

func (s *ListBatchTasksResponseBodyPageResult) SetCount(v int32) *ListBatchTasksResponseBodyPageResult {
	s.Count = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResult) SetPage(v int32) *ListBatchTasksResponseBodyPageResult {
	s.Page = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResult) SetPageSize(v int32) *ListBatchTasksResponseBodyPageResult {
	s.PageSize = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResult) SetResultData(v []*ListBatchTasksResponseBodyPageResultResultData) *ListBatchTasksResponseBodyPageResult {
	s.ResultData = v
	return s
}

func (s *ListBatchTasksResponseBodyPageResult) Validate() error {
	if s.ResultData != nil {
		for _, item := range s.ResultData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBatchTasksResponseBodyPageResultResultData struct {
	// example:
	//
	// 订单明细加工任务
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// /dwd
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 7090125821589888
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// SUCCESS
	LastSubmitStatus *string `json:"LastSubmitStatus,omitempty" xml:"LastSubmitStatus,omitempty"`
	// example:
	//
	// 3
	LastVersion *int32 `json:"LastVersion,omitempty" xml:"LastVersion,omitempty"`
	// example:
	//
	// dwd_order_detail
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// n_123456
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// dwd_order_detail
	NodeName           *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeOutputNameList []*string `json:"NodeOutputNameList,omitempty" xml:"NodeOutputNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	NodeType *int32 `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 10
	OperatorType *int32 `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 30001011
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Published   *bool   `json:"Published,omitempty" xml:"Published,omitempty"`
	Released    *bool   `json:"Released,omitempty" xml:"Released,omitempty"`
	// example:
	//
	// SUBMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBatchTasksResponseBodyPageResultResultData) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTasksResponseBodyPageResultResultData) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetDescription() *string {
	return s.Description
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetDirectory() *string {
	return s.Directory
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetFileId() *int64 {
	return s.FileId
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetLastSubmitStatus() *string {
	return s.LastSubmitStatus
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetLastVersion() *int32 {
	return s.LastVersion
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetName() *string {
	return s.Name
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetNodeId() *string {
	return s.NodeId
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetNodeName() *string {
	return s.NodeName
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetNodeOutputNameList() []*string {
	return s.NodeOutputNameList
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetNodeType() *int32 {
	return s.NodeType
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetOperatorType() *int32 {
	return s.OperatorType
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetPublished() *bool {
	return s.Published
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetReleased() *bool {
	return s.Released
}

func (s *ListBatchTasksResponseBodyPageResultResultData) GetStatus() *string {
	return s.Status
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetDescription(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.Description = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetDirectory(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.Directory = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetFileId(v int64) *ListBatchTasksResponseBodyPageResultResultData {
	s.FileId = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetLastSubmitStatus(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.LastSubmitStatus = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetLastVersion(v int32) *ListBatchTasksResponseBodyPageResultResultData {
	s.LastVersion = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetName(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.Name = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetNodeId(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.NodeId = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetNodeName(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.NodeName = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetNodeOutputNameList(v []*string) *ListBatchTasksResponseBodyPageResultResultData {
	s.NodeOutputNameList = v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetNodeType(v int32) *ListBatchTasksResponseBodyPageResultResultData {
	s.NodeType = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetOperatorType(v int32) *ListBatchTasksResponseBodyPageResultResultData {
	s.OperatorType = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetOwnerName(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.OwnerName = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetOwnerUserId(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.OwnerUserId = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetPublished(v bool) *ListBatchTasksResponseBodyPageResultResultData {
	s.Published = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetReleased(v bool) *ListBatchTasksResponseBodyPageResultResultData {
	s.Released = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) SetStatus(v string) *ListBatchTasksResponseBodyPageResultResultData {
	s.Status = &v
	return s
}

func (s *ListBatchTasksResponseBodyPageResultResultData) Validate() error {
	return dara.Validate(s)
}
