// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMinutesTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMinutesTextResponse
	GetStatusCode() *int32
	SetBody(v *QueryMinutesTextResponseBody) *QueryMinutesTextResponse
	GetBody() *QueryMinutesTextResponseBody
}

type QueryMinutesTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesTextResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMinutesTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMinutesTextResponse) GetBody() *QueryMinutesTextResponseBody {
	return s.Body
}

func (s *QueryMinutesTextResponse) SetHeaders(v map[string]*string) *QueryMinutesTextResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesTextResponse) SetStatusCode(v int32) *QueryMinutesTextResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesTextResponse) SetBody(v *QueryMinutesTextResponseBody) *QueryMinutesTextResponse {
	s.Body = v
	return s
}

func (s *QueryMinutesTextResponse) Validate() error {
	return dara.Validate(s)
}
