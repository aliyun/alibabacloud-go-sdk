// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITeacherSyncDialogueSuggestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAITeacherSyncDialogueSuggestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponse
	GetStatusCode() *int32
	SetBody(v *GetAITeacherSyncDialogueSuggestionResponseBody) *GetAITeacherSyncDialogueSuggestionResponse
	GetBody() *GetAITeacherSyncDialogueSuggestionResponseBody
}

type GetAITeacherSyncDialogueSuggestionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITeacherSyncDialogueSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponse) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) GetBody() *GetAITeacherSyncDialogueSuggestionResponseBody {
	return s.Body
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetHeaders(v map[string]*string) *GetAITeacherSyncDialogueSuggestionResponse {
	s.Headers = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetBody(v *GetAITeacherSyncDialogueSuggestionResponseBody) *GetAITeacherSyncDialogueSuggestionResponse {
	s.Body = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
