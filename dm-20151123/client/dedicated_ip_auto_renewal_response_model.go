// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpAutoRenewalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpAutoRenewalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpAutoRenewalResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpAutoRenewalResponseBody) *DedicatedIpAutoRenewalResponse
	GetBody() *DedicatedIpAutoRenewalResponseBody
}

type DedicatedIpAutoRenewalResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpAutoRenewalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpAutoRenewalResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpAutoRenewalResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpAutoRenewalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpAutoRenewalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpAutoRenewalResponse) GetBody() *DedicatedIpAutoRenewalResponseBody {
	return s.Body
}

func (s *DedicatedIpAutoRenewalResponse) SetHeaders(v map[string]*string) *DedicatedIpAutoRenewalResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpAutoRenewalResponse) SetStatusCode(v int32) *DedicatedIpAutoRenewalResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpAutoRenewalResponse) SetBody(v *DedicatedIpAutoRenewalResponseBody) *DedicatedIpAutoRenewalResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpAutoRenewalResponse) Validate() error {
	return dara.Validate(s)
}
