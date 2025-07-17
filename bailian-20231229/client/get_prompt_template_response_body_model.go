// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetPromptTemplateResponseBody
	GetContent() *string
	SetName(v string) *GetPromptTemplateResponseBody
	GetName() *string
	SetPromptTemplateId(v string) *GetPromptTemplateResponseBody
	GetPromptTemplateId() *string
	SetRequestId(v string) *GetPromptTemplateResponseBody
	GetRequestId() *string
	SetVariables(v []*string) *GetPromptTemplateResponseBody
	GetVariables() []*string
	SetWorkspaceId(v string) *GetPromptTemplateResponseBody
	GetWorkspaceId() *string
}

type GetPromptTemplateResponseBody struct {
	// The template content.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The template name.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 6e49109bfeb94a39bb268f4e483ccxxx
	PromptTemplateId *string `json:"promptTemplateId,omitempty" xml:"promptTemplateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The variables of the template.
	//
	// example:
	//
	// ["theme"]
	Variables []*string `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// llm-us9hjmt32nysdxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetPromptTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetPromptTemplateResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetPromptTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *GetPromptTemplateResponseBody) GetPromptTemplateId() *string {
	return s.PromptTemplateId
}

func (s *GetPromptTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPromptTemplateResponseBody) GetVariables() []*string {
	return s.Variables
}

func (s *GetPromptTemplateResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPromptTemplateResponseBody) SetContent(v string) *GetPromptTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetPromptTemplateResponseBody) SetName(v string) *GetPromptTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *GetPromptTemplateResponseBody) SetPromptTemplateId(v string) *GetPromptTemplateResponseBody {
	s.PromptTemplateId = &v
	return s
}

func (s *GetPromptTemplateResponseBody) SetRequestId(v string) *GetPromptTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPromptTemplateResponseBody) SetVariables(v []*string) *GetPromptTemplateResponseBody {
	s.Variables = v
	return s
}

func (s *GetPromptTemplateResponseBody) SetWorkspaceId(v string) *GetPromptTemplateResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetPromptTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
