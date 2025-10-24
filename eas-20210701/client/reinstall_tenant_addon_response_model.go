// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstallTenantAddonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReinstallTenantAddonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReinstallTenantAddonResponse
	GetStatusCode() *int32
	SetBody(v *ReinstallTenantAddonResponseBody) *ReinstallTenantAddonResponse
	GetBody() *ReinstallTenantAddonResponseBody
}

type ReinstallTenantAddonResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReinstallTenantAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReinstallTenantAddonResponse) String() string {
	return dara.Prettify(s)
}

func (s ReinstallTenantAddonResponse) GoString() string {
	return s.String()
}

func (s *ReinstallTenantAddonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReinstallTenantAddonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReinstallTenantAddonResponse) GetBody() *ReinstallTenantAddonResponseBody {
	return s.Body
}

func (s *ReinstallTenantAddonResponse) SetHeaders(v map[string]*string) *ReinstallTenantAddonResponse {
	s.Headers = v
	return s
}

func (s *ReinstallTenantAddonResponse) SetStatusCode(v int32) *ReinstallTenantAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *ReinstallTenantAddonResponse) SetBody(v *ReinstallTenantAddonResponseBody) *ReinstallTenantAddonResponse {
	s.Body = v
	return s
}

func (s *ReinstallTenantAddonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
