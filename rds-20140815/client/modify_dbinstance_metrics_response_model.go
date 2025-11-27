// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceMetricsResponseBody) *ModifyDBInstanceMetricsResponse
	GetBody() *ModifyDBInstanceMetricsResponseBody
}

type ModifyDBInstanceMetricsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceMetricsResponse) GetBody() *ModifyDBInstanceMetricsResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceMetricsResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMetricsResponse) SetStatusCode(v int32) *ModifyDBInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceMetricsResponse) SetBody(v *ModifyDBInstanceMetricsResponseBody) *ModifyDBInstanceMetricsResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
