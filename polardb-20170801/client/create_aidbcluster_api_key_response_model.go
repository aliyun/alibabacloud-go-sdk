// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAIDBClusterApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAIDBClusterApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAIDBClusterApiKeyResponseBody) *CreateAIDBClusterApiKeyResponse
	GetBody() *CreateAIDBClusterApiKeyResponseBody
}

type CreateAIDBClusterApiKeyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAIDBClusterApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAIDBClusterApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterApiKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAIDBClusterApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAIDBClusterApiKeyResponse) GetBody() *CreateAIDBClusterApiKeyResponseBody {
	return s.Body
}

func (s *CreateAIDBClusterApiKeyResponse) SetHeaders(v map[string]*string) *CreateAIDBClusterApiKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAIDBClusterApiKeyResponse) SetStatusCode(v int32) *CreateAIDBClusterApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAIDBClusterApiKeyResponse) SetBody(v *CreateAIDBClusterApiKeyResponseBody) *CreateAIDBClusterApiKeyResponse {
	s.Body = v
	return s
}

func (s *CreateAIDBClusterApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
