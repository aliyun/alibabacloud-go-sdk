// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowAliasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowAliasResponseBody) *DescribeFlowAliasResponse
	GetBody() *DescribeFlowAliasResponseBody
}

type DescribeFlowAliasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowAliasResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowAliasResponse) GetBody() *DescribeFlowAliasResponseBody {
	return s.Body
}

func (s *DescribeFlowAliasResponse) SetHeaders(v map[string]*string) *DescribeFlowAliasResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowAliasResponse) SetStatusCode(v int32) *DescribeFlowAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowAliasResponse) SetBody(v *DescribeFlowAliasResponseBody) *DescribeFlowAliasResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
