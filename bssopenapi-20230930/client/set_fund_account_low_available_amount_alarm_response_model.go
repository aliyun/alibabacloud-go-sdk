// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFundAccountLowAvailableAmountAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetFundAccountLowAvailableAmountAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetFundAccountLowAvailableAmountAlarmResponse
	GetStatusCode() *int32
	SetBody(v *SetFundAccountLowAvailableAmountAlarmResponseBody) *SetFundAccountLowAvailableAmountAlarmResponse
	GetBody() *SetFundAccountLowAvailableAmountAlarmResponseBody
}

type SetFundAccountLowAvailableAmountAlarmResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFundAccountLowAvailableAmountAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFundAccountLowAvailableAmountAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s SetFundAccountLowAvailableAmountAlarmResponse) GoString() string {
	return s.String()
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) GetBody() *SetFundAccountLowAvailableAmountAlarmResponseBody {
	return s.Body
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) SetHeaders(v map[string]*string) *SetFundAccountLowAvailableAmountAlarmResponse {
	s.Headers = v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) SetStatusCode(v int32) *SetFundAccountLowAvailableAmountAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) SetBody(v *SetFundAccountLowAvailableAmountAlarmResponseBody) *SetFundAccountLowAvailableAmountAlarmResponse {
	s.Body = v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) Validate() error {
	return dara.Validate(s)
}
