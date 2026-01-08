// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFlowTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatFlowTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatFlowTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetChatFlowTemplateResponseBody) *GetChatFlowTemplateResponse
	GetBody() *GetChatFlowTemplateResponseBody
}

type GetChatFlowTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatFlowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatFlowTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetChatFlowTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatFlowTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatFlowTemplateResponse) GetBody() *GetChatFlowTemplateResponseBody {
	return s.Body
}

func (s *GetChatFlowTemplateResponse) SetHeaders(v map[string]*string) *GetChatFlowTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetChatFlowTemplateResponse) SetStatusCode(v int32) *GetChatFlowTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatFlowTemplateResponse) SetBody(v *GetChatFlowTemplateResponseBody) *GetChatFlowTemplateResponse {
	s.Body = v
	return s
}

func (s *GetChatFlowTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
