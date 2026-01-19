// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerReportSendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerReportSendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerReportSendResponse
	GetStatusCode() *int32
	SetBody(v *TriggerReportSendResponseBody) *TriggerReportSendResponse
	GetBody() *TriggerReportSendResponseBody
}

type TriggerReportSendResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerReportSendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerReportSendResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerReportSendResponse) GoString() string {
	return s.String()
}

func (s *TriggerReportSendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerReportSendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerReportSendResponse) GetBody() *TriggerReportSendResponseBody {
	return s.Body
}

func (s *TriggerReportSendResponse) SetHeaders(v map[string]*string) *TriggerReportSendResponse {
	s.Headers = v
	return s
}

func (s *TriggerReportSendResponse) SetStatusCode(v int32) *TriggerReportSendResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerReportSendResponse) SetBody(v *TriggerReportSendResponseBody) *TriggerReportSendResponse {
	s.Body = v
	return s
}

func (s *TriggerReportSendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
