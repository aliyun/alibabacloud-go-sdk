// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourceStaticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryResourceStaticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryResourceStaticsResponse
	GetStatusCode() *int32
	SetBody(v *QueryResourceStaticsResponseBody) *QueryResourceStaticsResponse
	GetBody() *QueryResourceStaticsResponseBody
}

type QueryResourceStaticsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResourceStaticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryResourceStaticsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceStaticsResponse) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryResourceStaticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryResourceStaticsResponse) GetBody() *QueryResourceStaticsResponseBody {
	return s.Body
}

func (s *QueryResourceStaticsResponse) SetHeaders(v map[string]*string) *QueryResourceStaticsResponse {
	s.Headers = v
	return s
}

func (s *QueryResourceStaticsResponse) SetStatusCode(v int32) *QueryResourceStaticsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResourceStaticsResponse) SetBody(v *QueryResourceStaticsResponseBody) *QueryResourceStaticsResponse {
	s.Body = v
	return s
}

func (s *QueryResourceStaticsResponse) Validate() error {
	return dara.Validate(s)
}
