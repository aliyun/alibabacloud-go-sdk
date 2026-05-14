// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDescribeClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDescribeClientResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDescribeClientResponseBody) *ClinkDescribeClientResponse
	GetBody() *ClinkDescribeClientResponseBody
}

type ClinkDescribeClientResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDescribeClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDescribeClientResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientResponse) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDescribeClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDescribeClientResponse) GetBody() *ClinkDescribeClientResponseBody {
	return s.Body
}

func (s *ClinkDescribeClientResponse) SetHeaders(v map[string]*string) *ClinkDescribeClientResponse {
	s.Headers = v
	return s
}

func (s *ClinkDescribeClientResponse) SetStatusCode(v int32) *ClinkDescribeClientResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeClientResponse) SetBody(v *ClinkDescribeClientResponseBody) *ClinkDescribeClientResponse {
	s.Body = v
	return s
}

func (s *ClinkDescribeClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
