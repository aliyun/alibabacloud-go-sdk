// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpApiResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpApiResponseBody) *UpdateHttpApiResponse
	GetBody() *UpdateHttpApiResponseBody
}

type UpdateHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpApiResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpApiResponse) GetBody() *UpdateHttpApiResponseBody {
	return s.Body
}

func (s *UpdateHttpApiResponse) SetHeaders(v map[string]*string) *UpdateHttpApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpApiResponse) SetStatusCode(v int32) *UpdateHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpApiResponse) SetBody(v *UpdateHttpApiResponseBody) *UpdateHttpApiResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
