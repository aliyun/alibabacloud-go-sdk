// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSamplingLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJMeterSamplingLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJMeterSamplingLogsResponse
	GetStatusCode() *int32
	SetBody(v *GetJMeterSamplingLogsResponseBody) *GetJMeterSamplingLogsResponse
	GetBody() *GetJMeterSamplingLogsResponseBody
}

type GetJMeterSamplingLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJMeterSamplingLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJMeterSamplingLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSamplingLogsResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterSamplingLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJMeterSamplingLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJMeterSamplingLogsResponse) GetBody() *GetJMeterSamplingLogsResponseBody {
	return s.Body
}

func (s *GetJMeterSamplingLogsResponse) SetHeaders(v map[string]*string) *GetJMeterSamplingLogsResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterSamplingLogsResponse) SetStatusCode(v int32) *GetJMeterSamplingLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJMeterSamplingLogsResponse) SetBody(v *GetJMeterSamplingLogsResponseBody) *GetJMeterSamplingLogsResponse {
	s.Body = v
	return s
}

func (s *GetJMeterSamplingLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
