// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStorageGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStorageGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStorageGatewayResponseBody) *DescribeStorageGatewayResponse
	GetBody() *DescribeStorageGatewayResponseBody
}

type DescribeStorageGatewayResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStorageGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStorageGatewayResponse) GetBody() *DescribeStorageGatewayResponseBody {
	return s.Body
}

func (s *DescribeStorageGatewayResponse) SetHeaders(v map[string]*string) *DescribeStorageGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageGatewayResponse) SetStatusCode(v int32) *DescribeStorageGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageGatewayResponse) SetBody(v *DescribeStorageGatewayResponseBody) *DescribeStorageGatewayResponse {
	s.Body = v
	return s
}

func (s *DescribeStorageGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
