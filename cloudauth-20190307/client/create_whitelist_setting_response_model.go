// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhitelistSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWhitelistSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWhitelistSettingResponse
	GetStatusCode() *int32
	SetBody(v *CreateWhitelistSettingResponseBody) *CreateWhitelistSettingResponse
	GetBody() *CreateWhitelistSettingResponseBody
}

type CreateWhitelistSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWhitelistSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWhitelistSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistSettingResponse) GoString() string {
	return s.String()
}

func (s *CreateWhitelistSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWhitelistSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWhitelistSettingResponse) GetBody() *CreateWhitelistSettingResponseBody {
	return s.Body
}

func (s *CreateWhitelistSettingResponse) SetHeaders(v map[string]*string) *CreateWhitelistSettingResponse {
	s.Headers = v
	return s
}

func (s *CreateWhitelistSettingResponse) SetStatusCode(v int32) *CreateWhitelistSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWhitelistSettingResponse) SetBody(v *CreateWhitelistSettingResponseBody) *CreateWhitelistSettingResponse {
	s.Body = v
	return s
}

func (s *CreateWhitelistSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
