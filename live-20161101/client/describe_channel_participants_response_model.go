// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelParticipantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelParticipantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelParticipantsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelParticipantsResponseBody) *DescribeChannelParticipantsResponse
	GetBody() *DescribeChannelParticipantsResponseBody
}

type DescribeChannelParticipantsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelParticipantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelParticipantsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelParticipantsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelParticipantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelParticipantsResponse) GetBody() *DescribeChannelParticipantsResponseBody {
	return s.Body
}

func (s *DescribeChannelParticipantsResponse) SetHeaders(v map[string]*string) *DescribeChannelParticipantsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelParticipantsResponse) SetStatusCode(v int32) *DescribeChannelParticipantsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelParticipantsResponse) SetBody(v *DescribeChannelParticipantsResponseBody) *DescribeChannelParticipantsResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelParticipantsResponse) Validate() error {
	return dara.Validate(s)
}
