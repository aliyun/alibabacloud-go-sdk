// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *GetMonitorDataResponseBody) *GetMonitorDataResponse
	GetBody() *GetMonitorDataResponseBody
}

type GetMonitorDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonitorDataResponse) GetBody() *GetMonitorDataResponseBody {
	return s.Body
}

func (s *GetMonitorDataResponse) SetHeaders(v map[string]*string) *GetMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorDataResponse) SetStatusCode(v int32) *GetMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonitorDataResponse) SetBody(v *GetMonitorDataResponseBody) *GetMonitorDataResponse {
	s.Body = v
	return s
}

func (s *GetMonitorDataResponse) Validate() error {
	return dara.Validate(s)
}
