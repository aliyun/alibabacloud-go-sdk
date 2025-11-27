// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteADSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteADSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteADSettingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteADSettingResponseBody) *DeleteADSettingResponse
	GetBody() *DeleteADSettingResponseBody
}

type DeleteADSettingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteADSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteADSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteADSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteADSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteADSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteADSettingResponse) GetBody() *DeleteADSettingResponseBody {
	return s.Body
}

func (s *DeleteADSettingResponse) SetHeaders(v map[string]*string) *DeleteADSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteADSettingResponse) SetStatusCode(v int32) *DeleteADSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteADSettingResponse) SetBody(v *DeleteADSettingResponseBody) *DeleteADSettingResponse {
	s.Body = v
	return s
}

func (s *DeleteADSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
