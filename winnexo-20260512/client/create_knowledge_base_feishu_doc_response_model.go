// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFeishuDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseFeishuDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseFeishuDocResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseFeishuDocResponseBody) *CreateKnowledgeBaseFeishuDocResponse
	GetBody() *CreateKnowledgeBaseFeishuDocResponseBody
}

type CreateKnowledgeBaseFeishuDocResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseFeishuDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseFeishuDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFeishuDocResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFeishuDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseFeishuDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseFeishuDocResponse) GetBody() *CreateKnowledgeBaseFeishuDocResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseFeishuDocResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseFeishuDocResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponse) SetStatusCode(v int32) *CreateKnowledgeBaseFeishuDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponse) SetBody(v *CreateKnowledgeBaseFeishuDocResponseBody) *CreateKnowledgeBaseFeishuDocResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
