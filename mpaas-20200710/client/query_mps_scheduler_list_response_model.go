// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMpsSchedulerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMpsSchedulerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMpsSchedulerListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMpsSchedulerListResponseBody) *QueryMpsSchedulerListResponse
	GetBody() *QueryMpsSchedulerListResponseBody
}

type QueryMpsSchedulerListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMpsSchedulerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMpsSchedulerListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMpsSchedulerListResponse) GoString() string {
	return s.String()
}

func (s *QueryMpsSchedulerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMpsSchedulerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMpsSchedulerListResponse) GetBody() *QueryMpsSchedulerListResponseBody {
	return s.Body
}

func (s *QueryMpsSchedulerListResponse) SetHeaders(v map[string]*string) *QueryMpsSchedulerListResponse {
	s.Headers = v
	return s
}

func (s *QueryMpsSchedulerListResponse) SetStatusCode(v int32) *QueryMpsSchedulerListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMpsSchedulerListResponse) SetBody(v *QueryMpsSchedulerListResponseBody) *QueryMpsSchedulerListResponse {
	s.Body = v
	return s
}

func (s *QueryMpsSchedulerListResponse) Validate() error {
	return dara.Validate(s)
}
