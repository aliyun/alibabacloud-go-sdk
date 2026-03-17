// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSmartAccessGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSmartAccessGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSmartAccessGatewaysResponseBody) *DescribeSmartAccessGatewaysResponse
	GetBody() *DescribeSmartAccessGatewaysResponseBody
}

type DescribeSmartAccessGatewaysResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmartAccessGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmartAccessGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSmartAccessGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSmartAccessGatewaysResponse) GetBody() *DescribeSmartAccessGatewaysResponseBody {
	return s.Body
}

func (s *DescribeSmartAccessGatewaysResponse) SetHeaders(v map[string]*string) *DescribeSmartAccessGatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartAccessGatewaysResponse) SetStatusCode(v int32) *DescribeSmartAccessGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponse) SetBody(v *DescribeSmartAccessGatewaysResponseBody) *DescribeSmartAccessGatewaysResponse {
	s.Body = v
	return s
}

func (s *DescribeSmartAccessGatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
