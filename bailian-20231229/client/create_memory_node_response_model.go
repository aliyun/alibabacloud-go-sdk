// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemoryNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemoryNodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemoryNodeResponseBody) *CreateMemoryNodeResponse
	GetBody() *CreateMemoryNodeResponseBody
}

type CreateMemoryNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemoryNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemoryNodeResponse) GetBody() *CreateMemoryNodeResponseBody {
	return s.Body
}

func (s *CreateMemoryNodeResponse) SetHeaders(v map[string]*string) *CreateMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryNodeResponse) SetStatusCode(v int32) *CreateMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryNodeResponse) SetBody(v *CreateMemoryNodeResponseBody) *CreateMemoryNodeResponse {
	s.Body = v
	return s
}

func (s *CreateMemoryNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
