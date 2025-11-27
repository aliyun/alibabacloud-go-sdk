// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarImageClusteringTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSimilarImageClusteringTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSimilarImageClusteringTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSimilarImageClusteringTaskResponseBody) *CreateSimilarImageClusteringTaskResponse
	GetBody() *CreateSimilarImageClusteringTaskResponseBody
}

type CreateSimilarImageClusteringTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimilarImageClusteringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimilarImageClusteringTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimilarImageClusteringTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimilarImageClusteringTaskResponse) GetBody() *CreateSimilarImageClusteringTaskResponseBody {
	return s.Body
}

func (s *CreateSimilarImageClusteringTaskResponse) SetHeaders(v map[string]*string) *CreateSimilarImageClusteringTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponse) SetStatusCode(v int32) *CreateSimilarImageClusteringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponse) SetBody(v *CreateSimilarImageClusteringTaskResponseBody) *CreateSimilarImageClusteringTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
