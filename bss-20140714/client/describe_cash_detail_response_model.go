// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCashDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCashDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCashDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCashDetailResponseBody) *DescribeCashDetailResponse
	GetBody() *DescribeCashDetailResponseBody
}

type DescribeCashDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCashDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCashDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCashDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCashDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCashDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCashDetailResponse) GetBody() *DescribeCashDetailResponseBody {
	return s.Body
}

func (s *DescribeCashDetailResponse) SetHeaders(v map[string]*string) *DescribeCashDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCashDetailResponse) SetStatusCode(v int32) *DescribeCashDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCashDetailResponse) SetBody(v *DescribeCashDetailResponseBody) *DescribeCashDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCashDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
