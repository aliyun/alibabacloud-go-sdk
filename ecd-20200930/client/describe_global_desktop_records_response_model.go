// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDesktopRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalDesktopRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalDesktopRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalDesktopRecordsResponseBody) *DescribeGlobalDesktopRecordsResponse
	GetBody() *DescribeGlobalDesktopRecordsResponseBody
}

type DescribeGlobalDesktopRecordsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDesktopRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDesktopRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalDesktopRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalDesktopRecordsResponse) GetBody() *DescribeGlobalDesktopRecordsResponseBody {
	return s.Body
}

func (s *DescribeGlobalDesktopRecordsResponse) SetHeaders(v map[string]*string) *DescribeGlobalDesktopRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponse) SetStatusCode(v int32) *DescribeGlobalDesktopRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponse) SetBody(v *DescribeGlobalDesktopRecordsResponseBody) *DescribeGlobalDesktopRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
