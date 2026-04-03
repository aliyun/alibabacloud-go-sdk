// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteToolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteToolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteToolResponse
	GetStatusCode() *int32
	SetBody(v *ToolResult) *DeleteToolResponse
	GetBody() *ToolResult
}

type DeleteToolResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ToolResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteToolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteToolResponse) GoString() string {
	return s.String()
}

func (s *DeleteToolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteToolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteToolResponse) GetBody() *ToolResult {
	return s.Body
}

func (s *DeleteToolResponse) SetHeaders(v map[string]*string) *DeleteToolResponse {
	s.Headers = v
	return s
}

func (s *DeleteToolResponse) SetStatusCode(v int32) *DeleteToolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteToolResponse) SetBody(v *ToolResult) *DeleteToolResponse {
	s.Body = v
	return s
}

func (s *DeleteToolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
