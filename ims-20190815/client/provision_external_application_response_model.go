// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionExternalApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProvisionExternalApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProvisionExternalApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ProvisionExternalApplicationResponseBody) *ProvisionExternalApplicationResponse
	GetBody() *ProvisionExternalApplicationResponseBody
}

type ProvisionExternalApplicationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionExternalApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProvisionExternalApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ProvisionExternalApplicationResponse) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProvisionExternalApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProvisionExternalApplicationResponse) GetBody() *ProvisionExternalApplicationResponseBody {
	return s.Body
}

func (s *ProvisionExternalApplicationResponse) SetHeaders(v map[string]*string) *ProvisionExternalApplicationResponse {
	s.Headers = v
	return s
}

func (s *ProvisionExternalApplicationResponse) SetStatusCode(v int32) *ProvisionExternalApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ProvisionExternalApplicationResponse) SetBody(v *ProvisionExternalApplicationResponseBody) *ProvisionExternalApplicationResponse {
	s.Body = v
	return s
}

func (s *ProvisionExternalApplicationResponse) Validate() error {
	return dara.Validate(s)
}
