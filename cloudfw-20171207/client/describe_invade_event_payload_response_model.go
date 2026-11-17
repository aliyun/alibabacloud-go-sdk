// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventPayloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvadeEventPayloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvadeEventPayloadResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvadeEventPayloadResponseBody) *DescribeInvadeEventPayloadResponse
	GetBody() *DescribeInvadeEventPayloadResponseBody
}

type DescribeInvadeEventPayloadResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvadeEventPayloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvadeEventPayloadResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventPayloadResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventPayloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvadeEventPayloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvadeEventPayloadResponse) GetBody() *DescribeInvadeEventPayloadResponseBody {
	return s.Body
}

func (s *DescribeInvadeEventPayloadResponse) SetHeaders(v map[string]*string) *DescribeInvadeEventPayloadResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEventPayloadResponse) SetStatusCode(v int32) *DescribeInvadeEventPayloadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponse) SetBody(v *DescribeInvadeEventPayloadResponseBody) *DescribeInvadeEventPayloadResponse {
	s.Body = v
	return s
}

func (s *DescribeInvadeEventPayloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
