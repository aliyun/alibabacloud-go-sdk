// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatQuestionRespResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatQuestionRespResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatQuestionRespResponse
	GetStatusCode() *int32
	SetBody(v *GetChatQuestionRespResponseBody) *GetChatQuestionRespResponse
	GetBody() *GetChatQuestionRespResponseBody
}

type GetChatQuestionRespResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatQuestionRespResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatQuestionRespResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatQuestionRespResponse) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatQuestionRespResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatQuestionRespResponse) GetBody() *GetChatQuestionRespResponseBody {
	return s.Body
}

func (s *GetChatQuestionRespResponse) SetHeaders(v map[string]*string) *GetChatQuestionRespResponse {
	s.Headers = v
	return s
}

func (s *GetChatQuestionRespResponse) SetStatusCode(v int32) *GetChatQuestionRespResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatQuestionRespResponse) SetBody(v *GetChatQuestionRespResponseBody) *GetChatQuestionRespResponse {
	s.Body = v
	return s
}

func (s *GetChatQuestionRespResponse) Validate() error {
	return dara.Validate(s)
}
