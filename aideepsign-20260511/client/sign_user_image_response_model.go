// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignUserImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SignUserImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SignUserImageResponse
	GetStatusCode() *int32
	SetBody(v *SignUserImageResponseBody) *SignUserImageResponse
	GetBody() *SignUserImageResponseBody
}

type SignUserImageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignUserImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignUserImageResponse) String() string {
	return dara.Prettify(s)
}

func (s SignUserImageResponse) GoString() string {
	return s.String()
}

func (s *SignUserImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SignUserImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SignUserImageResponse) GetBody() *SignUserImageResponseBody {
	return s.Body
}

func (s *SignUserImageResponse) SetHeaders(v map[string]*string) *SignUserImageResponse {
	s.Headers = v
	return s
}

func (s *SignUserImageResponse) SetStatusCode(v int32) *SignUserImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SignUserImageResponse) SetBody(v *SignUserImageResponseBody) *SignUserImageResponse {
	s.Body = v
	return s
}

func (s *SignUserImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
