// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateToolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateToolResponse
	GetStatusCode() *int32
	SetBody(v *ToolResult) *CreateToolResponse
	GetBody() *ToolResult
}

type CreateToolResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ToolResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateToolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateToolResponse) GoString() string {
	return s.String()
}

func (s *CreateToolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateToolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateToolResponse) GetBody() *ToolResult {
	return s.Body
}

func (s *CreateToolResponse) SetHeaders(v map[string]*string) *CreateToolResponse {
	s.Headers = v
	return s
}

func (s *CreateToolResponse) SetStatusCode(v int32) *CreateToolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateToolResponse) SetBody(v *ToolResult) *CreateToolResponse {
	s.Body = v
	return s
}

func (s *CreateToolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
