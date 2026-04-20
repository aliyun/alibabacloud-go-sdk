// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIDBClusterTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIDBClusterTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIDBClusterTaskResponseBody) *DeleteAIDBClusterTaskResponse
	GetBody() *DeleteAIDBClusterTaskResponseBody
}

type DeleteAIDBClusterTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIDBClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIDBClusterTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIDBClusterTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIDBClusterTaskResponse) GetBody() *DeleteAIDBClusterTaskResponseBody {
	return s.Body
}

func (s *DeleteAIDBClusterTaskResponse) SetHeaders(v map[string]*string) *DeleteAIDBClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIDBClusterTaskResponse) SetStatusCode(v int32) *DeleteAIDBClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIDBClusterTaskResponse) SetBody(v *DeleteAIDBClusterTaskResponseBody) *DeleteAIDBClusterTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteAIDBClusterTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
