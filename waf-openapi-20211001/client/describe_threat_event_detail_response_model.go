// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeThreatEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeThreatEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeThreatEventDetailResponseBody) *DescribeThreatEventDetailResponse
	GetBody() *DescribeThreatEventDetailResponseBody
}

type DescribeThreatEventDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeThreatEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeThreatEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeThreatEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeThreatEventDetailResponse) GetBody() *DescribeThreatEventDetailResponseBody {
	return s.Body
}

func (s *DescribeThreatEventDetailResponse) SetHeaders(v map[string]*string) *DescribeThreatEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeThreatEventDetailResponse) SetStatusCode(v int32) *DescribeThreatEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeThreatEventDetailResponse) SetBody(v *DescribeThreatEventDetailResponseBody) *DescribeThreatEventDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeThreatEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
