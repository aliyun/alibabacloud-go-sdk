// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarImageClusterTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSimilarImageClusterTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSimilarImageClusterTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSimilarImageClusterTaskResponseBody) *CreateSimilarImageClusterTaskResponse
	GetBody() *CreateSimilarImageClusterTaskResponseBody
}

type CreateSimilarImageClusterTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimilarImageClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimilarImageClusterTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarImageClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusterTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimilarImageClusterTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimilarImageClusterTaskResponse) GetBody() *CreateSimilarImageClusterTaskResponseBody {
	return s.Body
}

func (s *CreateSimilarImageClusterTaskResponse) SetHeaders(v map[string]*string) *CreateSimilarImageClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSimilarImageClusterTaskResponse) SetStatusCode(v int32) *CreateSimilarImageClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimilarImageClusterTaskResponse) SetBody(v *CreateSimilarImageClusterTaskResponseBody) *CreateSimilarImageClusterTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSimilarImageClusterTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
