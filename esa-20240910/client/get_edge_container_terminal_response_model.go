// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerTerminalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerTerminalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerTerminalResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerTerminalResponseBody) *GetEdgeContainerTerminalResponse
	GetBody() *GetEdgeContainerTerminalResponseBody
}

type GetEdgeContainerTerminalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerTerminalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerTerminalResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerTerminalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerTerminalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerTerminalResponse) GetBody() *GetEdgeContainerTerminalResponseBody {
	return s.Body
}

func (s *GetEdgeContainerTerminalResponse) SetHeaders(v map[string]*string) *GetEdgeContainerTerminalResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerTerminalResponse) SetStatusCode(v int32) *GetEdgeContainerTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerTerminalResponse) SetBody(v *GetEdgeContainerTerminalResponseBody) *GetEdgeContainerTerminalResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerTerminalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
