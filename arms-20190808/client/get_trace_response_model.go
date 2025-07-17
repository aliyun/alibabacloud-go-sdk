// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTraceResponse
	GetStatusCode() *int32
	SetBody(v *GetTraceResponseBody) *GetTraceResponse
	GetBody() *GetTraceResponseBody
}

type GetTraceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTraceResponse) GetBody() *GetTraceResponseBody {
	return s.Body
}

func (s *GetTraceResponse) SetHeaders(v map[string]*string) *GetTraceResponse {
	s.Headers = v
	return s
}

func (s *GetTraceResponse) SetStatusCode(v int32) *GetTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceResponse) SetBody(v *GetTraceResponseBody) *GetTraceResponse {
	s.Body = v
	return s
}

func (s *GetTraceResponse) Validate() error {
	return dara.Validate(s)
}
