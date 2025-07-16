// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomDomainSampleRateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomDomainSampleRateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomDomainSampleRateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomDomainSampleRateResponseBody) *ModifyCustomDomainSampleRateResponse
	GetBody() *ModifyCustomDomainSampleRateResponseBody
}

type ModifyCustomDomainSampleRateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomDomainSampleRateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomDomainSampleRateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomDomainSampleRateResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomDomainSampleRateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomDomainSampleRateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomDomainSampleRateResponse) GetBody() *ModifyCustomDomainSampleRateResponseBody {
	return s.Body
}

func (s *ModifyCustomDomainSampleRateResponse) SetHeaders(v map[string]*string) *ModifyCustomDomainSampleRateResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomDomainSampleRateResponse) SetStatusCode(v int32) *ModifyCustomDomainSampleRateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponse) SetBody(v *ModifyCustomDomainSampleRateResponseBody) *ModifyCustomDomainSampleRateResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomDomainSampleRateResponse) Validate() error {
	return dara.Validate(s)
}
