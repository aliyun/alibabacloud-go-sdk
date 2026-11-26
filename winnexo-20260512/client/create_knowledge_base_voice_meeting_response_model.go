// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseVoiceMeetingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseVoiceMeetingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseVoiceMeetingResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseVoiceMeetingResponseBody) *CreateKnowledgeBaseVoiceMeetingResponse
	GetBody() *CreateKnowledgeBaseVoiceMeetingResponseBody
}

type CreateKnowledgeBaseVoiceMeetingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseVoiceMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseVoiceMeetingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseVoiceMeetingResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseVoiceMeetingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseVoiceMeetingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseVoiceMeetingResponse) GetBody() *CreateKnowledgeBaseVoiceMeetingResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseVoiceMeetingResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseVoiceMeetingResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponse) SetStatusCode(v int32) *CreateKnowledgeBaseVoiceMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponse) SetBody(v *CreateKnowledgeBaseVoiceMeetingResponseBody) *CreateKnowledgeBaseVoiceMeetingResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
