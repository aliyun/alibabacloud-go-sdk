// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupAdminSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceGroupAdminSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceGroupAdminSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceGroupAdminSettingResponseBody) *UpdateResourceGroupAdminSettingResponse
	GetBody() *UpdateResourceGroupAdminSettingResponseBody
}

type UpdateResourceGroupAdminSettingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupAdminSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupAdminSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupAdminSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAdminSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceGroupAdminSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceGroupAdminSettingResponse) GetBody() *UpdateResourceGroupAdminSettingResponseBody {
	return s.Body
}

func (s *UpdateResourceGroupAdminSettingResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupAdminSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupAdminSettingResponse) SetStatusCode(v int32) *UpdateResourceGroupAdminSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupAdminSettingResponse) SetBody(v *UpdateResourceGroupAdminSettingResponseBody) *UpdateResourceGroupAdminSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceGroupAdminSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
