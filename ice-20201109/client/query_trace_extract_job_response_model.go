// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceExtractJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTraceExtractJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTraceExtractJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryTraceExtractJobResponseBody) *QueryTraceExtractJobResponse
	GetBody() *QueryTraceExtractJobResponseBody
}

type QueryTraceExtractJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTraceExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTraceExtractJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceExtractJobResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTraceExtractJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTraceExtractJobResponse) GetBody() *QueryTraceExtractJobResponseBody {
	return s.Body
}

func (s *QueryTraceExtractJobResponse) SetHeaders(v map[string]*string) *QueryTraceExtractJobResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceExtractJobResponse) SetStatusCode(v int32) *QueryTraceExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceExtractJobResponse) SetBody(v *QueryTraceExtractJobResponseBody) *QueryTraceExtractJobResponse {
	s.Body = v
	return s
}

func (s *QueryTraceExtractJobResponse) Validate() error {
	return dara.Validate(s)
}
