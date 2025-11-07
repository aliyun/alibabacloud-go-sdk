// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVerifySettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVerifySettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVerifySettingResponse
	GetStatusCode() *int32
	SetBody(v *CreateVerifySettingResponseBody) *CreateVerifySettingResponse
	GetBody() *CreateVerifySettingResponseBody
}

type CreateVerifySettingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVerifySettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVerifySettingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVerifySettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVerifySettingResponse) GetBody() *CreateVerifySettingResponseBody {
	return s.Body
}

func (s *CreateVerifySettingResponse) SetHeaders(v map[string]*string) *CreateVerifySettingResponse {
	s.Headers = v
	return s
}

func (s *CreateVerifySettingResponse) SetStatusCode(v int32) *CreateVerifySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVerifySettingResponse) SetBody(v *CreateVerifySettingResponseBody) *CreateVerifySettingResponse {
	s.Body = v
	return s
}

func (s *CreateVerifySettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
