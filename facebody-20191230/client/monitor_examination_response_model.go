// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorExaminationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MonitorExaminationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MonitorExaminationResponse
	GetStatusCode() *int32
	SetBody(v *MonitorExaminationResponseBody) *MonitorExaminationResponse
	GetBody() *MonitorExaminationResponseBody
}

type MonitorExaminationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorExaminationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MonitorExaminationResponse) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponse) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MonitorExaminationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MonitorExaminationResponse) GetBody() *MonitorExaminationResponseBody {
	return s.Body
}

func (s *MonitorExaminationResponse) SetHeaders(v map[string]*string) *MonitorExaminationResponse {
	s.Headers = v
	return s
}

func (s *MonitorExaminationResponse) SetStatusCode(v int32) *MonitorExaminationResponse {
	s.StatusCode = &v
	return s
}

func (s *MonitorExaminationResponse) SetBody(v *MonitorExaminationResponseBody) *MonitorExaminationResponse {
	s.Body = v
	return s
}

func (s *MonitorExaminationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
