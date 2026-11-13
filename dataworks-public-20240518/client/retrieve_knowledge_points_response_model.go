// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgePointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrieveKnowledgePointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrieveKnowledgePointsResponse
	GetStatusCode() *int32
	SetBody(v *RetrieveKnowledgePointsResponseBody) *RetrieveKnowledgePointsResponse
	GetBody() *RetrieveKnowledgePointsResponseBody
}

type RetrieveKnowledgePointsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveKnowledgePointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveKnowledgePointsResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponse) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrieveKnowledgePointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrieveKnowledgePointsResponse) GetBody() *RetrieveKnowledgePointsResponseBody {
	return s.Body
}

func (s *RetrieveKnowledgePointsResponse) SetHeaders(v map[string]*string) *RetrieveKnowledgePointsResponse {
	s.Headers = v
	return s
}

func (s *RetrieveKnowledgePointsResponse) SetStatusCode(v int32) *RetrieveKnowledgePointsResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveKnowledgePointsResponse) SetBody(v *RetrieveKnowledgePointsResponseBody) *RetrieveKnowledgePointsResponse {
	s.Body = v
	return s
}

func (s *RetrieveKnowledgePointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
