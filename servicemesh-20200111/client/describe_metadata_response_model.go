// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetadataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetadataResponseBody) *DescribeMetadataResponse
	GetBody() *DescribeMetadataResponseBody
}

type DescribeMetadataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetadataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetadataResponse) GetBody() *DescribeMetadataResponseBody {
	return s.Body
}

func (s *DescribeMetadataResponse) SetHeaders(v map[string]*string) *DescribeMetadataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetadataResponse) SetStatusCode(v int32) *DescribeMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetadataResponse) SetBody(v *DescribeMetadataResponseBody) *DescribeMetadataResponse {
	s.Body = v
	return s
}

func (s *DescribeMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
