// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpProductToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpProductToolsResponseBody
	GetCode() *string
	SetData(v *ListMcpProductToolsResponseBodyData) *ListMcpProductToolsResponseBody
	GetData() *ListMcpProductToolsResponseBodyData
	SetMessage(v string) *ListMcpProductToolsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMcpProductToolsResponseBody
	GetRequestId() *string
}

type ListMcpProductToolsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListMcpProductToolsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMcpProductToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpProductToolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpProductToolsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpProductToolsResponseBody) GetData() *ListMcpProductToolsResponseBodyData {
	return s.Data
}

func (s *ListMcpProductToolsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpProductToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpProductToolsResponseBody) SetCode(v string) *ListMcpProductToolsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpProductToolsResponseBody) SetData(v *ListMcpProductToolsResponseBodyData) *ListMcpProductToolsResponseBody {
	s.Data = v
	return s
}

func (s *ListMcpProductToolsResponseBody) SetMessage(v string) *ListMcpProductToolsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpProductToolsResponseBody) SetRequestId(v string) *ListMcpProductToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpProductToolsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcpProductToolsResponseBodyData struct {
	// The list of MCP tools. Each element in the list represents an MCP tool.
	Tools []*ListMcpProductToolsResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s ListMcpProductToolsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMcpProductToolsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMcpProductToolsResponseBodyData) GetTools() []*ListMcpProductToolsResponseBodyDataTools {
	return s.Tools
}

func (s *ListMcpProductToolsResponseBodyData) SetTools(v []*ListMcpProductToolsResponseBodyDataTools) *ListMcpProductToolsResponseBodyData {
	s.Tools = v
	return s
}

func (s *ListMcpProductToolsResponseBodyData) Validate() error {
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpProductToolsResponseBodyDataTools struct {
	// The annotation information of the tool, which is a map structure. The API definition declares it only as map<any> without specifying its members or purposes.
	Annotations map[string]interface{} `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// The description of the tool.
	//
	// example:
	//
	// Functional description of the sample tool
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The input parameter structure definition of the tool, which is a map structure. The API definition declares it only as map<any> without specifying its members or constraints.
	InputSchema map[string]interface{} `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// The metadata information of the tool, which is a map structure. The API definition declares it only as map<any> without specifying its members or purposes.
	Meta map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the tool.
	//
	// example:
	//
	// test-tool
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The output result structure definition of the tool, which is a map structure. The API definition declares it only as map<any> without specifying its members or constraints.
	OutputSchema map[string]interface{} `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
	// The title of the tool.
	//
	// example:
	//
	// Sample tool
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListMcpProductToolsResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s ListMcpProductToolsResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *ListMcpProductToolsResponseBodyDataTools) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *ListMcpProductToolsResponseBodyDataTools) GetDescription() *string {
	return s.Description
}

func (s *ListMcpProductToolsResponseBodyDataTools) GetInputSchema() map[string]interface{} {
	return s.InputSchema
}

func (s *ListMcpProductToolsResponseBodyDataTools) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListMcpProductToolsResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *ListMcpProductToolsResponseBodyDataTools) GetOutputSchema() map[string]interface{} {
	return s.OutputSchema
}

func (s *ListMcpProductToolsResponseBodyDataTools) GetTitle() *string {
	return s.Title
}

func (s *ListMcpProductToolsResponseBodyDataTools) SetAnnotations(v map[string]interface{}) *ListMcpProductToolsResponseBodyDataTools {
	s.Annotations = v
	return s
}

func (s *ListMcpProductToolsResponseBodyDataTools) SetDescription(v string) *ListMcpProductToolsResponseBodyDataTools {
	s.Description = &v
	return s
}

func (s *ListMcpProductToolsResponseBodyDataTools) SetInputSchema(v map[string]interface{}) *ListMcpProductToolsResponseBodyDataTools {
	s.InputSchema = v
	return s
}

func (s *ListMcpProductToolsResponseBodyDataTools) SetMeta(v map[string]interface{}) *ListMcpProductToolsResponseBodyDataTools {
	s.Meta = v
	return s
}

func (s *ListMcpProductToolsResponseBodyDataTools) SetName(v string) *ListMcpProductToolsResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *ListMcpProductToolsResponseBodyDataTools) SetOutputSchema(v map[string]interface{}) *ListMcpProductToolsResponseBodyDataTools {
	s.OutputSchema = v
	return s
}

func (s *ListMcpProductToolsResponseBodyDataTools) SetTitle(v string) *ListMcpProductToolsResponseBodyDataTools {
	s.Title = &v
	return s
}

func (s *ListMcpProductToolsResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
