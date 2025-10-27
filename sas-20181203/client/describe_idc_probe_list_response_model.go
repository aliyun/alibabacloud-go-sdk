// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcProbeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIdcProbeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIdcProbeListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIdcProbeListResponseBody) *DescribeIdcProbeListResponse
	GetBody() *DescribeIdcProbeListResponseBody
}

type DescribeIdcProbeListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIdcProbeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIdcProbeListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcProbeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeIdcProbeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIdcProbeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIdcProbeListResponse) GetBody() *DescribeIdcProbeListResponseBody {
	return s.Body
}

func (s *DescribeIdcProbeListResponse) SetHeaders(v map[string]*string) *DescribeIdcProbeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeIdcProbeListResponse) SetStatusCode(v int32) *DescribeIdcProbeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIdcProbeListResponse) SetBody(v *DescribeIdcProbeListResponseBody) *DescribeIdcProbeListResponse {
	s.Body = v
	return s
}

func (s *DescribeIdcProbeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
