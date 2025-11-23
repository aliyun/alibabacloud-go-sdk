// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyDefaultVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDifyDefaultVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDifyDefaultVpcResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDifyDefaultVpcResponseBody) *DescribeDifyDefaultVpcResponse
	GetBody() *DescribeDifyDefaultVpcResponseBody
}

type DescribeDifyDefaultVpcResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDifyDefaultVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDifyDefaultVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyDefaultVpcResponse) GoString() string {
	return s.String()
}

func (s *DescribeDifyDefaultVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDifyDefaultVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDifyDefaultVpcResponse) GetBody() *DescribeDifyDefaultVpcResponseBody {
	return s.Body
}

func (s *DescribeDifyDefaultVpcResponse) SetHeaders(v map[string]*string) *DescribeDifyDefaultVpcResponse {
	s.Headers = v
	return s
}

func (s *DescribeDifyDefaultVpcResponse) SetStatusCode(v int32) *DescribeDifyDefaultVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponse) SetBody(v *DescribeDifyDefaultVpcResponseBody) *DescribeDifyDefaultVpcResponse {
	s.Body = v
	return s
}

func (s *DescribeDifyDefaultVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
