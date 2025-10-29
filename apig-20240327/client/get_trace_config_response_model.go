// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTraceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTraceConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetTraceConfigResponseBody) *GetTraceConfigResponse
	GetBody() *GetTraceConfigResponseBody
}

type GetTraceConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTraceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetTraceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTraceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTraceConfigResponse) GetBody() *GetTraceConfigResponseBody {
	return s.Body
}

func (s *GetTraceConfigResponse) SetHeaders(v map[string]*string) *GetTraceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetTraceConfigResponse) SetStatusCode(v int32) *GetTraceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceConfigResponse) SetBody(v *GetTraceConfigResponseBody) *GetTraceConfigResponse {
	s.Body = v
	return s
}

func (s *GetTraceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
