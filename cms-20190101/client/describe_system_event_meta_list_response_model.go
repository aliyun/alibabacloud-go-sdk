// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventMetaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemEventMetaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemEventMetaListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemEventMetaListResponseBody) *DescribeSystemEventMetaListResponse
	GetBody() *DescribeSystemEventMetaListResponseBody
}

type DescribeSystemEventMetaListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemEventMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemEventMetaListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventMetaListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventMetaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemEventMetaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemEventMetaListResponse) GetBody() *DescribeSystemEventMetaListResponseBody {
	return s.Body
}

func (s *DescribeSystemEventMetaListResponse) SetHeaders(v map[string]*string) *DescribeSystemEventMetaListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemEventMetaListResponse) SetStatusCode(v int32) *DescribeSystemEventMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemEventMetaListResponse) SetBody(v *DescribeSystemEventMetaListResponseBody) *DescribeSystemEventMetaListResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemEventMetaListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
