// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMinutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMinutesResponse
	GetStatusCode() *int32
	SetBody(v *QueryMinutesResponseBody) *QueryMinutesResponse
	GetBody() *QueryMinutesResponseBody
}

type QueryMinutesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMinutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMinutesResponse) GetBody() *QueryMinutesResponseBody {
	return s.Body
}

func (s *QueryMinutesResponse) SetHeaders(v map[string]*string) *QueryMinutesResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesResponse) SetStatusCode(v int32) *QueryMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesResponse) SetBody(v *QueryMinutesResponseBody) *QueryMinutesResponse {
	s.Body = v
	return s
}

func (s *QueryMinutesResponse) Validate() error {
	return dara.Validate(s)
}
