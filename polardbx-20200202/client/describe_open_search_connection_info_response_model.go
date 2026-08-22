// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchConnectionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchConnectionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchConnectionInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchConnectionInfoResponseBody) *DescribeOpenSearchConnectionInfoResponse
	GetBody() *DescribeOpenSearchConnectionInfoResponseBody
}

type DescribeOpenSearchConnectionInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchConnectionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchConnectionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchConnectionInfoResponse) GetBody() *DescribeOpenSearchConnectionInfoResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchConnectionInfoResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchConnectionInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponse) SetStatusCode(v int32) *DescribeOpenSearchConnectionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponse) SetBody(v *DescribeOpenSearchConnectionInfoResponseBody) *DescribeOpenSearchConnectionInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
