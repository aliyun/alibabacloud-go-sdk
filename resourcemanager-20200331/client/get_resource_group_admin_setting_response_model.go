// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupAdminSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupAdminSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupAdminSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupAdminSettingResponseBody) *GetResourceGroupAdminSettingResponse
	GetBody() *GetResourceGroupAdminSettingResponseBody
}

type GetResourceGroupAdminSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupAdminSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupAdminSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupAdminSettingResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupAdminSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupAdminSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupAdminSettingResponse) GetBody() *GetResourceGroupAdminSettingResponseBody {
	return s.Body
}

func (s *GetResourceGroupAdminSettingResponse) SetHeaders(v map[string]*string) *GetResourceGroupAdminSettingResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupAdminSettingResponse) SetStatusCode(v int32) *GetResourceGroupAdminSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupAdminSettingResponse) SetBody(v *GetResourceGroupAdminSettingResponseBody) *GetResourceGroupAdminSettingResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupAdminSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
