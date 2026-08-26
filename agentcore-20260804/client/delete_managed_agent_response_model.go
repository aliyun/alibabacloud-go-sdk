// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteManagedAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteManagedAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteManagedAgentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteManagedAgentResponseBody) *DeleteManagedAgentResponse
	GetBody() *DeleteManagedAgentResponseBody
}

type DeleteManagedAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteManagedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteManagedAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteManagedAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteManagedAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteManagedAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteManagedAgentResponse) GetBody() *DeleteManagedAgentResponseBody {
	return s.Body
}

func (s *DeleteManagedAgentResponse) SetHeaders(v map[string]*string) *DeleteManagedAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteManagedAgentResponse) SetStatusCode(v int32) *DeleteManagedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteManagedAgentResponse) SetBody(v *DeleteManagedAgentResponseBody) *DeleteManagedAgentResponse {
	s.Body = v
	return s
}

func (s *DeleteManagedAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
