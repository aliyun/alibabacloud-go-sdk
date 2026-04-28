// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelApisResponseBody) *DescribeModelApisResponse
	GetBody() *DescribeModelApisResponseBody
}

type DescribeModelApisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelApisResponse) GetBody() *DescribeModelApisResponseBody {
	return s.Body
}

func (s *DescribeModelApisResponse) SetHeaders(v map[string]*string) *DescribeModelApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelApisResponse) SetStatusCode(v int32) *DescribeModelApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelApisResponse) SetBody(v *DescribeModelApisResponseBody) *DescribeModelApisResponse {
	s.Body = v
	return s
}

func (s *DescribeModelApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
