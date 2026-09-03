// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttributePassingSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttributePassingSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttributePassingSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetAttributePassingSettingResponseBody) *GetAttributePassingSettingResponse
	GetBody() *GetAttributePassingSettingResponseBody
}

type GetAttributePassingSettingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttributePassingSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttributePassingSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttributePassingSettingResponse) GoString() string {
	return s.String()
}

func (s *GetAttributePassingSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttributePassingSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttributePassingSettingResponse) GetBody() *GetAttributePassingSettingResponseBody {
	return s.Body
}

func (s *GetAttributePassingSettingResponse) SetHeaders(v map[string]*string) *GetAttributePassingSettingResponse {
	s.Headers = v
	return s
}

func (s *GetAttributePassingSettingResponse) SetStatusCode(v int32) *GetAttributePassingSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttributePassingSettingResponse) SetBody(v *GetAttributePassingSettingResponseBody) *GetAttributePassingSettingResponse {
	s.Body = v
	return s
}

func (s *GetAttributePassingSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
