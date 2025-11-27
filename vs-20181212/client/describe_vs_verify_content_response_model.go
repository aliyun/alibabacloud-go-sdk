// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsVerifyContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsVerifyContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsVerifyContentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsVerifyContentResponseBody) *DescribeVsVerifyContentResponse
	GetBody() *DescribeVsVerifyContentResponseBody
}

type DescribeVsVerifyContentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsVerifyContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsVerifyContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsVerifyContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsVerifyContentResponse) GetBody() *DescribeVsVerifyContentResponseBody {
	return s.Body
}

func (s *DescribeVsVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeVsVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsVerifyContentResponse) SetStatusCode(v int32) *DescribeVsVerifyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsVerifyContentResponse) SetBody(v *DescribeVsVerifyContentResponseBody) *DescribeVsVerifyContentResponse {
	s.Body = v
	return s
}

func (s *DescribeVsVerifyContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
