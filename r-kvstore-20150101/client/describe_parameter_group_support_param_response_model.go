// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupSupportParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParameterGroupSupportParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParameterGroupSupportParamResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParameterGroupSupportParamResponseBody) *DescribeParameterGroupSupportParamResponse
	GetBody() *DescribeParameterGroupSupportParamResponseBody
}

type DescribeParameterGroupSupportParamResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterGroupSupportParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParameterGroupSupportParamResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupSupportParamResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupSupportParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParameterGroupSupportParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParameterGroupSupportParamResponse) GetBody() *DescribeParameterGroupSupportParamResponseBody {
	return s.Body
}

func (s *DescribeParameterGroupSupportParamResponse) SetHeaders(v map[string]*string) *DescribeParameterGroupSupportParamResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterGroupSupportParamResponse) SetStatusCode(v int32) *DescribeParameterGroupSupportParamResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterGroupSupportParamResponse) SetBody(v *DescribeParameterGroupSupportParamResponseBody) *DescribeParameterGroupSupportParamResponse {
	s.Body = v
	return s
}

func (s *DescribeParameterGroupSupportParamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
