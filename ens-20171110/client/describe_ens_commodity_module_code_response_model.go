// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsCommodityModuleCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsCommodityModuleCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsCommodityModuleCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsCommodityModuleCodeResponseBody) *DescribeEnsCommodityModuleCodeResponse
	GetBody() *DescribeEnsCommodityModuleCodeResponseBody
}

type DescribeEnsCommodityModuleCodeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsCommodityModuleCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsCommodityModuleCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityModuleCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityModuleCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsCommodityModuleCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsCommodityModuleCodeResponse) GetBody() *DescribeEnsCommodityModuleCodeResponseBody {
	return s.Body
}

func (s *DescribeEnsCommodityModuleCodeResponse) SetHeaders(v map[string]*string) *DescribeEnsCommodityModuleCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponse) SetStatusCode(v int32) *DescribeEnsCommodityModuleCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponse) SetBody(v *DescribeEnsCommodityModuleCodeResponseBody) *DescribeEnsCommodityModuleCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
