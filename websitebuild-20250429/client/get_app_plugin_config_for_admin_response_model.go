// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginConfigForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppPluginConfigForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppPluginConfigForAdminResponse
	GetStatusCode() *int32
	SetBody(v *GetAppPluginConfigForAdminResponseBody) *GetAppPluginConfigForAdminResponse
	GetBody() *GetAppPluginConfigForAdminResponseBody
}

type GetAppPluginConfigForAdminResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppPluginConfigForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppPluginConfigForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigForAdminResponse) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppPluginConfigForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppPluginConfigForAdminResponse) GetBody() *GetAppPluginConfigForAdminResponseBody {
	return s.Body
}

func (s *GetAppPluginConfigForAdminResponse) SetHeaders(v map[string]*string) *GetAppPluginConfigForAdminResponse {
	s.Headers = v
	return s
}

func (s *GetAppPluginConfigForAdminResponse) SetStatusCode(v int32) *GetAppPluginConfigForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponse) SetBody(v *GetAppPluginConfigForAdminResponseBody) *GetAppPluginConfigForAdminResponse {
	s.Body = v
	return s
}

func (s *GetAppPluginConfigForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
