// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatFlowLogSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChatFlowLogSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChatFlowLogSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChatFlowLogSettingResponseBody) *UpdateChatFlowLogSettingResponse
	GetBody() *UpdateChatFlowLogSettingResponseBody
}

type UpdateChatFlowLogSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChatFlowLogSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChatFlowLogSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatFlowLogSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateChatFlowLogSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChatFlowLogSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChatFlowLogSettingResponse) GetBody() *UpdateChatFlowLogSettingResponseBody {
	return s.Body
}

func (s *UpdateChatFlowLogSettingResponse) SetHeaders(v map[string]*string) *UpdateChatFlowLogSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateChatFlowLogSettingResponse) SetStatusCode(v int32) *UpdateChatFlowLogSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChatFlowLogSettingResponse) SetBody(v *UpdateChatFlowLogSettingResponseBody) *UpdateChatFlowLogSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateChatFlowLogSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
