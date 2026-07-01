// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreditDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreditDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreditDetailResponseBody) *DescribeCreditDetailResponse
	GetBody() *DescribeCreditDetailResponseBody
}

type DescribeCreditDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreditDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreditDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreditDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreditDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreditDetailResponse) GetBody() *DescribeCreditDetailResponseBody {
	return s.Body
}

func (s *DescribeCreditDetailResponse) SetHeaders(v map[string]*string) *DescribeCreditDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreditDetailResponse) SetStatusCode(v int32) *DescribeCreditDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreditDetailResponse) SetBody(v *DescribeCreditDetailResponseBody) *DescribeCreditDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCreditDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
