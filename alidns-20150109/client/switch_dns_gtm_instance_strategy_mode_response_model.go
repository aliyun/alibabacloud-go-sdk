// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDnsGtmInstanceStrategyModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchDnsGtmInstanceStrategyModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchDnsGtmInstanceStrategyModeResponse
	GetStatusCode() *int32
	SetBody(v *SwitchDnsGtmInstanceStrategyModeResponseBody) *SwitchDnsGtmInstanceStrategyModeResponse
	GetBody() *SwitchDnsGtmInstanceStrategyModeResponseBody
}

type SwitchDnsGtmInstanceStrategyModeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchDnsGtmInstanceStrategyModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchDnsGtmInstanceStrategyModeResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchDnsGtmInstanceStrategyModeResponse) GoString() string {
	return s.String()
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) GetBody() *SwitchDnsGtmInstanceStrategyModeResponseBody {
	return s.Body
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) SetHeaders(v map[string]*string) *SwitchDnsGtmInstanceStrategyModeResponse {
	s.Headers = v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) SetStatusCode(v int32) *SwitchDnsGtmInstanceStrategyModeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) SetBody(v *SwitchDnsGtmInstanceStrategyModeResponseBody) *SwitchDnsGtmInstanceStrategyModeResponse {
	s.Body = v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
