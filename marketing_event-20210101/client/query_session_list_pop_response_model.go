// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionListPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySessionListPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySessionListPopResponse
	GetStatusCode() *int32
	SetBody(v *QuerySessionListPopResponseBody) *QuerySessionListPopResponse
	GetBody() *QuerySessionListPopResponseBody
}

type QuerySessionListPopResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionListPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionListPopResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionListPopResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySessionListPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySessionListPopResponse) GetBody() *QuerySessionListPopResponseBody {
	return s.Body
}

func (s *QuerySessionListPopResponse) SetHeaders(v map[string]*string) *QuerySessionListPopResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionListPopResponse) SetStatusCode(v int32) *QuerySessionListPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionListPopResponse) SetBody(v *QuerySessionListPopResponseBody) *QuerySessionListPopResponse {
	s.Body = v
	return s
}

func (s *QuerySessionListPopResponse) Validate() error {
	return dara.Validate(s)
}
