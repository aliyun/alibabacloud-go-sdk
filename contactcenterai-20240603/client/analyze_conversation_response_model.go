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
	SetBody(v *AnalyzeConversationResponseBody) *AnalyzeConversationResponse
	GetBody() *AnalyzeConversationResponseBody
}

type AnalyzeConversationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *AnalyzeConversationResponse) SetBody(v *AnalyzeConversationResponseBody) *AnalyzeConversationResponse {
	s.Body = v
	return s
}

func (s *AnalyzeConversationResponse) Validate() error {
	return dara.Validate(s)
}
