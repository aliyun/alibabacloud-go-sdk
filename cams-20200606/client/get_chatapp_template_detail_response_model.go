// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappTemplateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatappTemplateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatappTemplateDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetChatappTemplateDetailResponseBody) *GetChatappTemplateDetailResponse
	GetBody() *GetChatappTemplateDetailResponseBody
}

type GetChatappTemplateDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappTemplateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatappTemplateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatappTemplateDetailResponse) GetBody() *GetChatappTemplateDetailResponseBody {
	return s.Body
}

func (s *GetChatappTemplateDetailResponse) SetHeaders(v map[string]*string) *GetChatappTemplateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetChatappTemplateDetailResponse) SetStatusCode(v int32) *GetChatappTemplateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponse) SetBody(v *GetChatappTemplateDetailResponseBody) *GetChatappTemplateDetailResponse {
	s.Body = v
	return s
}

func (s *GetChatappTemplateDetailResponse) Validate() error {
	return dara.Validate(s)
}
