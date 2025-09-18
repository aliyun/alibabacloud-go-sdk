// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCommandResponseBody
	GetAppId() *string
	SetCreateUserId(v string) *DescribeCommandResponseBody
	GetCreateUserId() *string
	SetCreateUserName(v string) *DescribeCommandResponseBody
	GetCreateUserName() *string
	SetDescription(v string) *DescribeCommandResponseBody
	GetDescription() *string
	SetDomainCode(v string) *DescribeCommandResponseBody
	GetDomainCode() *string
	SetDomainName(v string) *DescribeCommandResponseBody
	GetDomainName() *string
	SetGmtCreate(v string) *DescribeCommandResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *DescribeCommandResponseBody
	GetGmtModified() *string
	SetModifyUserId(v string) *DescribeCommandResponseBody
	GetModifyUserId() *string
	SetModifyUserName(v string) *DescribeCommandResponseBody
	GetModifyUserName() *string
	SetRequestId(v string) *DescribeCommandResponseBody
	GetRequestId() *string
	SetToolExamples(v []*DescribeCommandResponseBodyToolExamples) *DescribeCommandResponseBody
	GetToolExamples() []*DescribeCommandResponseBodyToolExamples
	SetToolId(v string) *DescribeCommandResponseBody
	GetToolId() *string
	SetToolName(v string) *DescribeCommandResponseBody
	GetToolName() *string
	SetToolParams(v []*DescribeCommandResponseBodyToolParams) *DescribeCommandResponseBody
	GetToolParams() []*DescribeCommandResponseBodyToolParams
}

type DescribeCommandResponseBody struct {
	// example:
	//
	// mm_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 232423
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// sdsd
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 56632343
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// example:
	//
	// xccvsd
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// xxx
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// xxx
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 57967
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// xxx
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// xxxx
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ToolExamples []*DescribeCommandResponseBodyToolExamples `json:"ToolExamples,omitempty" xml:"ToolExamples,omitempty" type:"Repeated"`
	// example:
	//
	// 65655
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// example:
	//
	// dsf34
	ToolName   *string                                  `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParams []*DescribeCommandResponseBodyToolParams `json:"ToolParams,omitempty" xml:"ToolParams,omitempty" type:"Repeated"`
}

func (s DescribeCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommandResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCommandResponseBody) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *DescribeCommandResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeCommandResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommandResponseBody) GetDomainCode() *string {
	return s.DomainCode
}

func (s *DescribeCommandResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCommandResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeCommandResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCommandResponseBody) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *DescribeCommandResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommandResponseBody) GetToolExamples() []*DescribeCommandResponseBodyToolExamples {
	return s.ToolExamples
}

func (s *DescribeCommandResponseBody) GetToolId() *string {
	return s.ToolId
}

func (s *DescribeCommandResponseBody) GetToolName() *string {
	return s.ToolName
}

func (s *DescribeCommandResponseBody) GetToolParams() []*DescribeCommandResponseBodyToolParams {
	return s.ToolParams
}

func (s *DescribeCommandResponseBody) SetAppId(v string) *DescribeCommandResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeCommandResponseBody) SetCreateUserId(v string) *DescribeCommandResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeCommandResponseBody) SetCreateUserName(v string) *DescribeCommandResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeCommandResponseBody) SetDescription(v string) *DescribeCommandResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCommandResponseBody) SetDomainCode(v string) *DescribeCommandResponseBody {
	s.DomainCode = &v
	return s
}

func (s *DescribeCommandResponseBody) SetDomainName(v string) *DescribeCommandResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeCommandResponseBody) SetGmtCreate(v string) *DescribeCommandResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCommandResponseBody) SetGmtModified(v string) *DescribeCommandResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeCommandResponseBody) SetModifyUserId(v string) *DescribeCommandResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeCommandResponseBody) SetModifyUserName(v string) *DescribeCommandResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeCommandResponseBody) SetRequestId(v string) *DescribeCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommandResponseBody) SetToolExamples(v []*DescribeCommandResponseBodyToolExamples) *DescribeCommandResponseBody {
	s.ToolExamples = v
	return s
}

func (s *DescribeCommandResponseBody) SetToolId(v string) *DescribeCommandResponseBody {
	s.ToolId = &v
	return s
}

func (s *DescribeCommandResponseBody) SetToolName(v string) *DescribeCommandResponseBody {
	s.ToolName = &v
	return s
}

func (s *DescribeCommandResponseBody) SetToolParams(v []*DescribeCommandResponseBodyToolParams) *DescribeCommandResponseBody {
	s.ToolParams = v
	return s
}

func (s *DescribeCommandResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCommandResponseBodyToolExamples struct {
	// example:
	//
	// xxx
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s DescribeCommandResponseBodyToolExamples) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandResponseBodyToolExamples) GoString() string {
	return s.String()
}

func (s *DescribeCommandResponseBodyToolExamples) GetQuery() *string {
	return s.Query
}

func (s *DescribeCommandResponseBodyToolExamples) SetQuery(v string) *DescribeCommandResponseBodyToolExamples {
	s.Query = &v
	return s
}

func (s *DescribeCommandResponseBodyToolExamples) Validate() error {
	return dara.Validate(s)
}

type DescribeCommandResponseBodyToolParams struct {
	// example:
	//
	// xxx
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// example:
	//
	// xxx
	ParamExample *string `json:"ParamExample,omitempty" xml:"ParamExample,omitempty"`
	// example:
	//
	// xxx
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
}

func (s DescribeCommandResponseBodyToolParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandResponseBodyToolParams) GoString() string {
	return s.String()
}

func (s *DescribeCommandResponseBodyToolParams) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *DescribeCommandResponseBodyToolParams) GetParamExample() *string {
	return s.ParamExample
}

func (s *DescribeCommandResponseBodyToolParams) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeCommandResponseBodyToolParams) SetParamDesc(v string) *DescribeCommandResponseBodyToolParams {
	s.ParamDesc = &v
	return s
}

func (s *DescribeCommandResponseBodyToolParams) SetParamExample(v string) *DescribeCommandResponseBodyToolParams {
	s.ParamExample = &v
	return s
}

func (s *DescribeCommandResponseBodyToolParams) SetParamName(v string) *DescribeCommandResponseBodyToolParams {
	s.ParamName = &v
	return s
}

func (s *DescribeCommandResponseBodyToolParams) Validate() error {
	return dara.Validate(s)
}
