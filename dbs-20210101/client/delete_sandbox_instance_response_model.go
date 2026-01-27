// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSandboxInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSandboxInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSandboxInstanceResponseBody) *DeleteSandboxInstanceResponse
	GetBody() *DeleteSandboxInstanceResponseBody
}

type DeleteSandboxInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSandboxInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSandboxInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSandboxInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSandboxInstanceResponse) GetBody() *DeleteSandboxInstanceResponseBody {
	return s.Body
}

func (s *DeleteSandboxInstanceResponse) SetHeaders(v map[string]*string) *DeleteSandboxInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSandboxInstanceResponse) SetStatusCode(v int32) *DeleteSandboxInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSandboxInstanceResponse) SetBody(v *DeleteSandboxInstanceResponseBody) *DeleteSandboxInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteSandboxInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
