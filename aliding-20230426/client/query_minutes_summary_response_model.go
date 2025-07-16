// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMinutesSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMinutesSummaryResponse
	GetStatusCode() *int32
	SetBody(v *QueryMinutesSummaryResponseBody) *QueryMinutesSummaryResponse
	GetBody() *QueryMinutesSummaryResponseBody
}

type QueryMinutesSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMinutesSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMinutesSummaryResponse) GetBody() *QueryMinutesSummaryResponseBody {
	return s.Body
}

func (s *QueryMinutesSummaryResponse) SetHeaders(v map[string]*string) *QueryMinutesSummaryResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesSummaryResponse) SetStatusCode(v int32) *QueryMinutesSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesSummaryResponse) SetBody(v *QueryMinutesSummaryResponseBody) *QueryMinutesSummaryResponse {
	s.Body = v
	return s
}

func (s *QueryMinutesSummaryResponse) Validate() error {
	return dara.Validate(s)
}
