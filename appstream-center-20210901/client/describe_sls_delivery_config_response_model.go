// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsDeliveryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsDeliveryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsDeliveryConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsDeliveryConfigResponseBody) *DescribeSlsDeliveryConfigResponse
	GetBody() *DescribeSlsDeliveryConfigResponseBody
}

type DescribeSlsDeliveryConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsDeliveryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsDeliveryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsDeliveryConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsDeliveryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsDeliveryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsDeliveryConfigResponse) GetBody() *DescribeSlsDeliveryConfigResponseBody {
	return s.Body
}

func (s *DescribeSlsDeliveryConfigResponse) SetHeaders(v map[string]*string) *DescribeSlsDeliveryConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsDeliveryConfigResponse) SetStatusCode(v int32) *DescribeSlsDeliveryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponse) SetBody(v *DescribeSlsDeliveryConfigResponseBody) *DescribeSlsDeliveryConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsDeliveryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
