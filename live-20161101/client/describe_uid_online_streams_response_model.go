// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUidOnlineStreamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUidOnlineStreamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUidOnlineStreamsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUidOnlineStreamsResponseBody) *DescribeUidOnlineStreamsResponse
	GetBody() *DescribeUidOnlineStreamsResponseBody
}

type DescribeUidOnlineStreamsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUidOnlineStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUidOnlineStreamsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUidOnlineStreamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUidOnlineStreamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUidOnlineStreamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUidOnlineStreamsResponse) GetBody() *DescribeUidOnlineStreamsResponseBody {
	return s.Body
}

func (s *DescribeUidOnlineStreamsResponse) SetHeaders(v map[string]*string) *DescribeUidOnlineStreamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUidOnlineStreamsResponse) SetStatusCode(v int32) *DescribeUidOnlineStreamsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUidOnlineStreamsResponse) SetBody(v *DescribeUidOnlineStreamsResponseBody) *DescribeUidOnlineStreamsResponse {
	s.Body = v
	return s
}

func (s *DescribeUidOnlineStreamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
