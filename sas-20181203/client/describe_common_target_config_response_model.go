// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonTargetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommonTargetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommonTargetConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommonTargetConfigResponseBody) *DescribeCommonTargetConfigResponse
	GetBody() *DescribeCommonTargetConfigResponseBody
}

type DescribeCommonTargetConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommonTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommonTargetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommonTargetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommonTargetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommonTargetConfigResponse) GetBody() *DescribeCommonTargetConfigResponseBody {
	return s.Body
}

func (s *DescribeCommonTargetConfigResponse) SetHeaders(v map[string]*string) *DescribeCommonTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommonTargetConfigResponse) SetStatusCode(v int32) *DescribeCommonTargetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommonTargetConfigResponse) SetBody(v *DescribeCommonTargetConfigResponseBody) *DescribeCommonTargetConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCommonTargetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
