// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListenerAccessControlAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeListenerAccessControlAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeListenerAccessControlAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeListenerAccessControlAttributeResponseBody) *DescribeListenerAccessControlAttributeResponse
	GetBody() *DescribeListenerAccessControlAttributeResponseBody
}

type DescribeListenerAccessControlAttributeResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeListenerAccessControlAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeListenerAccessControlAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerAccessControlAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeListenerAccessControlAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeListenerAccessControlAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeListenerAccessControlAttributeResponse) GetBody() *DescribeListenerAccessControlAttributeResponseBody {
	return s.Body
}

func (s *DescribeListenerAccessControlAttributeResponse) SetHeaders(v map[string]*string) *DescribeListenerAccessControlAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponse) SetStatusCode(v int32) *DescribeListenerAccessControlAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponse) SetBody(v *DescribeListenerAccessControlAttributeResponseBody) *DescribeListenerAccessControlAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
