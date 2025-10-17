// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcUserCntDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcUserCntDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcUserCntDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcUserCntDataResponseBody) *DescribeRtcUserCntDataResponse
	GetBody() *DescribeRtcUserCntDataResponseBody
}

type DescribeRtcUserCntDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcUserCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcUserCntDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcUserCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcUserCntDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcUserCntDataResponse) GetBody() *DescribeRtcUserCntDataResponseBody {
	return s.Body
}

func (s *DescribeRtcUserCntDataResponse) SetHeaders(v map[string]*string) *DescribeRtcUserCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcUserCntDataResponse) SetStatusCode(v int32) *DescribeRtcUserCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcUserCntDataResponse) SetBody(v *DescribeRtcUserCntDataResponseBody) *DescribeRtcUserCntDataResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcUserCntDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
