// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighDefinitionMonitorLogAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHighDefinitionMonitorLogAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHighDefinitionMonitorLogAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHighDefinitionMonitorLogAttributeResponseBody) *DescribeHighDefinitionMonitorLogAttributeResponse
	GetBody() *DescribeHighDefinitionMonitorLogAttributeResponseBody
}

type DescribeHighDefinitionMonitorLogAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHighDefinitionMonitorLogAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHighDefinitionMonitorLogAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighDefinitionMonitorLogAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponse) GetBody() *DescribeHighDefinitionMonitorLogAttributeResponseBody {
	return s.Body
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponse) SetHeaders(v map[string]*string) *DescribeHighDefinitionMonitorLogAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponse) SetStatusCode(v int32) *DescribeHighDefinitionMonitorLogAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponse) SetBody(v *DescribeHighDefinitionMonitorLogAttributeResponseBody) *DescribeHighDefinitionMonitorLogAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
