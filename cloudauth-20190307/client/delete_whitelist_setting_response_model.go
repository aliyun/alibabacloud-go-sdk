// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhitelistSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWhitelistSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWhitelistSettingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWhitelistSettingResponseBody) *DeleteWhitelistSettingResponse
	GetBody() *DeleteWhitelistSettingResponseBody
}

type DeleteWhitelistSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWhitelistSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWhitelistSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWhitelistSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWhitelistSettingResponse) GetBody() *DeleteWhitelistSettingResponseBody {
	return s.Body
}

func (s *DeleteWhitelistSettingResponse) SetHeaders(v map[string]*string) *DeleteWhitelistSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhitelistSettingResponse) SetStatusCode(v int32) *DeleteWhitelistSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWhitelistSettingResponse) SetBody(v *DeleteWhitelistSettingResponseBody) *DeleteWhitelistSettingResponse {
	s.Body = v
	return s
}

func (s *DeleteWhitelistSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
