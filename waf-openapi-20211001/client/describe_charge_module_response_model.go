// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChargeModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChargeModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChargeModuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChargeModuleResponseBody) *DescribeChargeModuleResponse
	GetBody() *DescribeChargeModuleResponseBody
}

type DescribeChargeModuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChargeModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChargeModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeModuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeChargeModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChargeModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChargeModuleResponse) GetBody() *DescribeChargeModuleResponseBody {
	return s.Body
}

func (s *DescribeChargeModuleResponse) SetHeaders(v map[string]*string) *DescribeChargeModuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeChargeModuleResponse) SetStatusCode(v int32) *DescribeChargeModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChargeModuleResponse) SetBody(v *DescribeChargeModuleResponseBody) *DescribeChargeModuleResponse {
	s.Body = v
	return s
}

func (s *DescribeChargeModuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
