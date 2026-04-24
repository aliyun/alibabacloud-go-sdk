// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOTokenUsageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMOTokenUsageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMOTokenUsageDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMOTokenUsageDetailResponseBody) *DescribeMOTokenUsageDetailResponse
	GetBody() *DescribeMOTokenUsageDetailResponseBody
}

type DescribeMOTokenUsageDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMOTokenUsageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMOTokenUsageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMOTokenUsageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMOTokenUsageDetailResponse) GetBody() *DescribeMOTokenUsageDetailResponseBody {
	return s.Body
}

func (s *DescribeMOTokenUsageDetailResponse) SetHeaders(v map[string]*string) *DescribeMOTokenUsageDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeMOTokenUsageDetailResponse) SetStatusCode(v int32) *DescribeMOTokenUsageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponse) SetBody(v *DescribeMOTokenUsageDetailResponseBody) *DescribeMOTokenUsageDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeMOTokenUsageDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
