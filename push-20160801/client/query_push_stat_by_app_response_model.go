// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushStatByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPushStatByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPushStatByAppResponse
	GetStatusCode() *int32
	SetBody(v *QueryPushStatByAppResponseBody) *QueryPushStatByAppResponse
	GetBody() *QueryPushStatByAppResponseBody
}

type QueryPushStatByAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPushStatByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPushStatByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByAppResponse) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPushStatByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPushStatByAppResponse) GetBody() *QueryPushStatByAppResponseBody {
	return s.Body
}

func (s *QueryPushStatByAppResponse) SetHeaders(v map[string]*string) *QueryPushStatByAppResponse {
	s.Headers = v
	return s
}

func (s *QueryPushStatByAppResponse) SetStatusCode(v int32) *QueryPushStatByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushStatByAppResponse) SetBody(v *QueryPushStatByAppResponseBody) *QueryPushStatByAppResponse {
	s.Body = v
	return s
}

func (s *QueryPushStatByAppResponse) Validate() error {
	return dara.Validate(s)
}
