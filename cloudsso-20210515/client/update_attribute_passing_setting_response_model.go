// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttributePassingSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAttributePassingSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAttributePassingSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAttributePassingSettingResponseBody) *UpdateAttributePassingSettingResponse
	GetBody() *UpdateAttributePassingSettingResponseBody
}

type UpdateAttributePassingSettingResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAttributePassingSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAttributePassingSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttributePassingSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateAttributePassingSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAttributePassingSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAttributePassingSettingResponse) GetBody() *UpdateAttributePassingSettingResponseBody {
	return s.Body
}

func (s *UpdateAttributePassingSettingResponse) SetHeaders(v map[string]*string) *UpdateAttributePassingSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateAttributePassingSettingResponse) SetStatusCode(v int32) *UpdateAttributePassingSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAttributePassingSettingResponse) SetBody(v *UpdateAttributePassingSettingResponseBody) *UpdateAttributePassingSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateAttributePassingSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
