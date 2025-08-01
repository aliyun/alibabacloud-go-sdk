// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlbSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySlbSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySlbSpecResponse
	GetStatusCode() *int32
	SetBody(v *QuerySlbSpecResponseBody) *QuerySlbSpecResponse
	GetBody() *QuerySlbSpecResponseBody
}

type QuerySlbSpecResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySlbSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySlbSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySlbSpecResponse) GoString() string {
	return s.String()
}

func (s *QuerySlbSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySlbSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySlbSpecResponse) GetBody() *QuerySlbSpecResponseBody {
	return s.Body
}

func (s *QuerySlbSpecResponse) SetHeaders(v map[string]*string) *QuerySlbSpecResponse {
	s.Headers = v
	return s
}

func (s *QuerySlbSpecResponse) SetStatusCode(v int32) *QuerySlbSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySlbSpecResponse) SetBody(v *QuerySlbSpecResponseBody) *QuerySlbSpecResponse {
	s.Body = v
	return s
}

func (s *QuerySlbSpecResponse) Validate() error {
	return dara.Validate(s)
}
