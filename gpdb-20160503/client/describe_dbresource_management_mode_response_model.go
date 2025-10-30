// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceManagementModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBResourceManagementModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBResourceManagementModeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBResourceManagementModeResponseBody) *DescribeDBResourceManagementModeResponse
	GetBody() *DescribeDBResourceManagementModeResponseBody
}

type DescribeDBResourceManagementModeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBResourceManagementModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBResourceManagementModeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceManagementModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceManagementModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBResourceManagementModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBResourceManagementModeResponse) GetBody() *DescribeDBResourceManagementModeResponseBody {
	return s.Body
}

func (s *DescribeDBResourceManagementModeResponse) SetHeaders(v map[string]*string) *DescribeDBResourceManagementModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBResourceManagementModeResponse) SetStatusCode(v int32) *DescribeDBResourceManagementModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBResourceManagementModeResponse) SetBody(v *DescribeDBResourceManagementModeResponseBody) *DescribeDBResourceManagementModeResponse {
	s.Body = v
	return s
}

func (s *DescribeDBResourceManagementModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
