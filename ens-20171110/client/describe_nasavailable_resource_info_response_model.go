// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNASAvailableResourceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNASAvailableResourceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNASAvailableResourceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNASAvailableResourceInfoResponseBody) *DescribeNASAvailableResourceInfoResponse
	GetBody() *DescribeNASAvailableResourceInfoResponseBody
}

type DescribeNASAvailableResourceInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNASAvailableResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNASAvailableResourceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASAvailableResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeNASAvailableResourceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNASAvailableResourceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNASAvailableResourceInfoResponse) GetBody() *DescribeNASAvailableResourceInfoResponseBody {
	return s.Body
}

func (s *DescribeNASAvailableResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeNASAvailableResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponse) SetStatusCode(v int32) *DescribeNASAvailableResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponse) SetBody(v *DescribeNASAvailableResourceInfoResponseBody) *DescribeNASAvailableResourceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
