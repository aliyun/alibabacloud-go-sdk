// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonOverallConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommonOverallConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommonOverallConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommonOverallConfigResponseBody) *DescribeCommonOverallConfigResponse
	GetBody() *DescribeCommonOverallConfigResponseBody
}

type DescribeCommonOverallConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommonOverallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommonOverallConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonOverallConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommonOverallConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommonOverallConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommonOverallConfigResponse) GetBody() *DescribeCommonOverallConfigResponseBody {
	return s.Body
}

func (s *DescribeCommonOverallConfigResponse) SetHeaders(v map[string]*string) *DescribeCommonOverallConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommonOverallConfigResponse) SetStatusCode(v int32) *DescribeCommonOverallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommonOverallConfigResponse) SetBody(v *DescribeCommonOverallConfigResponseBody) *DescribeCommonOverallConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCommonOverallConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
