// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommerceSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCommerceSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCommerceSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetCommerceSettingResponseBody) *GetCommerceSettingResponse
	GetBody() *GetCommerceSettingResponseBody
}

type GetCommerceSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommerceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommerceSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCommerceSettingResponse) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCommerceSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCommerceSettingResponse) GetBody() *GetCommerceSettingResponseBody {
	return s.Body
}

func (s *GetCommerceSettingResponse) SetHeaders(v map[string]*string) *GetCommerceSettingResponse {
	s.Headers = v
	return s
}

func (s *GetCommerceSettingResponse) SetStatusCode(v int32) *GetCommerceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommerceSettingResponse) SetBody(v *GetCommerceSettingResponseBody) *GetCommerceSettingResponse {
	s.Body = v
	return s
}

func (s *GetCommerceSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
