// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeConversationResponse
	GetStatusCode() *int32
	SetId(v string) *AnalyzeConversationResponse
	GetId() *string
	SetEvent(v string) *AnalyzeConversationResponse
	GetEvent() *string
	SetBody(v *AnalyzeConversationResponseBody) *AnalyzeConversationResponse
	GetBody() *AnalyzeConversationResponseBody
}

type AnalyzeConversationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                          `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                          `json:"event,omitempty" xml:"event,omitempty"`
	Body       *AnalyzeConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeConversationResponse) GetId() *string {
	return s.Id
}

func (s *AnalyzeConversationResponse) GetEvent() *string {
	return s.Event
}

func (s *AnalyzeConversationResponse) GetBody() *AnalyzeConversationResponseBody {
	return s.Body
}

func (s *AnalyzeConversationResponse) SetHeaders(v map[string]*string) *AnalyzeConversationResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeConversationResponse) SetStatusCode(v int32) *AnalyzeConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeConversationResponse) SetId(v string) *AnalyzeConversationResponse {
	s.Id = &v
	return s
}

func (s *AnalyzeConversationResponse) SetEvent(v string) *AnalyzeConversationResponse {
	s.Event = &v
	return s
}

func (s *AnalyzeConversationResponse) SetBody(v *AnalyzeConversationResponseBody) *AnalyzeConversationResponse {
	s.Body = v
	return s
}

func (s *AnalyzeConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
