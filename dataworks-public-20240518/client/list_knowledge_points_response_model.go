// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgePointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgePointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgePointsResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgePointsResponseBody) *ListKnowledgePointsResponse
	GetBody() *ListKnowledgePointsResponseBody
}

type ListKnowledgePointsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgePointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgePointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgePointsResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgePointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgePointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgePointsResponse) GetBody() *ListKnowledgePointsResponseBody {
	return s.Body
}

func (s *ListKnowledgePointsResponse) SetHeaders(v map[string]*string) *ListKnowledgePointsResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgePointsResponse) SetStatusCode(v int32) *ListKnowledgePointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgePointsResponse) SetBody(v *ListKnowledgePointsResponseBody) *ListKnowledgePointsResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgePointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
