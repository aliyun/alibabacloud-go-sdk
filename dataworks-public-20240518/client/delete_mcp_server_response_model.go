// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcpServerResponseBody) *DeleteMcpServerResponse
	GetBody() *DeleteMcpServerResponseBody
}

type DeleteMcpServerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcpServerResponse) GetBody() *DeleteMcpServerResponseBody {
	return s.Body
}

func (s *DeleteMcpServerResponse) SetHeaders(v map[string]*string) *DeleteMcpServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpServerResponse) SetStatusCode(v int32) *DeleteMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpServerResponse) SetBody(v *DeleteMcpServerResponseBody) *DeleteMcpServerResponse {
	s.Body = v
	return s
}

func (s *DeleteMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
