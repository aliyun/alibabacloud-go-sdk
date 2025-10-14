// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupNotifyPolicyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorGroupNotifyPolicyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorGroupNotifyPolicyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorGroupNotifyPolicyListResponseBody) *DescribeMonitorGroupNotifyPolicyListResponse
	GetBody() *DescribeMonitorGroupNotifyPolicyListResponseBody
}

type DescribeMonitorGroupNotifyPolicyListResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorGroupNotifyPolicyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) GetBody() *DescribeMonitorGroupNotifyPolicyListResponseBody {
	return s.Body
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupNotifyPolicyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) SetStatusCode(v int32) *DescribeMonitorGroupNotifyPolicyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) SetBody(v *DescribeMonitorGroupNotifyPolicyListResponseBody) *DescribeMonitorGroupNotifyPolicyListResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
