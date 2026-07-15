// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaResponseBody) *UpdateMediaResponse
	GetBody() *UpdateMediaResponseBody
}

type UpdateMediaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaResponse) GetBody() *UpdateMediaResponseBody {
	return s.Body
}

func (s *UpdateMediaResponse) SetHeaders(v map[string]*string) *UpdateMediaResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaResponse) SetStatusCode(v int32) *UpdateMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaResponse) SetBody(v *UpdateMediaResponseBody) *UpdateMediaResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
