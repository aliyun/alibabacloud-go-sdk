// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetUserSettingResponseBody) *GetUserSettingResponse
	GetBody() *GetUserSettingResponseBody
}

type GetUserSettingResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserSettingResponse) GoString() string {
	return s.String()
}

func (s *GetUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserSettingResponse) GetBody() *GetUserSettingResponseBody {
	return s.Body
}

func (s *GetUserSettingResponse) SetHeaders(v map[string]*string) *GetUserSettingResponse {
	s.Headers = v
	return s
}

func (s *GetUserSettingResponse) SetStatusCode(v int32) *GetUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserSettingResponse) SetBody(v *GetUserSettingResponseBody) *GetUserSettingResponse {
	s.Body = v
	return s
}

func (s *GetUserSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
