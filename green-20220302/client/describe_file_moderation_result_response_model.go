// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileModerationResultResponseBody) *DescribeFileModerationResultResponse
	GetBody() *DescribeFileModerationResultResponseBody
}

type DescribeFileModerationResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileModerationResultResponse) GetBody() *DescribeFileModerationResultResponseBody {
	return s.Body
}

func (s *DescribeFileModerationResultResponse) SetHeaders(v map[string]*string) *DescribeFileModerationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileModerationResultResponse) SetStatusCode(v int32) *DescribeFileModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileModerationResultResponse) SetBody(v *DescribeFileModerationResultResponseBody) *DescribeFileModerationResultResponse {
	s.Body = v
	return s
}

func (s *DescribeFileModerationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
