// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderSessionListPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrderSessionListPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrderSessionListPopResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrderSessionListPopResponseBody) *QueryOrderSessionListPopResponse
	GetBody() *QueryOrderSessionListPopResponseBody
}

type QueryOrderSessionListPopResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderSessionListPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderSessionListPopResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderSessionListPopResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrderSessionListPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrderSessionListPopResponse) GetBody() *QueryOrderSessionListPopResponseBody {
	return s.Body
}

func (s *QueryOrderSessionListPopResponse) SetHeaders(v map[string]*string) *QueryOrderSessionListPopResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderSessionListPopResponse) SetStatusCode(v int32) *QueryOrderSessionListPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderSessionListPopResponse) SetBody(v *QueryOrderSessionListPopResponseBody) *QueryOrderSessionListPopResponse {
	s.Body = v
	return s
}

func (s *QueryOrderSessionListPopResponse) Validate() error {
	return dara.Validate(s)
}
