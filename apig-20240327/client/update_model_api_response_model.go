// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelApiResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelApiResponseBody) *UpdateModelApiResponse
	GetBody() *UpdateModelApiResponseBody
}

type UpdateModelApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelApiResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelApiResponse) GetBody() *UpdateModelApiResponseBody {
	return s.Body
}

func (s *UpdateModelApiResponse) SetHeaders(v map[string]*string) *UpdateModelApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelApiResponse) SetStatusCode(v int32) *UpdateModelApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelApiResponse) SetBody(v *UpdateModelApiResponseBody) *UpdateModelApiResponse {
	s.Body = v
	return s
}

func (s *UpdateModelApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
