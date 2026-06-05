// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppFileResponseBody) *UpdateAppFileResponse
	GetBody() *UpdateAppFileResponseBody
}

type UpdateAppFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppFileResponse) GetBody() *UpdateAppFileResponseBody {
	return s.Body
}

func (s *UpdateAppFileResponse) SetHeaders(v map[string]*string) *UpdateAppFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppFileResponse) SetStatusCode(v int32) *UpdateAppFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppFileResponse) SetBody(v *UpdateAppFileResponseBody) *UpdateAppFileResponse {
	s.Body = v
	return s
}

func (s *UpdateAppFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
