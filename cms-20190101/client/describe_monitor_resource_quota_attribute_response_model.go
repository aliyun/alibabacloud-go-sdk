// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorResourceQuotaAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorResourceQuotaAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorResourceQuotaAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorResourceQuotaAttributeResponseBody) *DescribeMonitorResourceQuotaAttributeResponse
	GetBody() *DescribeMonitorResourceQuotaAttributeResponseBody
}

type DescribeMonitorResourceQuotaAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorResourceQuotaAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) GetBody() *DescribeMonitorResourceQuotaAttributeResponseBody {
	return s.Body
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) SetHeaders(v map[string]*string) *DescribeMonitorResourceQuotaAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) SetStatusCode(v int32) *DescribeMonitorResourceQuotaAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) SetBody(v *DescribeMonitorResourceQuotaAttributeResponseBody) *DescribeMonitorResourceQuotaAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) Validate() error {
	return dara.Validate(s)
}
