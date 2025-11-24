// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeASMGatewayImportedServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeASMGatewayImportedServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeASMGatewayImportedServicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeASMGatewayImportedServicesResponseBody) *DescribeASMGatewayImportedServicesResponse
	GetBody() *DescribeASMGatewayImportedServicesResponseBody
}

type DescribeASMGatewayImportedServicesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeASMGatewayImportedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeASMGatewayImportedServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeASMGatewayImportedServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeASMGatewayImportedServicesResponse) GetBody() *DescribeASMGatewayImportedServicesResponseBody {
	return s.Body
}

func (s *DescribeASMGatewayImportedServicesResponse) SetHeaders(v map[string]*string) *DescribeASMGatewayImportedServicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponse) SetStatusCode(v int32) *DescribeASMGatewayImportedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponse) SetBody(v *DescribeASMGatewayImportedServicesResponseBody) *DescribeASMGatewayImportedServicesResponse {
	s.Body = v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
