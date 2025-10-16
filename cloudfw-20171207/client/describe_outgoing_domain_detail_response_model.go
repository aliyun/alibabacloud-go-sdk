// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingDomainDetailResponseBody) *DescribeOutgoingDomainDetailResponse
	GetBody() *DescribeOutgoingDomainDetailResponseBody
}

type DescribeOutgoingDomainDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingDomainDetailResponse) GetBody() *DescribeOutgoingDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeOutgoingDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDomainDetailResponse) SetStatusCode(v int32) *DescribeOutgoingDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponse) SetBody(v *DescribeOutgoingDomainDetailResponseBody) *DescribeOutgoingDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingDomainDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
