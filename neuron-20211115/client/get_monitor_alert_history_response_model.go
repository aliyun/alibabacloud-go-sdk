// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorAlertHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonitorAlertHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonitorAlertHistoryResponse
	GetStatusCode() *int32
	SetBody(v *MonitorAlertHistoryPageResult) *GetMonitorAlertHistoryResponse
	GetBody() *MonitorAlertHistoryPageResult
}

type GetMonitorAlertHistoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorAlertHistoryPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorAlertHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorAlertHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorAlertHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonitorAlertHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonitorAlertHistoryResponse) GetBody() *MonitorAlertHistoryPageResult {
	return s.Body
}

func (s *GetMonitorAlertHistoryResponse) SetHeaders(v map[string]*string) *GetMonitorAlertHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorAlertHistoryResponse) SetStatusCode(v int32) *GetMonitorAlertHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonitorAlertHistoryResponse) SetBody(v *MonitorAlertHistoryPageResult) *GetMonitorAlertHistoryResponse {
	s.Body = v
	return s
}

func (s *GetMonitorAlertHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
