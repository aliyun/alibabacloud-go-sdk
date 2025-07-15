// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *StartDesktopsResponseBody) *StartDesktopsResponse
	GetBody() *StartDesktopsResponseBody
}

type StartDesktopsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDesktopsResponse) GetBody() *StartDesktopsResponseBody {
	return s.Body
}

func (s *StartDesktopsResponse) SetHeaders(v map[string]*string) *StartDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StartDesktopsResponse) SetStatusCode(v int32) *StartDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDesktopsResponse) SetBody(v *StartDesktopsResponseBody) *StartDesktopsResponse {
	s.Body = v
	return s
}

func (s *StartDesktopsResponse) Validate() error {
	return dara.Validate(s)
}
