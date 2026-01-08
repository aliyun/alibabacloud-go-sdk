// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommerceSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCommerceSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCommerceSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCommerceSettingResponseBody) *UpdateCommerceSettingResponse
	GetBody() *UpdateCommerceSettingResponseBody
}

type UpdateCommerceSettingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCommerceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCommerceSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommerceSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommerceSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCommerceSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCommerceSettingResponse) GetBody() *UpdateCommerceSettingResponseBody {
	return s.Body
}

func (s *UpdateCommerceSettingResponse) SetHeaders(v map[string]*string) *UpdateCommerceSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateCommerceSettingResponse) SetStatusCode(v int32) *UpdateCommerceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCommerceSettingResponse) SetBody(v *UpdateCommerceSettingResponseBody) *UpdateCommerceSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateCommerceSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
