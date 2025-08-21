// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPushRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPushRecordsResponse
	GetStatusCode() *int32
	SetBody(v *QueryPushRecordsResponseBody) *QueryPushRecordsResponse
	GetBody() *QueryPushRecordsResponseBody
}

type QueryPushRecordsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPushRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPushRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPushRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPushRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPushRecordsResponse) GetBody() *QueryPushRecordsResponseBody {
	return s.Body
}

func (s *QueryPushRecordsResponse) SetHeaders(v map[string]*string) *QueryPushRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryPushRecordsResponse) SetStatusCode(v int32) *QueryPushRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushRecordsResponse) SetBody(v *QueryPushRecordsResponseBody) *QueryPushRecordsResponse {
	s.Body = v
	return s
}

func (s *QueryPushRecordsResponse) Validate() error {
	return dara.Validate(s)
}
