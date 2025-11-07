// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryWhiteListSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageQueryWhiteListSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageQueryWhiteListSettingResponse
	GetStatusCode() *int32
	SetBody(v *PageQueryWhiteListSettingResponseBody) *PageQueryWhiteListSettingResponse
	GetBody() *PageQueryWhiteListSettingResponseBody
}

type PageQueryWhiteListSettingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageQueryWhiteListSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageQueryWhiteListSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s PageQueryWhiteListSettingResponse) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageQueryWhiteListSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageQueryWhiteListSettingResponse) GetBody() *PageQueryWhiteListSettingResponseBody {
	return s.Body
}

func (s *PageQueryWhiteListSettingResponse) SetHeaders(v map[string]*string) *PageQueryWhiteListSettingResponse {
	s.Headers = v
	return s
}

func (s *PageQueryWhiteListSettingResponse) SetStatusCode(v int32) *PageQueryWhiteListSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *PageQueryWhiteListSettingResponse) SetBody(v *PageQueryWhiteListSettingResponseBody) *PageQueryWhiteListSettingResponse {
	s.Body = v
	return s
}

func (s *PageQueryWhiteListSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
