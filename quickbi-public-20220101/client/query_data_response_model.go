// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryDataResponseBody) *QueryDataResponse
	GetBody() *QueryDataResponseBody
}

type QueryDataResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDataResponse) GetBody() *QueryDataResponseBody {
	return s.Body
}

func (s *QueryDataResponse) SetHeaders(v map[string]*string) *QueryDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDataResponse) SetStatusCode(v int32) *QueryDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataResponse) SetBody(v *QueryDataResponseBody) *QueryDataResponse {
	s.Body = v
	return s
}

func (s *QueryDataResponse) Validate() error {
	return dara.Validate(s)
}
