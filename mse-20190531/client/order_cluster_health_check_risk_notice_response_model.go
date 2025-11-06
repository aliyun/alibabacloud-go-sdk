// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderClusterHealthCheckRiskNoticeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OrderClusterHealthCheckRiskNoticeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OrderClusterHealthCheckRiskNoticeResponse
	GetStatusCode() *int32
	SetBody(v *OrderClusterHealthCheckRiskNoticeResponseBody) *OrderClusterHealthCheckRiskNoticeResponse
	GetBody() *OrderClusterHealthCheckRiskNoticeResponseBody
}

type OrderClusterHealthCheckRiskNoticeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderClusterHealthCheckRiskNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OrderClusterHealthCheckRiskNoticeResponse) String() string {
	return dara.Prettify(s)
}

func (s OrderClusterHealthCheckRiskNoticeResponse) GoString() string {
	return s.String()
}

func (s *OrderClusterHealthCheckRiskNoticeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OrderClusterHealthCheckRiskNoticeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OrderClusterHealthCheckRiskNoticeResponse) GetBody() *OrderClusterHealthCheckRiskNoticeResponseBody {
	return s.Body
}

func (s *OrderClusterHealthCheckRiskNoticeResponse) SetHeaders(v map[string]*string) *OrderClusterHealthCheckRiskNoticeResponse {
	s.Headers = v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponse) SetStatusCode(v int32) *OrderClusterHealthCheckRiskNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponse) SetBody(v *OrderClusterHealthCheckRiskNoticeResponseBody) *OrderClusterHealthCheckRiskNoticeResponse {
	s.Body = v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
