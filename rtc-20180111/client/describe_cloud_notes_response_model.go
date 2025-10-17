// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudNotesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudNotesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudNotesResponseBody) *DescribeCloudNotesResponse
	GetBody() *DescribeCloudNotesResponseBody
}

type DescribeCloudNotesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudNotesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudNotesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudNotesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudNotesResponse) GetBody() *DescribeCloudNotesResponseBody {
	return s.Body
}

func (s *DescribeCloudNotesResponse) SetHeaders(v map[string]*string) *DescribeCloudNotesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudNotesResponse) SetStatusCode(v int32) *DescribeCloudNotesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudNotesResponse) SetBody(v *DescribeCloudNotesResponseBody) *DescribeCloudNotesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudNotesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
