// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsSignResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsSignResponseBody) *QuerySmsSignResponse
	GetBody() *QuerySmsSignResponseBody
}

type QuerySmsSignResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsSignResponse) GetBody() *QuerySmsSignResponseBody {
	return s.Body
}

func (s *QuerySmsSignResponse) SetHeaders(v map[string]*string) *QuerySmsSignResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsSignResponse) SetStatusCode(v int32) *QuerySmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsSignResponse) SetBody(v *QuerySmsSignResponseBody) *QuerySmsSignResponse {
	s.Body = v
	return s
}

func (s *QuerySmsSignResponse) Validate() error {
	return dara.Validate(s)
}
