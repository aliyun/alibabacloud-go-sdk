// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemoryNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemoryNodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMemoryNodeResponseBody) *UpdateMemoryNodeResponse
	GetBody() *UpdateMemoryNodeResponseBody
}

type UpdateMemoryNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemoryNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemoryNodeResponse) GetBody() *UpdateMemoryNodeResponseBody {
	return s.Body
}

func (s *UpdateMemoryNodeResponse) SetHeaders(v map[string]*string) *UpdateMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryNodeResponse) SetStatusCode(v int32) *UpdateMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryNodeResponse) SetBody(v *UpdateMemoryNodeResponseBody) *UpdateMemoryNodeResponse {
	s.Body = v
	return s
}

func (s *UpdateMemoryNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
