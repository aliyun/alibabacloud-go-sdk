// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSuspEventUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSuspEventUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSuspEventUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *SaveSuspEventUserSettingResponseBody) *SaveSuspEventUserSettingResponse
	GetBody() *SaveSuspEventUserSettingResponseBody
}

type SaveSuspEventUserSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSuspEventUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSuspEventUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSuspEventUserSettingResponse) GoString() string {
	return s.String()
}

func (s *SaveSuspEventUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSuspEventUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSuspEventUserSettingResponse) GetBody() *SaveSuspEventUserSettingResponseBody {
	return s.Body
}

func (s *SaveSuspEventUserSettingResponse) SetHeaders(v map[string]*string) *SaveSuspEventUserSettingResponse {
	s.Headers = v
	return s
}

func (s *SaveSuspEventUserSettingResponse) SetStatusCode(v int32) *SaveSuspEventUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSuspEventUserSettingResponse) SetBody(v *SaveSuspEventUserSettingResponseBody) *SaveSuspEventUserSettingResponse {
	s.Body = v
	return s
}

func (s *SaveSuspEventUserSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
