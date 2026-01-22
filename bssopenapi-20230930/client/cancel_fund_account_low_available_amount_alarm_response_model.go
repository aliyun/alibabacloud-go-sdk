// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFundAccountLowAvailableAmountAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelFundAccountLowAvailableAmountAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelFundAccountLowAvailableAmountAlarmResponse
	GetStatusCode() *int32
	SetBody(v *CancelFundAccountLowAvailableAmountAlarmResponseBody) *CancelFundAccountLowAvailableAmountAlarmResponse
	GetBody() *CancelFundAccountLowAvailableAmountAlarmResponseBody
}

type CancelFundAccountLowAvailableAmountAlarmResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelFundAccountLowAvailableAmountAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelFundAccountLowAvailableAmountAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelFundAccountLowAvailableAmountAlarmResponse) GoString() string {
	return s.String()
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) GetBody() *CancelFundAccountLowAvailableAmountAlarmResponseBody {
	return s.Body
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) SetHeaders(v map[string]*string) *CancelFundAccountLowAvailableAmountAlarmResponse {
	s.Headers = v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) SetStatusCode(v int32) *CancelFundAccountLowAvailableAmountAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) SetBody(v *CancelFundAccountLowAvailableAmountAlarmResponseBody) *CancelFundAccountLowAvailableAmountAlarmResponse {
	s.Body = v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
