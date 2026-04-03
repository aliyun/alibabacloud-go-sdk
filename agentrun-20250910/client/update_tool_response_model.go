// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateToolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateToolResponse
	GetStatusCode() *int32
	SetBody(v *ToolResult) *UpdateToolResponse
	GetBody() *ToolResult
}

type UpdateToolResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ToolResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateToolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolResponse) GoString() string {
	return s.String()
}

func (s *UpdateToolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateToolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateToolResponse) GetBody() *ToolResult {
	return s.Body
}

func (s *UpdateToolResponse) SetHeaders(v map[string]*string) *UpdateToolResponse {
	s.Headers = v
	return s
}

func (s *UpdateToolResponse) SetStatusCode(v int32) *UpdateToolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateToolResponse) SetBody(v *ToolResult) *UpdateToolResponse {
	s.Body = v
	return s
}

func (s *UpdateToolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
