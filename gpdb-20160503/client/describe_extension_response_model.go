// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExtensionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExtensionResponseBody) *DescribeExtensionResponse
	GetBody() *DescribeExtensionResponseBody
}

type DescribeExtensionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionResponse) GoString() string {
	return s.String()
}

func (s *DescribeExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExtensionResponse) GetBody() *DescribeExtensionResponseBody {
	return s.Body
}

func (s *DescribeExtensionResponse) SetHeaders(v map[string]*string) *DescribeExtensionResponse {
	s.Headers = v
	return s
}

func (s *DescribeExtensionResponse) SetStatusCode(v int32) *DescribeExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExtensionResponse) SetBody(v *DescribeExtensionResponseBody) *DescribeExtensionResponse {
	s.Body = v
	return s
}

func (s *DescribeExtensionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
