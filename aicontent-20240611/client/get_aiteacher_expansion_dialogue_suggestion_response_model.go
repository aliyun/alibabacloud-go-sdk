// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITeacherExpansionDialogueSuggestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAITeacherExpansionDialogueSuggestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponse
	GetStatusCode() *int32
	SetBody(v *GetAITeacherExpansionDialogueSuggestionResponseBody) *GetAITeacherExpansionDialogueSuggestionResponse
	GetBody() *GetAITeacherExpansionDialogueSuggestionResponseBody
}

type GetAITeacherExpansionDialogueSuggestionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITeacherExpansionDialogueSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponse) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) GetBody() *GetAITeacherExpansionDialogueSuggestionResponseBody {
	return s.Body
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetHeaders(v map[string]*string) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.Headers = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetBody(v *GetAITeacherExpansionDialogueSuggestionResponseBody) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.Body = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
