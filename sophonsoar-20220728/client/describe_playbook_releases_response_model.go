// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookReleasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlaybookReleasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlaybookReleasesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlaybookReleasesResponseBody) *DescribePlaybookReleasesResponse
	GetBody() *DescribePlaybookReleasesResponseBody
}

type DescribePlaybookReleasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookReleasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookReleasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookReleasesResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlaybookReleasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlaybookReleasesResponse) GetBody() *DescribePlaybookReleasesResponseBody {
	return s.Body
}

func (s *DescribePlaybookReleasesResponse) SetHeaders(v map[string]*string) *DescribePlaybookReleasesResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookReleasesResponse) SetStatusCode(v int32) *DescribePlaybookReleasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookReleasesResponse) SetBody(v *DescribePlaybookReleasesResponseBody) *DescribePlaybookReleasesResponse {
	s.Body = v
	return s
}

func (s *DescribePlaybookReleasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
