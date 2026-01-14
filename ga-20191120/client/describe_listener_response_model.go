// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeListenerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeListenerResponseBody) *DescribeListenerResponse
	GetBody() *DescribeListenerResponseBody
}

type DescribeListenerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponse) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeListenerResponse) GetBody() *DescribeListenerResponseBody {
	return s.Body
}

func (s *DescribeListenerResponse) SetHeaders(v map[string]*string) *DescribeListenerResponse {
	s.Headers = v
	return s
}

func (s *DescribeListenerResponse) SetStatusCode(v int32) *DescribeListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListenerResponse) SetBody(v *DescribeListenerResponseBody) *DescribeListenerResponse {
	s.Body = v
	return s
}

func (s *DescribeListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
