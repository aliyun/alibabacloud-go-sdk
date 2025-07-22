// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountLowAvailableAmountAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFundAccountLowAvailableAmountAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFundAccountLowAvailableAmountAlarmResponse
	GetStatusCode() *int32
	SetBody(v *GetFundAccountLowAvailableAmountAlarmResponseBody) *GetFundAccountLowAvailableAmountAlarmResponse
	GetBody() *GetFundAccountLowAvailableAmountAlarmResponseBody
}

type GetFundAccountLowAvailableAmountAlarmResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountLowAvailableAmountAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountLowAvailableAmountAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountLowAvailableAmountAlarmResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) GetBody() *GetFundAccountLowAvailableAmountAlarmResponseBody {
	return s.Body
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) SetHeaders(v map[string]*string) *GetFundAccountLowAvailableAmountAlarmResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) SetStatusCode(v int32) *GetFundAccountLowAvailableAmountAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) SetBody(v *GetFundAccountLowAvailableAmountAlarmResponseBody) *GetFundAccountLowAvailableAmountAlarmResponse {
	s.Body = v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) Validate() error {
	return dara.Validate(s)
}
