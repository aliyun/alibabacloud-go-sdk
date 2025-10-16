// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSitesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOfficeSitesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOfficeSitesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOfficeSitesResponseBody) *DescribeOfficeSitesResponse
	GetBody() *DescribeOfficeSitesResponseBody
}

type DescribeOfficeSitesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOfficeSitesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOfficeSitesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOfficeSitesResponse) GetBody() *DescribeOfficeSitesResponseBody {
	return s.Body
}

func (s *DescribeOfficeSitesResponse) SetHeaders(v map[string]*string) *DescribeOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOfficeSitesResponse) SetStatusCode(v int32) *DescribeOfficeSitesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOfficeSitesResponse) SetBody(v *DescribeOfficeSitesResponseBody) *DescribeOfficeSitesResponse {
	s.Body = v
	return s
}

func (s *DescribeOfficeSitesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
