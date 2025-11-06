// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *CancelUserSettingResponseBody) *CancelUserSettingResponse
	GetBody() *CancelUserSettingResponseBody
}

type CancelUserSettingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelUserSettingResponse) GoString() string {
	return s.String()
}

func (s *CancelUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelUserSettingResponse) GetBody() *CancelUserSettingResponseBody {
	return s.Body
}

func (s *CancelUserSettingResponse) SetHeaders(v map[string]*string) *CancelUserSettingResponse {
	s.Headers = v
	return s
}

func (s *CancelUserSettingResponse) SetStatusCode(v int32) *CancelUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelUserSettingResponse) SetBody(v *CancelUserSettingResponseBody) *CancelUserSettingResponse {
	s.Body = v
	return s
}

func (s *CancelUserSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
