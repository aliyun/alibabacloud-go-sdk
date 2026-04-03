// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetToolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetToolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetToolResponse
	GetStatusCode() *int32
	SetBody(v *ToolResult) *GetToolResponse
	GetBody() *ToolResult
}

type GetToolResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ToolResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetToolResponse) String() string {
	return dara.Prettify(s)
}

func (s GetToolResponse) GoString() string {
	return s.String()
}

func (s *GetToolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetToolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetToolResponse) GetBody() *ToolResult {
	return s.Body
}

func (s *GetToolResponse) SetHeaders(v map[string]*string) *GetToolResponse {
	s.Headers = v
	return s
}

func (s *GetToolResponse) SetStatusCode(v int32) *GetToolResponse {
	s.StatusCode = &v
	return s
}

func (s *GetToolResponse) SetBody(v *ToolResult) *GetToolResponse {
	s.Body = v
	return s
}

func (s *GetToolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
