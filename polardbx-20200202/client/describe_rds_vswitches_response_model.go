// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVswitchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsVswitchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsVswitchesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsVswitchesResponseBody) *DescribeRdsVswitchesResponse
	GetBody() *DescribeRdsVswitchesResponseBody
}

type DescribeRdsVswitchesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsVswitchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsVswitchesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVswitchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVswitchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsVswitchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsVswitchesResponse) GetBody() *DescribeRdsVswitchesResponseBody {
	return s.Body
}

func (s *DescribeRdsVswitchesResponse) SetHeaders(v map[string]*string) *DescribeRdsVswitchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVswitchesResponse) SetStatusCode(v int32) *DescribeRdsVswitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVswitchesResponse) SetBody(v *DescribeRdsVswitchesResponseBody) *DescribeRdsVswitchesResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsVswitchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
