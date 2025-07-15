// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRtcAsrTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRtcAsrTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRtcAsrTasksResponse
	GetStatusCode() *int32
	SetBody(v *QueryRtcAsrTasksResponseBody) *QueryRtcAsrTasksResponse
	GetBody() *QueryRtcAsrTasksResponseBody
}

type QueryRtcAsrTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRtcAsrTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRtcAsrTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRtcAsrTasksResponse) GoString() string {
	return s.String()
}

func (s *QueryRtcAsrTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRtcAsrTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRtcAsrTasksResponse) GetBody() *QueryRtcAsrTasksResponseBody {
	return s.Body
}

func (s *QueryRtcAsrTasksResponse) SetHeaders(v map[string]*string) *QueryRtcAsrTasksResponse {
	s.Headers = v
	return s
}

func (s *QueryRtcAsrTasksResponse) SetStatusCode(v int32) *QueryRtcAsrTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRtcAsrTasksResponse) SetBody(v *QueryRtcAsrTasksResponseBody) *QueryRtcAsrTasksResponse {
	s.Body = v
	return s
}

func (s *QueryRtcAsrTasksResponse) Validate() error {
	return dara.Validate(s)
}
