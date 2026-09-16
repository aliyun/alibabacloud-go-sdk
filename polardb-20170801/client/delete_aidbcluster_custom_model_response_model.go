// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterCustomModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIDBClusterCustomModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIDBClusterCustomModelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIDBClusterCustomModelResponseBody) *DeleteAIDBClusterCustomModelResponse
	GetBody() *DeleteAIDBClusterCustomModelResponseBody
}

type DeleteAIDBClusterCustomModelResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIDBClusterCustomModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIDBClusterCustomModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterCustomModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterCustomModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIDBClusterCustomModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIDBClusterCustomModelResponse) GetBody() *DeleteAIDBClusterCustomModelResponseBody {
	return s.Body
}

func (s *DeleteAIDBClusterCustomModelResponse) SetHeaders(v map[string]*string) *DeleteAIDBClusterCustomModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIDBClusterCustomModelResponse) SetStatusCode(v int32) *DeleteAIDBClusterCustomModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelResponse) SetBody(v *DeleteAIDBClusterCustomModelResponseBody) *DeleteAIDBClusterCustomModelResponse {
	s.Body = v
	return s
}

func (s *DeleteAIDBClusterCustomModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
