// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushSchedulerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPushSchedulerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPushSchedulerListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPushSchedulerListResponseBody) *QueryPushSchedulerListResponse
	GetBody() *QueryPushSchedulerListResponseBody
}

type QueryPushSchedulerListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPushSchedulerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPushSchedulerListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPushSchedulerListResponse) GoString() string {
	return s.String()
}

func (s *QueryPushSchedulerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPushSchedulerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPushSchedulerListResponse) GetBody() *QueryPushSchedulerListResponseBody {
	return s.Body
}

func (s *QueryPushSchedulerListResponse) SetHeaders(v map[string]*string) *QueryPushSchedulerListResponse {
	s.Headers = v
	return s
}

func (s *QueryPushSchedulerListResponse) SetStatusCode(v int32) *QueryPushSchedulerListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushSchedulerListResponse) SetBody(v *QueryPushSchedulerListResponseBody) *QueryPushSchedulerListResponse {
	s.Body = v
	return s
}

func (s *QueryPushSchedulerListResponse) Validate() error {
	return dara.Validate(s)
}
