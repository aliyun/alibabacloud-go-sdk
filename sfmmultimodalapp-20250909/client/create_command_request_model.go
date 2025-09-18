// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateCommandRequest
	GetAppId() *string
	SetDomainCode(v string) *CreateCommandRequest
	GetDomainCode() *string
	SetDomainName(v string) *CreateCommandRequest
	GetDomainName() *string
	SetToolDescription(v string) *CreateCommandRequest
	GetToolDescription() *string
	SetToolExamples(v []*CreateCommandRequestToolExamples) *CreateCommandRequest
	GetToolExamples() []*CreateCommandRequestToolExamples
	SetToolName(v string) *CreateCommandRequest
	GetToolName() *string
	SetToolParams(v []*CreateCommandRequestToolParams) *CreateCommandRequest
	GetToolParams() []*CreateCommandRequestToolParams
	SetWorkspaceId(v string) *CreateCommandRequest
	GetWorkspaceId() *string
}

type CreateCommandRequest struct {
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
	// This parameter is required.
	//
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
	ToolExamples    []*CreateCommandRequestToolExamples `json:"ToolExamples,omitempty" xml:"ToolExamples,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// open_bx
	ToolName   *string                           `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParams []*CreateCommandRequestToolParams `json:"ToolParams,omitempty" xml:"ToolParams,omitempty" type:"Repeated"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequest) GoString() string {
	return s.String()
}

func (s *CreateCommandRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateCommandRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *CreateCommandRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCommandRequest) GetToolDescription() *string {
	return s.ToolDescription
}

func (s *CreateCommandRequest) GetToolExamples() []*CreateCommandRequestToolExamples {
	return s.ToolExamples
}

func (s *CreateCommandRequest) GetToolName() *string {
	return s.ToolName
}

func (s *CreateCommandRequest) GetToolParams() []*CreateCommandRequestToolParams {
	return s.ToolParams
}

func (s *CreateCommandRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCommandRequest) SetAppId(v string) *CreateCommandRequest {
	s.AppId = &v
	return s
}

func (s *CreateCommandRequest) SetDomainCode(v string) *CreateCommandRequest {
	s.DomainCode = &v
	return s
}

func (s *CreateCommandRequest) SetDomainName(v string) *CreateCommandRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCommandRequest) SetToolDescription(v string) *CreateCommandRequest {
	s.ToolDescription = &v
	return s
}

func (s *CreateCommandRequest) SetToolExamples(v []*CreateCommandRequestToolExamples) *CreateCommandRequest {
	s.ToolExamples = v
	return s
}

func (s *CreateCommandRequest) SetToolName(v string) *CreateCommandRequest {
	s.ToolName = &v
	return s
}

func (s *CreateCommandRequest) SetToolParams(v []*CreateCommandRequestToolParams) *CreateCommandRequest {
	s.ToolParams = v
	return s
}

func (s *CreateCommandRequest) SetWorkspaceId(v string) *CreateCommandRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCommandRequest) Validate() error {
	return dara.Validate(s)
}

type CreateCommandRequestToolExamples struct {
	// example:
	//
	// 给我xxx
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateCommandRequestToolExamples) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequestToolExamples) GoString() string {
	return s.String()
}

func (s *CreateCommandRequestToolExamples) GetQuery() *string {
	return s.Query
}

func (s *CreateCommandRequestToolExamples) SetQuery(v string) *CreateCommandRequestToolExamples {
	s.Query = &v
	return s
}

func (s *CreateCommandRequestToolExamples) Validate() error {
	return dara.Validate(s)
}

type CreateCommandRequestToolParams struct {
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

func (s CreateCommandRequestToolParams) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequestToolParams) GoString() string {
	return s.String()
}

func (s *CreateCommandRequestToolParams) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *CreateCommandRequestToolParams) GetParamExample() *string {
	return s.ParamExample
}

func (s *CreateCommandRequestToolParams) GetParamName() *string {
	return s.ParamName
}

func (s *CreateCommandRequestToolParams) SetParamDesc(v string) *CreateCommandRequestToolParams {
	s.ParamDesc = &v
	return s
}

func (s *CreateCommandRequestToolParams) SetParamExample(v string) *CreateCommandRequestToolParams {
	s.ParamExample = &v
	return s
}

func (s *CreateCommandRequestToolParams) SetParamName(v string) *CreateCommandRequestToolParams {
	s.ParamName = &v
	return s
}

func (s *CreateCommandRequestToolParams) Validate() error {
	return dara.Validate(s)
}
