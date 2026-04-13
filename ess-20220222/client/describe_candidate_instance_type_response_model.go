// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCandidateInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCandidateInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCandidateInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCandidateInstanceTypeResponseBody) *DescribeCandidateInstanceTypeResponse
	GetBody() *DescribeCandidateInstanceTypeResponseBody
}

type DescribeCandidateInstanceTypeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCandidateInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCandidateInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCandidateInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCandidateInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCandidateInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCandidateInstanceTypeResponse) GetBody() *DescribeCandidateInstanceTypeResponseBody {
	return s.Body
}

func (s *DescribeCandidateInstanceTypeResponse) SetHeaders(v map[string]*string) *DescribeCandidateInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCandidateInstanceTypeResponse) SetStatusCode(v int32) *DescribeCandidateInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCandidateInstanceTypeResponse) SetBody(v *DescribeCandidateInstanceTypeResponseBody) *DescribeCandidateInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeCandidateInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
