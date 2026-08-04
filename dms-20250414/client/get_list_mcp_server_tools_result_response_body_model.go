// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListMcpServerToolsResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetListMcpServerToolsResultResponseBodyData) *GetListMcpServerToolsResultResponseBody
	GetData() *GetListMcpServerToolsResultResponseBodyData
	SetErrorCode(v string) *GetListMcpServerToolsResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetListMcpServerToolsResultResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetListMcpServerToolsResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetListMcpServerToolsResultResponseBody
	GetSuccess() *string
}

type GetListMcpServerToolsResultResponseBody struct {
	// The MCP Server connectivity detection result. The business status is distinguished by the State field.
	Data *GetListMcpServerToolsResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return code. The value success is returned if the request succeeds. An error code is returned if the request fails.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when a system-level request failure occurs.
	//
	// example:
	//
	// no mcp connect test record for session
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate this API call.
	//
	// example:
	//
	// 550***544
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetListMcpServerToolsResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetListMcpServerToolsResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetListMcpServerToolsResultResponseBody) GetData() *GetListMcpServerToolsResultResponseBodyData {
	return s.Data
}

func (s *GetListMcpServerToolsResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetListMcpServerToolsResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetListMcpServerToolsResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListMcpServerToolsResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetListMcpServerToolsResultResponseBody) SetData(v *GetListMcpServerToolsResultResponseBodyData) *GetListMcpServerToolsResultResponseBody {
	s.Data = v
	return s
}

func (s *GetListMcpServerToolsResultResponseBody) SetErrorCode(v string) *GetListMcpServerToolsResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBody) SetErrorMessage(v string) *GetListMcpServerToolsResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBody) SetRequestId(v string) *GetListMcpServerToolsResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBody) SetSuccess(v string) *GetListMcpServerToolsResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetListMcpServerToolsResultResponseBodyData struct {
	// Indicates whether the MCP Server is accessible. The value is true only when State is success.
	//
	// example:
	//
	// true
	Accessible *bool `json:"Accessible,omitempty" xml:"Accessible,omitempty"`
	// The detection status. Valid values:
	//
	// - pending: The detection is in progress.
	//
	// - success: The detection succeeded.
	//
	// - failed: The detection failed or timed out.
	//
	// The top-level Success field can be true in all three business states.
	//
	// example:
	//
	// success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The list of detected MCP tools. A non-empty list is returned only when State is success.
	Tools []*GetListMcpServerToolsResultResponseBodyDataTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
}

func (s GetListMcpServerToolsResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetListMcpServerToolsResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetListMcpServerToolsResultResponseBodyData) GetAccessible() *bool {
	return s.Accessible
}

func (s *GetListMcpServerToolsResultResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetListMcpServerToolsResultResponseBodyData) GetTools() []*GetListMcpServerToolsResultResponseBodyDataTools {
	return s.Tools
}

func (s *GetListMcpServerToolsResultResponseBodyData) SetAccessible(v bool) *GetListMcpServerToolsResultResponseBodyData {
	s.Accessible = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBodyData) SetState(v string) *GetListMcpServerToolsResultResponseBodyData {
	s.State = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBodyData) SetTools(v []*GetListMcpServerToolsResultResponseBodyDataTools) *GetListMcpServerToolsResultResponseBodyData {
	s.Tools = v
	return s
}

func (s *GetListMcpServerToolsResultResponseBodyData) Validate() error {
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

type GetListMcpServerToolsResultResponseBodyDataTools struct {
	// The description of the MCP tool functionality.
	//
	// example:
	//
	// query user information by user ID
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The JSON Schema string of the tool input parameters.
	//
	// example:
	//
	// {"type":"object","properties":{"userId":{"type":"string"}}}
	InputSchema *string `json:"InputSchema,omitempty" xml:"InputSchema,omitempty"`
	// The MCP tool name.
	//
	// example:
	//
	// query_user
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetListMcpServerToolsResultResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s GetListMcpServerToolsResultResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *GetListMcpServerToolsResultResponseBodyDataTools) GetDescription() *string {
	return s.Description
}

func (s *GetListMcpServerToolsResultResponseBodyDataTools) GetInputSchema() *string {
	return s.InputSchema
}

func (s *GetListMcpServerToolsResultResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *GetListMcpServerToolsResultResponseBodyDataTools) SetDescription(v string) *GetListMcpServerToolsResultResponseBodyDataTools {
	s.Description = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBodyDataTools) SetInputSchema(v string) *GetListMcpServerToolsResultResponseBodyDataTools {
	s.InputSchema = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBodyDataTools) SetName(v string) *GetListMcpServerToolsResultResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *GetListMcpServerToolsResultResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
