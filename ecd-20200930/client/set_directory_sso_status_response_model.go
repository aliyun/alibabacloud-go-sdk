// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDirectorySsoStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDirectorySsoStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDirectorySsoStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDirectorySsoStatusResponseBody) *SetDirectorySsoStatusResponse
	GetBody() *SetDirectorySsoStatusResponseBody
}

type SetDirectorySsoStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDirectorySsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDirectorySsoStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDirectorySsoStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDirectorySsoStatusResponse) GetBody() *SetDirectorySsoStatusResponseBody {
	return s.Body
}

func (s *SetDirectorySsoStatusResponse) SetHeaders(v map[string]*string) *SetDirectorySsoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDirectorySsoStatusResponse) SetStatusCode(v int32) *SetDirectorySsoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDirectorySsoStatusResponse) SetBody(v *SetDirectorySsoStatusResponseBody) *SetDirectorySsoStatusResponse {
	s.Body = v
	return s
}

func (s *SetDirectorySsoStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
