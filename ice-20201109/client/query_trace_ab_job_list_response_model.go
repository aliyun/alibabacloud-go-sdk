// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceAbJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTraceAbJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTraceAbJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryTraceAbJobListResponseBody) *QueryTraceAbJobListResponse
	GetBody() *QueryTraceAbJobListResponseBody
}

type QueryTraceAbJobListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTraceAbJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTraceAbJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTraceAbJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTraceAbJobListResponse) GetBody() *QueryTraceAbJobListResponseBody {
	return s.Body
}

func (s *QueryTraceAbJobListResponse) SetHeaders(v map[string]*string) *QueryTraceAbJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceAbJobListResponse) SetStatusCode(v int32) *QueryTraceAbJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceAbJobListResponse) SetBody(v *QueryTraceAbJobListResponseBody) *QueryTraceAbJobListResponse {
	s.Body = v
	return s
}

func (s *QueryTraceAbJobListResponse) Validate() error {
	return dara.Validate(s)
}
