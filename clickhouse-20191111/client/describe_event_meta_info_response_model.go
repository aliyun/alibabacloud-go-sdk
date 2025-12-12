// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventMetaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventMetaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventMetaInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventMetaInfoResponseBody) *DescribeEventMetaInfoResponse
	GetBody() *DescribeEventMetaInfoResponseBody
}

type DescribeEventMetaInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventMetaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventMetaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventMetaInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventMetaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventMetaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventMetaInfoResponse) GetBody() *DescribeEventMetaInfoResponseBody {
	return s.Body
}

func (s *DescribeEventMetaInfoResponse) SetHeaders(v map[string]*string) *DescribeEventMetaInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventMetaInfoResponse) SetStatusCode(v int32) *DescribeEventMetaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventMetaInfoResponse) SetBody(v *DescribeEventMetaInfoResponseBody) *DescribeEventMetaInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeEventMetaInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
