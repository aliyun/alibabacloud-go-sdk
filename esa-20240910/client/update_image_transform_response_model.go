// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageTransformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageTransformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageTransformResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageTransformResponseBody) *UpdateImageTransformResponse
	GetBody() *UpdateImageTransformResponseBody
}

type UpdateImageTransformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageTransformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageTransformResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageTransformResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageTransformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageTransformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageTransformResponse) GetBody() *UpdateImageTransformResponseBody {
	return s.Body
}

func (s *UpdateImageTransformResponse) SetHeaders(v map[string]*string) *UpdateImageTransformResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageTransformResponse) SetStatusCode(v int32) *UpdateImageTransformResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageTransformResponse) SetBody(v *UpdateImageTransformResponseBody) *UpdateImageTransformResponse {
	s.Body = v
	return s
}

func (s *UpdateImageTransformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
