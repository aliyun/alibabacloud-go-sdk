// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseBatchTaskDependencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ParseBatchTaskDependencyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ParseBatchTaskDependencyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ParseBatchTaskDependencyResponseBody
	GetMessage() *string
	SetParseResult(v *ParseBatchTaskDependencyResponseBodyParseResult) *ParseBatchTaskDependencyResponseBody
	GetParseResult() *ParseBatchTaskDependencyResponseBodyParseResult
	SetRequestId(v string) *ParseBatchTaskDependencyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ParseBatchTaskDependencyResponseBody
	GetSuccess() *bool
}

type ParseBatchTaskDependencyResponseBody struct {
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
	Message     *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	ParseResult *ParseBatchTaskDependencyResponseBodyParseResult `json:"ParseResult,omitempty" xml:"ParseResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ParseBatchTaskDependencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyResponseBody) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ParseBatchTaskDependencyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ParseBatchTaskDependencyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ParseBatchTaskDependencyResponseBody) GetParseResult() *ParseBatchTaskDependencyResponseBodyParseResult {
	return s.ParseResult
}

func (s *ParseBatchTaskDependencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ParseBatchTaskDependencyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ParseBatchTaskDependencyResponseBody) SetCode(v string) *ParseBatchTaskDependencyResponseBody {
	s.Code = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBody) SetHttpStatusCode(v int32) *ParseBatchTaskDependencyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBody) SetMessage(v string) *ParseBatchTaskDependencyResponseBody {
	s.Message = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBody) SetParseResult(v *ParseBatchTaskDependencyResponseBodyParseResult) *ParseBatchTaskDependencyResponseBody {
	s.ParseResult = v
	return s
}

func (s *ParseBatchTaskDependencyResponseBody) SetRequestId(v string) *ParseBatchTaskDependencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBody) SetSuccess(v bool) *ParseBatchTaskDependencyResponseBody {
	s.Success = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ParseBatchTaskDependencyResponseBodyParseResult struct {
	DependNodeList []*ParseBatchTaskDependencyResponseBodyParseResultDependNodeList `json:"DependNodeList,omitempty" xml:"DependNodeList,omitempty" type:"Repeated"`
}

func (s ParseBatchTaskDependencyResponseBodyParseResult) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyResponseBodyParseResult) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyResponseBodyParseResult) GetDependNodeList() []*ParseBatchTaskDependencyResponseBodyParseResultDependNodeList {
	return s.DependNodeList
}

func (s *ParseBatchTaskDependencyResponseBodyParseResult) SetDependNodeList(v []*ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) *ParseBatchTaskDependencyResponseBodyParseResult {
	s.DependNodeList = v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResult) Validate() error {
	return dara.Validate(s)
}

type ParseBatchTaskDependencyResponseBodyParseResultDependNodeList struct {
	// example:
	//
	// input
	NodeIoType           *string                                                                              `json:"NodeIoType,omitempty" xml:"NodeIoType,omitempty"`
	ScheduleNodeInfoList []*ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList `json:"ScheduleNodeInfoList,omitempty" xml:"ScheduleNodeInfoList,omitempty" type:"Repeated"`
}

func (s ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) GetNodeIoType() *string {
	return s.NodeIoType
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) GetScheduleNodeInfoList() []*ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	return s.ScheduleNodeInfoList
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) SetNodeIoType(v string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeList {
	s.NodeIoType = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) SetScheduleNodeInfoList(v []*ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeList {
	s.ScheduleNodeInfoList = v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeList) Validate() error {
	return dara.Validate(s)
}

type ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList struct {
	FieldList []*string `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
	// example:
	//
	// n_11013121
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// test11
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// t_test
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 201122301
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// t_test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GetFieldList() []*string {
	return s.FieldList
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GetNodeId() *string {
	return s.NodeId
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GetNodeName() *string {
	return s.NodeName
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GetOutputName() *string {
	return s.OutputName
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) GetTableName() *string {
	return s.TableName
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) SetFieldList(v []*string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	s.FieldList = v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) SetNodeId(v string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	s.NodeId = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) SetNodeName(v string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	s.NodeName = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) SetOutputName(v string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	s.OutputName = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) SetOwnerName(v string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	s.OwnerName = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) SetOwnerUserId(v string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	s.OwnerUserId = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) SetTableName(v string) *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList {
	s.TableName = &v
	return s
}

func (s *ParseBatchTaskDependencyResponseBodyParseResultDependNodeListScheduleNodeInfoList) Validate() error {
	return dara.Validate(s)
}
