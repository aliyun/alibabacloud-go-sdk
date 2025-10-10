// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAppInfoResponseBody) *QueryAppInfoResponse
	GetBody() *QueryAppInfoResponseBody
}

type QueryAppInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAppInfoResponse) GetBody() *QueryAppInfoResponseBody {
	return s.Body
}

func (s *QueryAppInfoResponse) SetHeaders(v map[string]*string) *QueryAppInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAppInfoResponse) SetStatusCode(v int32) *QueryAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppInfoResponse) SetBody(v *QueryAppInfoResponseBody) *QueryAppInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAppInfoResponse) Validate() error {
	return dara.Validate(s)
}
