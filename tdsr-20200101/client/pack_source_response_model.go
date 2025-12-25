// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PackSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PackSourceResponse
	GetStatusCode() *int32
	SetBody(v *PackSourceResponseBody) *PackSourceResponse
	GetBody() *PackSourceResponseBody
}

type PackSourceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PackSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PackSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s PackSourceResponse) GoString() string {
	return s.String()
}

func (s *PackSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PackSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PackSourceResponse) GetBody() *PackSourceResponseBody {
	return s.Body
}

func (s *PackSourceResponse) SetHeaders(v map[string]*string) *PackSourceResponse {
	s.Headers = v
	return s
}

func (s *PackSourceResponse) SetStatusCode(v int32) *PackSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PackSourceResponse) SetBody(v *PackSourceResponseBody) *PackSourceResponse {
	s.Body = v
	return s
}

func (s *PackSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
