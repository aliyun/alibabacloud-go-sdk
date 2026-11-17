// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseBuildTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseBuildTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBaseBuildTaskResultResponseBody) *GetKnowledgeBaseBuildTaskResultResponse
	GetBody() *GetKnowledgeBaseBuildTaskResultResponseBody
}

type GetKnowledgeBaseBuildTaskResultResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBaseBuildTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseBuildTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseBuildTaskResultResponse) GetBody() *GetKnowledgeBaseBuildTaskResultResponseBody {
	return s.Body
}

func (s *GetKnowledgeBaseBuildTaskResultResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseBuildTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponse) SetStatusCode(v int32) *GetKnowledgeBaseBuildTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponse) SetBody(v *GetKnowledgeBaseBuildTaskResultResponseBody) *GetKnowledgeBaseBuildTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
