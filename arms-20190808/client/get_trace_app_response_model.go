// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTraceAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTraceAppResponse
	GetStatusCode() *int32
	SetBody(v *GetTraceAppResponseBody) *GetTraceAppResponse
	GetBody() *GetTraceAppResponseBody
}

type GetTraceAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppResponse) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTraceAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTraceAppResponse) GetBody() *GetTraceAppResponseBody {
	return s.Body
}

func (s *GetTraceAppResponse) SetHeaders(v map[string]*string) *GetTraceAppResponse {
	s.Headers = v
	return s
}

func (s *GetTraceAppResponse) SetStatusCode(v int32) *GetTraceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceAppResponse) SetBody(v *GetTraceAppResponseBody) *GetTraceAppResponse {
	s.Body = v
	return s
}

func (s *GetTraceAppResponse) Validate() error {
	return dara.Validate(s)
}
