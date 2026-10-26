// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentConversationLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadAgentConversationLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadAgentConversationLogResponse
	GetStatusCode() *int32
	SetBody(v *UploadAgentConversationLogResponseBody) *UploadAgentConversationLogResponse
	GetBody() *UploadAgentConversationLogResponseBody
}

type UploadAgentConversationLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAgentConversationLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAgentConversationLogResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentConversationLogResponse) GoString() string {
	return s.String()
}

func (s *UploadAgentConversationLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAgentConversationLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadAgentConversationLogResponse) GetBody() *UploadAgentConversationLogResponseBody {
	return s.Body
}

func (s *UploadAgentConversationLogResponse) SetHeaders(v map[string]*string) *UploadAgentConversationLogResponse {
	s.Headers = v
	return s
}

func (s *UploadAgentConversationLogResponse) SetStatusCode(v int32) *UploadAgentConversationLogResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAgentConversationLogResponse) SetBody(v *UploadAgentConversationLogResponseBody) *UploadAgentConversationLogResponse {
	s.Body = v
	return s
}

func (s *UploadAgentConversationLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
