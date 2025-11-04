// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeParseSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeParseSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeParseSettingResponse
	GetStatusCode() *int32
	SetBody(v *ChangeParseSettingResponseBody) *ChangeParseSettingResponse
	GetBody() *ChangeParseSettingResponseBody
}

type ChangeParseSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeParseSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeParseSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeParseSettingResponse) GoString() string {
	return s.String()
}

func (s *ChangeParseSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeParseSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeParseSettingResponse) GetBody() *ChangeParseSettingResponseBody {
	return s.Body
}

func (s *ChangeParseSettingResponse) SetHeaders(v map[string]*string) *ChangeParseSettingResponse {
	s.Headers = v
	return s
}

func (s *ChangeParseSettingResponse) SetStatusCode(v int32) *ChangeParseSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeParseSettingResponse) SetBody(v *ChangeParseSettingResponseBody) *ChangeParseSettingResponse {
	s.Body = v
	return s
}

func (s *ChangeParseSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
