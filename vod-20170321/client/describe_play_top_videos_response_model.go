// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayTopVideosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayTopVideosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayTopVideosResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayTopVideosResponseBody) *DescribePlayTopVideosResponse
	GetBody() *DescribePlayTopVideosResponseBody
}

type DescribePlayTopVideosResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayTopVideosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayTopVideosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayTopVideosResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayTopVideosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayTopVideosResponse) GetBody() *DescribePlayTopVideosResponseBody {
	return s.Body
}

func (s *DescribePlayTopVideosResponse) SetHeaders(v map[string]*string) *DescribePlayTopVideosResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayTopVideosResponse) SetStatusCode(v int32) *DescribePlayTopVideosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayTopVideosResponse) SetBody(v *DescribePlayTopVideosResponseBody) *DescribePlayTopVideosResponse {
	s.Body = v
	return s
}

func (s *DescribePlayTopVideosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
