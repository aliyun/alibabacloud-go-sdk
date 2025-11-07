// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaceGuardRiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaceGuardRiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaceGuardRiskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaceGuardRiskResponseBody) *DescribeFaceGuardRiskResponse
	GetBody() *DescribeFaceGuardRiskResponseBody
}

type DescribeFaceGuardRiskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaceGuardRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaceGuardRiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceGuardRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaceGuardRiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaceGuardRiskResponse) GetBody() *DescribeFaceGuardRiskResponseBody {
	return s.Body
}

func (s *DescribeFaceGuardRiskResponse) SetHeaders(v map[string]*string) *DescribeFaceGuardRiskResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceGuardRiskResponse) SetStatusCode(v int32) *DescribeFaceGuardRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaceGuardRiskResponse) SetBody(v *DescribeFaceGuardRiskResponseBody) *DescribeFaceGuardRiskResponse {
	s.Body = v
	return s
}

func (s *DescribeFaceGuardRiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
