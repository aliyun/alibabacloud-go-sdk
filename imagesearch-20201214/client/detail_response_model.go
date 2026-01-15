// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetailResponse
	GetStatusCode() *int32
	SetBody(v *DetailResponseBody) *DetailResponse
	GetBody() *DetailResponseBody
}

type DetailResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DetailResponse) GoString() string {
	return s.String()
}

func (s *DetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetailResponse) GetBody() *DetailResponseBody {
	return s.Body
}

func (s *DetailResponse) SetHeaders(v map[string]*string) *DetailResponse {
	s.Headers = v
	return s
}

func (s *DetailResponse) SetStatusCode(v int32) *DetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailResponse) SetBody(v *DetailResponseBody) *DetailResponse {
	s.Body = v
	return s
}

func (s *DetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
