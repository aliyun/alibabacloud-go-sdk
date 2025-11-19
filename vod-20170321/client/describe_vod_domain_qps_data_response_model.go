// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainQpsDataResponseBody) *DescribeVodDomainQpsDataResponse
	GetBody() *DescribeVodDomainQpsDataResponseBody
}

type DescribeVodDomainQpsDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainQpsDataResponse) GetBody() *DescribeVodDomainQpsDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainQpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainQpsDataResponse) SetStatusCode(v int32) *DescribeVodDomainQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainQpsDataResponse) SetBody(v *DescribeVodDomainQpsDataResponseBody) *DescribeVodDomainQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainQpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
