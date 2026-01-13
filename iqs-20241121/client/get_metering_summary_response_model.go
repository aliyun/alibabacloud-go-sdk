// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeteringSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMeteringSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMeteringSummaryResponse
	GetStatusCode() *int32
	SetBody(v *MeteringSummaryResult) *GetMeteringSummaryResponse
	GetBody() *MeteringSummaryResult
}

type GetMeteringSummaryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MeteringSummaryResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMeteringSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMeteringSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetMeteringSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMeteringSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMeteringSummaryResponse) GetBody() *MeteringSummaryResult {
	return s.Body
}

func (s *GetMeteringSummaryResponse) SetHeaders(v map[string]*string) *GetMeteringSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetMeteringSummaryResponse) SetStatusCode(v int32) *GetMeteringSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeteringSummaryResponse) SetBody(v *MeteringSummaryResult) *GetMeteringSummaryResponse {
	s.Body = v
	return s
}

func (s *GetMeteringSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
