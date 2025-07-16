// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMinutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartMinutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartMinutesResponse
	GetStatusCode() *int32
	SetBody(v *StartMinutesResponseBody) *StartMinutesResponse
	GetBody() *StartMinutesResponseBody
}

type StartMinutesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMinutesResponse) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesResponse) GoString() string {
	return s.String()
}

func (s *StartMinutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartMinutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartMinutesResponse) GetBody() *StartMinutesResponseBody {
	return s.Body
}

func (s *StartMinutesResponse) SetHeaders(v map[string]*string) *StartMinutesResponse {
	s.Headers = v
	return s
}

func (s *StartMinutesResponse) SetStatusCode(v int32) *StartMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMinutesResponse) SetBody(v *StartMinutesResponseBody) *StartMinutesResponse {
	s.Body = v
	return s
}

func (s *StartMinutesResponse) Validate() error {
	return dara.Validate(s)
}
