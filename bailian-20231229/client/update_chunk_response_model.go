// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChunkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChunkResponseBody) *UpdateChunkResponse
	GetBody() *UpdateChunkResponseBody
}

type UpdateChunkResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChunkResponse) GoString() string {
	return s.String()
}

func (s *UpdateChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChunkResponse) GetBody() *UpdateChunkResponseBody {
	return s.Body
}

func (s *UpdateChunkResponse) SetHeaders(v map[string]*string) *UpdateChunkResponse {
	s.Headers = v
	return s
}

func (s *UpdateChunkResponse) SetStatusCode(v int32) *UpdateChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChunkResponse) SetBody(v *UpdateChunkResponseBody) *UpdateChunkResponse {
	s.Body = v
	return s
}

func (s *UpdateChunkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
