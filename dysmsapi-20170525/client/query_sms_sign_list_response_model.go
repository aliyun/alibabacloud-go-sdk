// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsSignListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsSignListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsSignListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsSignListResponseBody) *QuerySmsSignListResponse
	GetBody() *QuerySmsSignListResponseBody
}

type QuerySmsSignListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsSignListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsSignListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignListResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsSignListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsSignListResponse) GetBody() *QuerySmsSignListResponseBody {
	return s.Body
}

func (s *QuerySmsSignListResponse) SetHeaders(v map[string]*string) *QuerySmsSignListResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsSignListResponse) SetStatusCode(v int32) *QuerySmsSignListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsSignListResponse) SetBody(v *QuerySmsSignListResponseBody) *QuerySmsSignListResponse {
	s.Body = v
	return s
}

func (s *QuerySmsSignListResponse) Validate() error {
	return dara.Validate(s)
}
