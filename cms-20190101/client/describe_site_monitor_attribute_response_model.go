// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteMonitorAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteMonitorAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteMonitorAttributeResponseBody) *DescribeSiteMonitorAttributeResponse
	GetBody() *DescribeSiteMonitorAttributeResponseBody
}

type DescribeSiteMonitorAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteMonitorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteMonitorAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteMonitorAttributeResponse) GetBody() *DescribeSiteMonitorAttributeResponseBody {
	return s.Body
}

func (s *DescribeSiteMonitorAttributeResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponse) SetStatusCode(v int32) *DescribeSiteMonitorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponse) SetBody(v *DescribeSiteMonitorAttributeResponseBody) *DescribeSiteMonitorAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponse) Validate() error {
	return dara.Validate(s)
}
