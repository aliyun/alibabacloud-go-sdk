// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileResponseBody) *UpdateFileResponse
	GetBody() *UpdateFileResponseBody
}

type UpdateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileResponse) GetBody() *UpdateFileResponseBody {
	return s.Body
}

func (s *UpdateFileResponse) SetHeaders(v map[string]*string) *UpdateFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileResponse) SetStatusCode(v int32) *UpdateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileResponse) SetBody(v *UpdateFileResponseBody) *UpdateFileResponse {
	s.Body = v
	return s
}

func (s *UpdateFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
