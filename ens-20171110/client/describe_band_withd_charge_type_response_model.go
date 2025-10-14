// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandWithdChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBandWithdChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBandWithdChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBandWithdChargeTypeResponseBody) *DescribeBandWithdChargeTypeResponse
	GetBody() *DescribeBandWithdChargeTypeResponseBody
}

type DescribeBandWithdChargeTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBandWithdChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBandWithdChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandWithdChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandWithdChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBandWithdChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBandWithdChargeTypeResponse) GetBody() *DescribeBandWithdChargeTypeResponseBody {
	return s.Body
}

func (s *DescribeBandWithdChargeTypeResponse) SetHeaders(v map[string]*string) *DescribeBandWithdChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) SetStatusCode(v int32) *DescribeBandWithdChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) SetBody(v *DescribeBandWithdChargeTypeResponseBody) *DescribeBandWithdChargeTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
