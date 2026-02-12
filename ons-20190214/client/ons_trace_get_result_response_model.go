// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceGetResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTraceGetResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTraceGetResultResponse
	GetStatusCode() *int32
	SetBody(v *OnsTraceGetResultResponseBody) *OnsTraceGetResultResponse
	GetBody() *OnsTraceGetResultResponseBody
}

type OnsTraceGetResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTraceGetResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTraceGetResultResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponse) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTraceGetResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTraceGetResultResponse) GetBody() *OnsTraceGetResultResponseBody {
	return s.Body
}

func (s *OnsTraceGetResultResponse) SetHeaders(v map[string]*string) *OnsTraceGetResultResponse {
	s.Headers = v
	return s
}

func (s *OnsTraceGetResultResponse) SetStatusCode(v int32) *OnsTraceGetResultResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTraceGetResultResponse) SetBody(v *OnsTraceGetResultResponseBody) *OnsTraceGetResultResponse {
	s.Body = v
	return s
}

func (s *OnsTraceGetResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
