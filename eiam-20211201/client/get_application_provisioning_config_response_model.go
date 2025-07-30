// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationProvisioningConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationProvisioningConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationProvisioningConfigResponseBody) *GetApplicationProvisioningConfigResponse
	GetBody() *GetApplicationProvisioningConfigResponseBody
}

type GetApplicationProvisioningConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisioningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisioningConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationProvisioningConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationProvisioningConfigResponse) GetBody() *GetApplicationProvisioningConfigResponseBody {
	return s.Body
}

func (s *GetApplicationProvisioningConfigResponse) SetHeaders(v map[string]*string) *GetApplicationProvisioningConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisioningConfigResponse) SetStatusCode(v int32) *GetApplicationProvisioningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponse) SetBody(v *GetApplicationProvisioningConfigResponseBody) *GetApplicationProvisioningConfigResponse {
	s.Body = v
	return s
}

func (s *GetApplicationProvisioningConfigResponse) Validate() error {
	return dara.Validate(s)
}
