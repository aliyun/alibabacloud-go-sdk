// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeVhostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeVhostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeVhostResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeVhostResponseBody) *CreateMcubeVhostResponse
	GetBody() *CreateMcubeVhostResponseBody
}

type CreateMcubeVhostResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeVhostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeVhostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeVhostResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeVhostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeVhostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeVhostResponse) GetBody() *CreateMcubeVhostResponseBody {
	return s.Body
}

func (s *CreateMcubeVhostResponse) SetHeaders(v map[string]*string) *CreateMcubeVhostResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeVhostResponse) SetStatusCode(v int32) *CreateMcubeVhostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeVhostResponse) SetBody(v *CreateMcubeVhostResponseBody) *CreateMcubeVhostResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeVhostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
