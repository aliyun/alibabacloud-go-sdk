// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteMonitorLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteMonitorLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteMonitorLogResponseBody) *DescribeSiteMonitorLogResponse
	GetBody() *DescribeSiteMonitorLogResponseBody
}

type DescribeSiteMonitorLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteMonitorLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteMonitorLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteMonitorLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteMonitorLogResponse) GetBody() *DescribeSiteMonitorLogResponseBody {
	return s.Body
}

func (s *DescribeSiteMonitorLogResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorLogResponse) SetStatusCode(v int32) *DescribeSiteMonitorLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteMonitorLogResponse) SetBody(v *DescribeSiteMonitorLogResponseBody) *DescribeSiteMonitorLogResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteMonitorLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
