// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceAbJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTraceAbJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTraceAbJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryTraceAbJobResponseBody) *QueryTraceAbJobResponse
	GetBody() *QueryTraceAbJobResponseBody
}

type QueryTraceAbJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTraceAbJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTraceAbJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTraceAbJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTraceAbJobResponse) GetBody() *QueryTraceAbJobResponseBody {
	return s.Body
}

func (s *QueryTraceAbJobResponse) SetHeaders(v map[string]*string) *QueryTraceAbJobResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceAbJobResponse) SetStatusCode(v int32) *QueryTraceAbJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceAbJobResponse) SetBody(v *QueryTraceAbJobResponseBody) *QueryTraceAbJobResponse {
	s.Body = v
	return s
}

func (s *QueryTraceAbJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
