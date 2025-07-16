// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomDomainSampleRateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomDomainSampleRateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomDomainSampleRateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomDomainSampleRateResponseBody) *DeleteCustomDomainSampleRateResponse
	GetBody() *DeleteCustomDomainSampleRateResponseBody
}

type DeleteCustomDomainSampleRateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomDomainSampleRateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomDomainSampleRateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomDomainSampleRateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainSampleRateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomDomainSampleRateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomDomainSampleRateResponse) GetBody() *DeleteCustomDomainSampleRateResponseBody {
	return s.Body
}

func (s *DeleteCustomDomainSampleRateResponse) SetHeaders(v map[string]*string) *DeleteCustomDomainSampleRateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomDomainSampleRateResponse) SetStatusCode(v int32) *DeleteCustomDomainSampleRateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomDomainSampleRateResponse) SetBody(v *DeleteCustomDomainSampleRateResponseBody) *DeleteCustomDomainSampleRateResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomDomainSampleRateResponse) Validate() error {
	return dara.Validate(s)
}
