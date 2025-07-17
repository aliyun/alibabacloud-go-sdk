// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTraceAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTraceAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetTraceAppConfigResponseBody) *GetTraceAppConfigResponse
	GetBody() *GetTraceAppConfigResponseBody
}

type GetTraceAppConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppConfigResponse) GoString() string {
	return s.String()
}

func (s *GetTraceAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTraceAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTraceAppConfigResponse) GetBody() *GetTraceAppConfigResponseBody {
	return s.Body
}

func (s *GetTraceAppConfigResponse) SetHeaders(v map[string]*string) *GetTraceAppConfigResponse {
	s.Headers = v
	return s
}

func (s *GetTraceAppConfigResponse) SetStatusCode(v int32) *GetTraceAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceAppConfigResponse) SetBody(v *GetTraceAppConfigResponseBody) *GetTraceAppConfigResponse {
	s.Body = v
	return s
}

func (s *GetTraceAppConfigResponse) Validate() error {
	return dara.Validate(s)
}
