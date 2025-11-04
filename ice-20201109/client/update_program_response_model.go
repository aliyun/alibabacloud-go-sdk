// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProgramResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProgramResponseBody) *UpdateProgramResponse
	GetBody() *UpdateProgramResponseBody
}

type UpdateProgramResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProgramResponse) GoString() string {
	return s.String()
}

func (s *UpdateProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProgramResponse) GetBody() *UpdateProgramResponseBody {
	return s.Body
}

func (s *UpdateProgramResponse) SetHeaders(v map[string]*string) *UpdateProgramResponse {
	s.Headers = v
	return s
}

func (s *UpdateProgramResponse) SetStatusCode(v int32) *UpdateProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProgramResponse) SetBody(v *UpdateProgramResponseBody) *UpdateProgramResponse {
	s.Body = v
	return s
}

func (s *UpdateProgramResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
