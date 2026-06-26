// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApikeyAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApikeyAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApikeyAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApikeyAttributeResponseBody) *DescribeApikeyAttributeResponse
	GetBody() *DescribeApikeyAttributeResponseBody
}

type DescribeApikeyAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApikeyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApikeyAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApikeyAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeApikeyAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApikeyAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApikeyAttributeResponse) GetBody() *DescribeApikeyAttributeResponseBody {
	return s.Body
}

func (s *DescribeApikeyAttributeResponse) SetHeaders(v map[string]*string) *DescribeApikeyAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeApikeyAttributeResponse) SetStatusCode(v int32) *DescribeApikeyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApikeyAttributeResponse) SetBody(v *DescribeApikeyAttributeResponseBody) *DescribeApikeyAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeApikeyAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
