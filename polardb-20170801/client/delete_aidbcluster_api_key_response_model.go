// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIDBClusterApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIDBClusterApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIDBClusterApiKeyResponseBody) *DeleteAIDBClusterApiKeyResponse
	GetBody() *DeleteAIDBClusterApiKeyResponseBody
}

type DeleteAIDBClusterApiKeyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIDBClusterApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIDBClusterApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterApiKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIDBClusterApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIDBClusterApiKeyResponse) GetBody() *DeleteAIDBClusterApiKeyResponseBody {
	return s.Body
}

func (s *DeleteAIDBClusterApiKeyResponse) SetHeaders(v map[string]*string) *DeleteAIDBClusterApiKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIDBClusterApiKeyResponse) SetStatusCode(v int32) *DeleteAIDBClusterApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIDBClusterApiKeyResponse) SetBody(v *DeleteAIDBClusterApiKeyResponseBody) *DeleteAIDBClusterApiKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteAIDBClusterApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
