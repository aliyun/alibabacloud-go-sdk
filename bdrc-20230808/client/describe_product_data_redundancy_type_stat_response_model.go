// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductDataRedundancyTypeStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductDataRedundancyTypeStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductDataRedundancyTypeStatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductDataRedundancyTypeStatResponseBody) *DescribeProductDataRedundancyTypeStatResponse
	GetBody() *DescribeProductDataRedundancyTypeStatResponseBody
}

type DescribeProductDataRedundancyTypeStatResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductDataRedundancyTypeStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductDataRedundancyTypeStatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductDataRedundancyTypeStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductDataRedundancyTypeStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductDataRedundancyTypeStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductDataRedundancyTypeStatResponse) GetBody() *DescribeProductDataRedundancyTypeStatResponseBody {
	return s.Body
}

func (s *DescribeProductDataRedundancyTypeStatResponse) SetHeaders(v map[string]*string) *DescribeProductDataRedundancyTypeStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponse) SetStatusCode(v int32) *DescribeProductDataRedundancyTypeStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponse) SetBody(v *DescribeProductDataRedundancyTypeStatResponseBody) *DescribeProductDataRedundancyTypeStatResponse {
	s.Body = v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
