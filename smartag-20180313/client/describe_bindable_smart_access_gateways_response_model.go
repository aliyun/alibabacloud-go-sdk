// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBindableSmartAccessGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBindableSmartAccessGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBindableSmartAccessGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBindableSmartAccessGatewaysResponseBody) *DescribeBindableSmartAccessGatewaysResponse
	GetBody() *DescribeBindableSmartAccessGatewaysResponseBody
}

type DescribeBindableSmartAccessGatewaysResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBindableSmartAccessGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBindableSmartAccessGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBindableSmartAccessGatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeBindableSmartAccessGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBindableSmartAccessGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBindableSmartAccessGatewaysResponse) GetBody() *DescribeBindableSmartAccessGatewaysResponseBody {
	return s.Body
}

func (s *DescribeBindableSmartAccessGatewaysResponse) SetHeaders(v map[string]*string) *DescribeBindableSmartAccessGatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponse) SetStatusCode(v int32) *DescribeBindableSmartAccessGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponse) SetBody(v *DescribeBindableSmartAccessGatewaysResponseBody) *DescribeBindableSmartAccessGatewaysResponse {
	s.Body = v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
