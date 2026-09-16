// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAIDBClusterCustomModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterAIDBClusterCustomModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterAIDBClusterCustomModelResponse
	GetStatusCode() *int32
	SetBody(v *RegisterAIDBClusterCustomModelResponseBody) *RegisterAIDBClusterCustomModelResponse
	GetBody() *RegisterAIDBClusterCustomModelResponseBody
}

type RegisterAIDBClusterCustomModelResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterAIDBClusterCustomModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterAIDBClusterCustomModelResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterAIDBClusterCustomModelResponse) GoString() string {
	return s.String()
}

func (s *RegisterAIDBClusterCustomModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterAIDBClusterCustomModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterAIDBClusterCustomModelResponse) GetBody() *RegisterAIDBClusterCustomModelResponseBody {
	return s.Body
}

func (s *RegisterAIDBClusterCustomModelResponse) SetHeaders(v map[string]*string) *RegisterAIDBClusterCustomModelResponse {
	s.Headers = v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponse) SetStatusCode(v int32) *RegisterAIDBClusterCustomModelResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponse) SetBody(v *RegisterAIDBClusterCustomModelResponseBody) *RegisterAIDBClusterCustomModelResponse {
	s.Body = v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
