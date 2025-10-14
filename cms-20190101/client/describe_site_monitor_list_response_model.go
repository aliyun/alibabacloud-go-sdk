// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteMonitorListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteMonitorListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteMonitorListResponseBody) *DescribeSiteMonitorListResponse
	GetBody() *DescribeSiteMonitorListResponseBody
}

type DescribeSiteMonitorListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteMonitorListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteMonitorListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteMonitorListResponse) GetBody() *DescribeSiteMonitorListResponseBody {
	return s.Body
}

func (s *DescribeSiteMonitorListResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorListResponse) SetStatusCode(v int32) *DescribeSiteMonitorListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteMonitorListResponse) SetBody(v *DescribeSiteMonitorListResponseBody) *DescribeSiteMonitorListResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteMonitorListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
