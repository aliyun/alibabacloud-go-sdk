// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteWafSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteWafSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteWafSettingsResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteWafSettingsResponseBody) *GetSiteWafSettingsResponse
	GetBody() *GetSiteWafSettingsResponseBody
}

type GetSiteWafSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteWafSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteWafSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteWafSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetSiteWafSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteWafSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteWafSettingsResponse) GetBody() *GetSiteWafSettingsResponseBody {
	return s.Body
}

func (s *GetSiteWafSettingsResponse) SetHeaders(v map[string]*string) *GetSiteWafSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetSiteWafSettingsResponse) SetStatusCode(v int32) *GetSiteWafSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteWafSettingsResponse) SetBody(v *GetSiteWafSettingsResponseBody) *GetSiteWafSettingsResponse {
	s.Body = v
	return s
}

func (s *GetSiteWafSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
