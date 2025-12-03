// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpdateVpcInfoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpdateVpcInfoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpdateVpcInfoTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpdateVpcInfoTaskResponseBody) *DescribeUpdateVpcInfoTaskResponse
	GetBody() *DescribeUpdateVpcInfoTaskResponseBody
}

type DescribeUpdateVpcInfoTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpdateVpcInfoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpdateVpcInfoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpdateVpcInfoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpdateVpcInfoTaskResponse) GetBody() *DescribeUpdateVpcInfoTaskResponseBody {
	return s.Body
}

func (s *DescribeUpdateVpcInfoTaskResponse) SetHeaders(v map[string]*string) *DescribeUpdateVpcInfoTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponse) SetStatusCode(v int32) *DescribeUpdateVpcInfoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponse) SetBody(v *DescribeUpdateVpcInfoTaskResponseBody) *DescribeUpdateVpcInfoTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
