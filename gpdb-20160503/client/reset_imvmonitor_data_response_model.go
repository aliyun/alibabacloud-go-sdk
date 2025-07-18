// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetIMVMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetIMVMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetIMVMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *ResetIMVMonitorDataResponseBody) *ResetIMVMonitorDataResponse
	GetBody() *ResetIMVMonitorDataResponseBody
}

type ResetIMVMonitorDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetIMVMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetIMVMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetIMVMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *ResetIMVMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetIMVMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetIMVMonitorDataResponse) GetBody() *ResetIMVMonitorDataResponseBody {
	return s.Body
}

func (s *ResetIMVMonitorDataResponse) SetHeaders(v map[string]*string) *ResetIMVMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *ResetIMVMonitorDataResponse) SetStatusCode(v int32) *ResetIMVMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetIMVMonitorDataResponse) SetBody(v *ResetIMVMonitorDataResponseBody) *ResetIMVMonitorDataResponse {
	s.Body = v
	return s
}

func (s *ResetIMVMonitorDataResponse) Validate() error {
	return dara.Validate(s)
}
