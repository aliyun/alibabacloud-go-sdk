// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessibleImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessibleImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessibleImagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessibleImagesResponseBody) *DescribeAccessibleImagesResponse
	GetBody() *DescribeAccessibleImagesResponseBody
}

type DescribeAccessibleImagesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessibleImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessibleImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessibleImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessibleImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessibleImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessibleImagesResponse) GetBody() *DescribeAccessibleImagesResponseBody {
	return s.Body
}

func (s *DescribeAccessibleImagesResponse) SetHeaders(v map[string]*string) *DescribeAccessibleImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessibleImagesResponse) SetStatusCode(v int32) *DescribeAccessibleImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessibleImagesResponse) SetBody(v *DescribeAccessibleImagesResponseBody) *DescribeAccessibleImagesResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessibleImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
