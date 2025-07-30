// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationProvisioningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationProvisioningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationProvisioningResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationProvisioningResponseBody) *DisableApplicationProvisioningResponse
	GetBody() *DisableApplicationProvisioningResponseBody
}

type DisableApplicationProvisioningResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationProvisioningResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationProvisioningResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationProvisioningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationProvisioningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationProvisioningResponse) GetBody() *DisableApplicationProvisioningResponseBody {
	return s.Body
}

func (s *DisableApplicationProvisioningResponse) SetHeaders(v map[string]*string) *DisableApplicationProvisioningResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationProvisioningResponse) SetStatusCode(v int32) *DisableApplicationProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationProvisioningResponse) SetBody(v *DisableApplicationProvisioningResponseBody) *DisableApplicationProvisioningResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationProvisioningResponse) Validate() error {
	return dara.Validate(s)
}
