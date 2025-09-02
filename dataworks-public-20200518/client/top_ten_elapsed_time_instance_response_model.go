// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTopTenElapsedTimeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TopTenElapsedTimeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TopTenElapsedTimeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *TopTenElapsedTimeInstanceResponseBody) *TopTenElapsedTimeInstanceResponse
	GetBody() *TopTenElapsedTimeInstanceResponseBody
}

type TopTenElapsedTimeInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TopTenElapsedTimeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TopTenElapsedTimeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s TopTenElapsedTimeInstanceResponse) GoString() string {
	return s.String()
}

func (s *TopTenElapsedTimeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TopTenElapsedTimeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TopTenElapsedTimeInstanceResponse) GetBody() *TopTenElapsedTimeInstanceResponseBody {
	return s.Body
}

func (s *TopTenElapsedTimeInstanceResponse) SetHeaders(v map[string]*string) *TopTenElapsedTimeInstanceResponse {
	s.Headers = v
	return s
}

func (s *TopTenElapsedTimeInstanceResponse) SetStatusCode(v int32) *TopTenElapsedTimeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponse) SetBody(v *TopTenElapsedTimeInstanceResponseBody) *TopTenElapsedTimeInstanceResponse {
	s.Body = v
	return s
}

func (s *TopTenElapsedTimeInstanceResponse) Validate() error {
	return dara.Validate(s)
}
