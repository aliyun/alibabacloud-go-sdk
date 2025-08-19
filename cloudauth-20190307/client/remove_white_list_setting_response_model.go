// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWhiteListSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveWhiteListSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveWhiteListSettingResponse
	GetStatusCode() *int32
	SetBody(v *RemoveWhiteListSettingResponseBody) *RemoveWhiteListSettingResponse
	GetBody() *RemoveWhiteListSettingResponseBody
}

type RemoveWhiteListSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveWhiteListSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveWhiteListSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveWhiteListSettingResponse) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveWhiteListSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveWhiteListSettingResponse) GetBody() *RemoveWhiteListSettingResponseBody {
	return s.Body
}

func (s *RemoveWhiteListSettingResponse) SetHeaders(v map[string]*string) *RemoveWhiteListSettingResponse {
	s.Headers = v
	return s
}

func (s *RemoveWhiteListSettingResponse) SetStatusCode(v int32) *RemoveWhiteListSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveWhiteListSettingResponse) SetBody(v *RemoveWhiteListSettingResponseBody) *RemoveWhiteListSettingResponse {
	s.Body = v
	return s
}

func (s *RemoveWhiteListSettingResponse) Validate() error {
	return dara.Validate(s)
}
