// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsMessageTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsMessageTraceResponse
	GetStatusCode() *int32
	SetBody(v *OnsMessageTraceResponseBody) *OnsMessageTraceResponse
	GetBody() *OnsMessageTraceResponseBody
}

type OnsMessageTraceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageTraceResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsMessageTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsMessageTraceResponse) GetBody() *OnsMessageTraceResponseBody {
	return s.Body
}

func (s *OnsMessageTraceResponse) SetHeaders(v map[string]*string) *OnsMessageTraceResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageTraceResponse) SetStatusCode(v int32) *OnsMessageTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageTraceResponse) SetBody(v *OnsMessageTraceResponseBody) *OnsMessageTraceResponse {
	s.Body = v
	return s
}

func (s *OnsMessageTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
