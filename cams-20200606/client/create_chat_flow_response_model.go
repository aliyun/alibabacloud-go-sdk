// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatFlowResponseBody) *CreateChatFlowResponse
	GetBody() *CreateChatFlowResponseBody
}

type CreateChatFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateChatFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatFlowResponse) GetBody() *CreateChatFlowResponseBody {
	return s.Body
}

func (s *CreateChatFlowResponse) SetHeaders(v map[string]*string) *CreateChatFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateChatFlowResponse) SetStatusCode(v int32) *CreateChatFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatFlowResponse) SetBody(v *CreateChatFlowResponseBody) *CreateChatFlowResponse {
	s.Body = v
	return s
}

func (s *CreateChatFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
