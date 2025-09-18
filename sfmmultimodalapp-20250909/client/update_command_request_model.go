// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateCommandRequest
	GetAppId() *string
	SetDomainCode(v string) *UpdateCommandRequest
	GetDomainCode() *string
	SetDomainName(v string) *UpdateCommandRequest
	GetDomainName() *string
	SetToolDescription(v string) *UpdateCommandRequest
	GetToolDescription() *string
	SetToolExamples(v []*UpdateCommandRequestToolExamples) *UpdateCommandRequest
	GetToolExamples() []*UpdateCommandRequestToolExamples
	SetToolId(v string) *UpdateCommandRequest
	GetToolId() *string
	SetToolName(v string) *UpdateCommandRequest
	GetToolName() *string
	SetToolParams(v []*UpdateCommandRequestToolParams) *UpdateCommandRequest
	GetToolParams() []*UpdateCommandRequestToolParams
	SetWorkspaceId(v string) *UpdateCommandRequest
	GetWorkspaceId() *string
}

type UpdateCommandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_axaxaaa
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 659864545
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// example:
	//
	// shopping_t
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	ToolDescription *string                             `json:"ToolDescription,omitempty" xml:"ToolDescription,omitempty"`
	ToolExamples    []*UpdateCommandRequestToolExamples `json:"ToolExamples,omitempty" xml:"ToolExamples,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 8293382932xxx
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open_bx
	ToolName   *string                           `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParams []*UpdateCommandRequestToolParams `json:"ToolParams,omitempty" xml:"ToolParams,omitempty" type:"Repeated"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommandRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCommandRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *UpdateCommandRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateCommandRequest) GetToolDescription() *string {
	return s.ToolDescription
}

func (s *UpdateCommandRequest) GetToolExamples() []*UpdateCommandRequestToolExamples {
	return s.ToolExamples
}

func (s *UpdateCommandRequest) GetToolId() *string {
	return s.ToolId
}

func (s *UpdateCommandRequest) GetToolName() *string {
	return s.ToolName
}

func (s *UpdateCommandRequest) GetToolParams() []*UpdateCommandRequestToolParams {
	return s.ToolParams
}

func (s *UpdateCommandRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCommandRequest) SetAppId(v string) *UpdateCommandRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCommandRequest) SetDomainCode(v string) *UpdateCommandRequest {
	s.DomainCode = &v
	return s
}

func (s *UpdateCommandRequest) SetDomainName(v string) *UpdateCommandRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateCommandRequest) SetToolDescription(v string) *UpdateCommandRequest {
	s.ToolDescription = &v
	return s
}

func (s *UpdateCommandRequest) SetToolExamples(v []*UpdateCommandRequestToolExamples) *UpdateCommandRequest {
	s.ToolExamples = v
	return s
}

func (s *UpdateCommandRequest) SetToolId(v string) *UpdateCommandRequest {
	s.ToolId = &v
	return s
}

func (s *UpdateCommandRequest) SetToolName(v string) *UpdateCommandRequest {
	s.ToolName = &v
	return s
}

func (s *UpdateCommandRequest) SetToolParams(v []*UpdateCommandRequestToolParams) *UpdateCommandRequest {
	s.ToolParams = v
	return s
}

func (s *UpdateCommandRequest) SetWorkspaceId(v string) *UpdateCommandRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCommandRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateCommandRequestToolExamples struct {
	// example:
	//
	// 给我xxx
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateCommandRequestToolExamples) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandRequestToolExamples) GoString() string {
	return s.String()
}

func (s *UpdateCommandRequestToolExamples) GetQuery() *string {
	return s.Query
}

func (s *UpdateCommandRequestToolExamples) SetQuery(v string) *UpdateCommandRequestToolExamples {
	s.Query = &v
	return s
}

func (s *UpdateCommandRequestToolExamples) Validate() error {
	return dara.Validate(s)
}

type UpdateCommandRequestToolParams struct {
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
	// xxxx
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
}

func (s UpdateCommandRequestToolParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandRequestToolParams) GoString() string {
	return s.String()
}

func (s *UpdateCommandRequestToolParams) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *UpdateCommandRequestToolParams) GetParamExample() *string {
	return s.ParamExample
}

func (s *UpdateCommandRequestToolParams) GetParamName() *string {
	return s.ParamName
}

func (s *UpdateCommandRequestToolParams) SetParamDesc(v string) *UpdateCommandRequestToolParams {
	s.ParamDesc = &v
	return s
}

func (s *UpdateCommandRequestToolParams) SetParamExample(v string) *UpdateCommandRequestToolParams {
	s.ParamExample = &v
	return s
}

func (s *UpdateCommandRequestToolParams) SetParamName(v string) *UpdateCommandRequestToolParams {
	s.ParamName = &v
	return s
}

func (s *UpdateCommandRequestToolParams) Validate() error {
	return dara.Validate(s)
}
