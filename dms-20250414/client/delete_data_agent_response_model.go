// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAgentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAgentResponseBody) *DeleteDataAgentResponse
	GetBody() *DeleteDataAgentResponseBody
}

type DeleteDataAgentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAgentResponse) GetBody() *DeleteDataAgentResponseBody {
	return s.Body
}

func (s *DeleteDataAgentResponse) SetHeaders(v map[string]*string) *DeleteDataAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAgentResponse) SetStatusCode(v int32) *DeleteDataAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAgentResponse) SetBody(v *DeleteDataAgentResponseBody) *DeleteDataAgentResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
