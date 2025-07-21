// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowLogSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatFlowLogSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatFlowLogSettingResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatFlowLogSettingResponseBody) *CreateChatFlowLogSettingResponse
	GetBody() *CreateChatFlowLogSettingResponseBody
}

type CreateChatFlowLogSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatFlowLogSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatFlowLogSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowLogSettingResponse) GoString() string {
	return s.String()
}

func (s *CreateChatFlowLogSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatFlowLogSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatFlowLogSettingResponse) GetBody() *CreateChatFlowLogSettingResponseBody {
	return s.Body
}

func (s *CreateChatFlowLogSettingResponse) SetHeaders(v map[string]*string) *CreateChatFlowLogSettingResponse {
	s.Headers = v
	return s
}

func (s *CreateChatFlowLogSettingResponse) SetStatusCode(v int32) *CreateChatFlowLogSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatFlowLogSettingResponse) SetBody(v *CreateChatFlowLogSettingResponseBody) *CreateChatFlowLogSettingResponse {
	s.Body = v
	return s
}

func (s *CreateChatFlowLogSettingResponse) Validate() error {
	return dara.Validate(s)
}
