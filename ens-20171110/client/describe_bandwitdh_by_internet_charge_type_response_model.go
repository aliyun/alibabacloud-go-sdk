// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwitdhByInternetChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBandwitdhByInternetChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBandwitdhByInternetChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBandwitdhByInternetChargeTypeResponseBody) *DescribeBandwitdhByInternetChargeTypeResponse
	GetBody() *DescribeBandwitdhByInternetChargeTypeResponseBody
}

type DescribeBandwitdhByInternetChargeTypeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBandwitdhByInternetChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBandwitdhByInternetChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) GetBody() *DescribeBandwitdhByInternetChargeTypeResponseBody {
	return s.Body
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetHeaders(v map[string]*string) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetStatusCode(v int32) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetBody(v *DescribeBandwitdhByInternetChargeTypeResponseBody) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
