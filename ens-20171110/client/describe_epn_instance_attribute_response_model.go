// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEpnInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEpnInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEpnInstanceAttributeResponseBody) *DescribeEpnInstanceAttributeResponse
	GetBody() *DescribeEpnInstanceAttributeResponseBody
}

type DescribeEpnInstanceAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEpnInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEpnInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEpnInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEpnInstanceAttributeResponse) GetBody() *DescribeEpnInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeEpnInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeEpnInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetStatusCode(v int32) *DescribeEpnInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetBody(v *DescribeEpnInstanceAttributeResponseBody) *DescribeEpnInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
