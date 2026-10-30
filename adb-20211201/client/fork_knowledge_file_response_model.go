// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkKnowledgeFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForkKnowledgeFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForkKnowledgeFileResponse
	GetStatusCode() *int32
	SetBody(v *ForkKnowledgeFileResponseBody) *ForkKnowledgeFileResponse
	GetBody() *ForkKnowledgeFileResponseBody
}

type ForkKnowledgeFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForkKnowledgeFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForkKnowledgeFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ForkKnowledgeFileResponse) GoString() string {
	return s.String()
}

func (s *ForkKnowledgeFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForkKnowledgeFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForkKnowledgeFileResponse) GetBody() *ForkKnowledgeFileResponseBody {
	return s.Body
}

func (s *ForkKnowledgeFileResponse) SetHeaders(v map[string]*string) *ForkKnowledgeFileResponse {
	s.Headers = v
	return s
}

func (s *ForkKnowledgeFileResponse) SetStatusCode(v int32) *ForkKnowledgeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ForkKnowledgeFileResponse) SetBody(v *ForkKnowledgeFileResponseBody) *ForkKnowledgeFileResponse {
	s.Body = v
	return s
}

func (s *ForkKnowledgeFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
