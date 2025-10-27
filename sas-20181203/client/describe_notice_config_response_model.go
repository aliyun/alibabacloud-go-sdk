// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNoticeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNoticeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNoticeConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNoticeConfigResponseBody) *DescribeNoticeConfigResponse
	GetBody() *DescribeNoticeConfigResponseBody
}

type DescribeNoticeConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNoticeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNoticeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNoticeConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeNoticeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNoticeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNoticeConfigResponse) GetBody() *DescribeNoticeConfigResponseBody {
	return s.Body
}

func (s *DescribeNoticeConfigResponse) SetHeaders(v map[string]*string) *DescribeNoticeConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeNoticeConfigResponse) SetStatusCode(v int32) *DescribeNoticeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNoticeConfigResponse) SetBody(v *DescribeNoticeConfigResponseBody) *DescribeNoticeConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeNoticeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
