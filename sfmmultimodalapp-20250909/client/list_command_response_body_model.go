// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCommandResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCommandResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCommandResponseBody
	GetRequestId() *string
	SetToolInfoList(v []*ListCommandResponseBodyToolInfoList) *ListCommandResponseBody
	GetToolInfoList() []*ListCommandResponseBodyToolInfoList
	SetTotalCount(v int32) *ListCommandResponseBody
	GetTotalCount() *int32
}

type ListCommandResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxx
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ToolInfoList []*ListCommandResponseBodyToolInfoList `json:"ToolInfoList,omitempty" xml:"ToolInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommandResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommandResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCommandResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommandResponseBody) GetToolInfoList() []*ListCommandResponseBodyToolInfoList {
	return s.ToolInfoList
}

func (s *ListCommandResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCommandResponseBody) SetPageNumber(v int32) *ListCommandResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommandResponseBody) SetPageSize(v int32) *ListCommandResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommandResponseBody) SetRequestId(v string) *ListCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommandResponseBody) SetToolInfoList(v []*ListCommandResponseBodyToolInfoList) *ListCommandResponseBody {
	s.ToolInfoList = v
	return s
}

func (s *ListCommandResponseBody) SetTotalCount(v int32) *ListCommandResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCommandResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCommandResponseBodyToolInfoList struct {
	// example:
	//
	// mm-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 676776778678
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xxx
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 44574578797
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// example:
	//
	// xxh_xx
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// xxx
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// xxxx
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 2334346345
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// xxx
	ModifyUserName *string                                            `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	ToolExamples   []*ListCommandResponseBodyToolInfoListToolExamples `json:"ToolExamples,omitempty" xml:"ToolExamples,omitempty" type:"Repeated"`
	// example:
	//
	// 6734396796
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// example:
	//
	// sxxxx
	ToolName   *string                                          `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParams []*ListCommandResponseBodyToolInfoListToolParams `json:"ToolParams,omitempty" xml:"ToolParams,omitempty" type:"Repeated"`
}

func (s ListCommandResponseBodyToolInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListCommandResponseBodyToolInfoList) GoString() string {
	return s.String()
}

func (s *ListCommandResponseBodyToolInfoList) GetAppId() *string {
	return s.AppId
}

func (s *ListCommandResponseBodyToolInfoList) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListCommandResponseBodyToolInfoList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListCommandResponseBodyToolInfoList) GetDescription() *string {
	return s.Description
}

func (s *ListCommandResponseBodyToolInfoList) GetDomainCode() *string {
	return s.DomainCode
}

func (s *ListCommandResponseBodyToolInfoList) GetDomainName() *string {
	return s.DomainName
}

func (s *ListCommandResponseBodyToolInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListCommandResponseBodyToolInfoList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListCommandResponseBodyToolInfoList) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *ListCommandResponseBodyToolInfoList) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListCommandResponseBodyToolInfoList) GetToolExamples() []*ListCommandResponseBodyToolInfoListToolExamples {
	return s.ToolExamples
}

func (s *ListCommandResponseBodyToolInfoList) GetToolId() *string {
	return s.ToolId
}

func (s *ListCommandResponseBodyToolInfoList) GetToolName() *string {
	return s.ToolName
}

func (s *ListCommandResponseBodyToolInfoList) GetToolParams() []*ListCommandResponseBodyToolInfoListToolParams {
	return s.ToolParams
}

func (s *ListCommandResponseBodyToolInfoList) SetAppId(v string) *ListCommandResponseBodyToolInfoList {
	s.AppId = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetCreateUserId(v string) *ListCommandResponseBodyToolInfoList {
	s.CreateUserId = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetCreateUserName(v string) *ListCommandResponseBodyToolInfoList {
	s.CreateUserName = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetDescription(v string) *ListCommandResponseBodyToolInfoList {
	s.Description = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetDomainCode(v string) *ListCommandResponseBodyToolInfoList {
	s.DomainCode = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetDomainName(v string) *ListCommandResponseBodyToolInfoList {
	s.DomainName = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetGmtCreate(v string) *ListCommandResponseBodyToolInfoList {
	s.GmtCreate = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetGmtModified(v string) *ListCommandResponseBodyToolInfoList {
	s.GmtModified = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetModifyUserId(v string) *ListCommandResponseBodyToolInfoList {
	s.ModifyUserId = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetModifyUserName(v string) *ListCommandResponseBodyToolInfoList {
	s.ModifyUserName = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetToolExamples(v []*ListCommandResponseBodyToolInfoListToolExamples) *ListCommandResponseBodyToolInfoList {
	s.ToolExamples = v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetToolId(v string) *ListCommandResponseBodyToolInfoList {
	s.ToolId = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetToolName(v string) *ListCommandResponseBodyToolInfoList {
	s.ToolName = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) SetToolParams(v []*ListCommandResponseBodyToolInfoListToolParams) *ListCommandResponseBodyToolInfoList {
	s.ToolParams = v
	return s
}

func (s *ListCommandResponseBodyToolInfoList) Validate() error {
	return dara.Validate(s)
}

type ListCommandResponseBodyToolInfoListToolExamples struct {
	// example:
	//
	// xxx
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListCommandResponseBodyToolInfoListToolExamples) String() string {
	return dara.Prettify(s)
}

func (s ListCommandResponseBodyToolInfoListToolExamples) GoString() string {
	return s.String()
}

func (s *ListCommandResponseBodyToolInfoListToolExamples) GetQuery() *string {
	return s.Query
}

func (s *ListCommandResponseBodyToolInfoListToolExamples) SetQuery(v string) *ListCommandResponseBodyToolInfoListToolExamples {
	s.Query = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoListToolExamples) Validate() error {
	return dara.Validate(s)
}

type ListCommandResponseBodyToolInfoListToolParams struct {
	// example:
	//
	// xxx
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// example:
	//
	// xx
	ParamExample *string `json:"ParamExample,omitempty" xml:"ParamExample,omitempty"`
	// example:
	//
	// xxxx
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
}

func (s ListCommandResponseBodyToolInfoListToolParams) String() string {
	return dara.Prettify(s)
}

func (s ListCommandResponseBodyToolInfoListToolParams) GoString() string {
	return s.String()
}

func (s *ListCommandResponseBodyToolInfoListToolParams) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *ListCommandResponseBodyToolInfoListToolParams) GetParamExample() *string {
	return s.ParamExample
}

func (s *ListCommandResponseBodyToolInfoListToolParams) GetParamName() *string {
	return s.ParamName
}

func (s *ListCommandResponseBodyToolInfoListToolParams) SetParamDesc(v string) *ListCommandResponseBodyToolInfoListToolParams {
	s.ParamDesc = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoListToolParams) SetParamExample(v string) *ListCommandResponseBodyToolInfoListToolParams {
	s.ParamExample = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoListToolParams) SetParamName(v string) *ListCommandResponseBodyToolInfoListToolParams {
	s.ParamName = &v
	return s
}

func (s *ListCommandResponseBodyToolInfoListToolParams) Validate() error {
	return dara.Validate(s)
}
