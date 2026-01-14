// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAcceleratorAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAcceleratorAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAcceleratorAutoRenewAttributeResponseBody) *DescribeAcceleratorAutoRenewAttributeResponse
	GetBody() *DescribeAcceleratorAutoRenewAttributeResponseBody
}

type DescribeAcceleratorAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAcceleratorAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAcceleratorAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) GetBody() *DescribeAcceleratorAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeAcceleratorAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeAcceleratorAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) SetBody(v *DescribeAcceleratorAutoRenewAttributeResponseBody) *DescribeAcceleratorAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
