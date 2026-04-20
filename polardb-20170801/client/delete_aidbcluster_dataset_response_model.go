// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIDBClusterDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIDBClusterDatasetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIDBClusterDatasetResponseBody) *DeleteAIDBClusterDatasetResponse
	GetBody() *DeleteAIDBClusterDatasetResponseBody
}

type DeleteAIDBClusterDatasetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIDBClusterDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIDBClusterDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIDBClusterDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIDBClusterDatasetResponse) GetBody() *DeleteAIDBClusterDatasetResponseBody {
	return s.Body
}

func (s *DeleteAIDBClusterDatasetResponse) SetHeaders(v map[string]*string) *DeleteAIDBClusterDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIDBClusterDatasetResponse) SetStatusCode(v int32) *DeleteAIDBClusterDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIDBClusterDatasetResponse) SetBody(v *DeleteAIDBClusterDatasetResponseBody) *DeleteAIDBClusterDatasetResponse {
	s.Body = v
	return s
}

func (s *DeleteAIDBClusterDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
