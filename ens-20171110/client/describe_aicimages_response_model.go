// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAICImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAICImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAICImagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAICImagesResponseBody) *DescribeAICImagesResponse
	GetBody() *DescribeAICImagesResponseBody
}

type DescribeAICImagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAICImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAICImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAICImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAICImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAICImagesResponse) GetBody() *DescribeAICImagesResponseBody {
	return s.Body
}

func (s *DescribeAICImagesResponse) SetHeaders(v map[string]*string) *DescribeAICImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAICImagesResponse) SetStatusCode(v int32) *DescribeAICImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAICImagesResponse) SetBody(v *DescribeAICImagesResponseBody) *DescribeAICImagesResponse {
	s.Body = v
	return s
}

func (s *DescribeAICImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
