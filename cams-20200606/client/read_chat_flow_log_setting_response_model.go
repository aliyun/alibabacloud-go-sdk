// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadChatFlowLogSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadChatFlowLogSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadChatFlowLogSettingResponse
	GetStatusCode() *int32
	SetBody(v *ReadChatFlowLogSettingResponseBody) *ReadChatFlowLogSettingResponse
	GetBody() *ReadChatFlowLogSettingResponseBody
}

type ReadChatFlowLogSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadChatFlowLogSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadChatFlowLogSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadChatFlowLogSettingResponse) GoString() string {
	return s.String()
}

func (s *ReadChatFlowLogSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadChatFlowLogSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadChatFlowLogSettingResponse) GetBody() *ReadChatFlowLogSettingResponseBody {
	return s.Body
}

func (s *ReadChatFlowLogSettingResponse) SetHeaders(v map[string]*string) *ReadChatFlowLogSettingResponse {
	s.Headers = v
	return s
}

func (s *ReadChatFlowLogSettingResponse) SetStatusCode(v int32) *ReadChatFlowLogSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadChatFlowLogSettingResponse) SetBody(v *ReadChatFlowLogSettingResponseBody) *ReadChatFlowLogSettingResponse {
	s.Body = v
	return s
}

func (s *ReadChatFlowLogSettingResponse) Validate() error {
	return dara.Validate(s)
}
