// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTempFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTempFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTempFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTempFileResponseBody) *UpdateTempFileResponse
	GetBody() *UpdateTempFileResponseBody
}

type UpdateTempFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTempFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTempFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTempFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateTempFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTempFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTempFileResponse) GetBody() *UpdateTempFileResponseBody {
	return s.Body
}

func (s *UpdateTempFileResponse) SetHeaders(v map[string]*string) *UpdateTempFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateTempFileResponse) SetStatusCode(v int32) *UpdateTempFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTempFileResponse) SetBody(v *UpdateTempFileResponseBody) *UpdateTempFileResponse {
	s.Body = v
	return s
}

func (s *UpdateTempFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
