// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfsTenantQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPxfsTenantQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPxfsTenantQuotaResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPxfsTenantQuotaResponseBody) *ModifyPxfsTenantQuotaResponse
	GetBody() *ModifyPxfsTenantQuotaResponseBody
}

type ModifyPxfsTenantQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPxfsTenantQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPxfsTenantQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsTenantQuotaResponse) GoString() string {
	return s.String()
}

func (s *ModifyPxfsTenantQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPxfsTenantQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPxfsTenantQuotaResponse) GetBody() *ModifyPxfsTenantQuotaResponseBody {
	return s.Body
}

func (s *ModifyPxfsTenantQuotaResponse) SetHeaders(v map[string]*string) *ModifyPxfsTenantQuotaResponse {
	s.Headers = v
	return s
}

func (s *ModifyPxfsTenantQuotaResponse) SetStatusCode(v int32) *ModifyPxfsTenantQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponse) SetBody(v *ModifyPxfsTenantQuotaResponseBody) *ModifyPxfsTenantQuotaResponse {
	s.Body = v
	return s
}

func (s *ModifyPxfsTenantQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
