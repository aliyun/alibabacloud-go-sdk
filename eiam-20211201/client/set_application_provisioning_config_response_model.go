// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApplicationProvisioningConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApplicationProvisioningConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetApplicationProvisioningConfigResponseBody) *SetApplicationProvisioningConfigResponse
	GetBody() *SetApplicationProvisioningConfigResponseBody
}

type SetApplicationProvisioningConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationProvisioningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationProvisioningConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningConfigResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApplicationProvisioningConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApplicationProvisioningConfigResponse) GetBody() *SetApplicationProvisioningConfigResponseBody {
	return s.Body
}

func (s *SetApplicationProvisioningConfigResponse) SetHeaders(v map[string]*string) *SetApplicationProvisioningConfigResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationProvisioningConfigResponse) SetStatusCode(v int32) *SetApplicationProvisioningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationProvisioningConfigResponse) SetBody(v *SetApplicationProvisioningConfigResponseBody) *SetApplicationProvisioningConfigResponse {
	s.Body = v
	return s
}

func (s *SetApplicationProvisioningConfigResponse) Validate() error {
	return dara.Validate(s)
}
