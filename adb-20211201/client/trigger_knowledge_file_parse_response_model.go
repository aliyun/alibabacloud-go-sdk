// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerKnowledgeFileParseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerKnowledgeFileParseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerKnowledgeFileParseResponse
	GetStatusCode() *int32
	SetBody(v *TriggerKnowledgeFileParseResponseBody) *TriggerKnowledgeFileParseResponse
	GetBody() *TriggerKnowledgeFileParseResponseBody
}

type TriggerKnowledgeFileParseResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerKnowledgeFileParseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerKnowledgeFileParseResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerKnowledgeFileParseResponse) GoString() string {
	return s.String()
}

func (s *TriggerKnowledgeFileParseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerKnowledgeFileParseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerKnowledgeFileParseResponse) GetBody() *TriggerKnowledgeFileParseResponseBody {
	return s.Body
}

func (s *TriggerKnowledgeFileParseResponse) SetHeaders(v map[string]*string) *TriggerKnowledgeFileParseResponse {
	s.Headers = v
	return s
}

func (s *TriggerKnowledgeFileParseResponse) SetStatusCode(v int32) *TriggerKnowledgeFileParseResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerKnowledgeFileParseResponse) SetBody(v *TriggerKnowledgeFileParseResponseBody) *TriggerKnowledgeFileParseResponse {
	s.Body = v
	return s
}

func (s *TriggerKnowledgeFileParseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
