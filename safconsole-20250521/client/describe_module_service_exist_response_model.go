// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleServiceExistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModuleServiceExistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModuleServiceExistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModuleServiceExistResponseBody) *DescribeModuleServiceExistResponse
	GetBody() *DescribeModuleServiceExistResponseBody
}

type DescribeModuleServiceExistResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModuleServiceExistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModuleServiceExistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleServiceExistResponse) GoString() string {
	return s.String()
}

func (s *DescribeModuleServiceExistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModuleServiceExistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModuleServiceExistResponse) GetBody() *DescribeModuleServiceExistResponseBody {
	return s.Body
}

func (s *DescribeModuleServiceExistResponse) SetHeaders(v map[string]*string) *DescribeModuleServiceExistResponse {
	s.Headers = v
	return s
}

func (s *DescribeModuleServiceExistResponse) SetStatusCode(v int32) *DescribeModuleServiceExistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModuleServiceExistResponse) SetBody(v *DescribeModuleServiceExistResponseBody) *DescribeModuleServiceExistResponse {
	s.Body = v
	return s
}

func (s *DescribeModuleServiceExistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
