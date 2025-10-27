// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentlessAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentlessAssetResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentlessAssetResponseBody) *ListAgentlessAssetResponse
	GetBody() *ListAgentlessAssetResponseBody
}

type ListAgentlessAssetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentlessAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentlessAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessAssetResponse) GoString() string {
	return s.String()
}

func (s *ListAgentlessAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentlessAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentlessAssetResponse) GetBody() *ListAgentlessAssetResponseBody {
	return s.Body
}

func (s *ListAgentlessAssetResponse) SetHeaders(v map[string]*string) *ListAgentlessAssetResponse {
	s.Headers = v
	return s
}

func (s *ListAgentlessAssetResponse) SetStatusCode(v int32) *ListAgentlessAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentlessAssetResponse) SetBody(v *ListAgentlessAssetResponseBody) *ListAgentlessAssetResponse {
	s.Body = v
	return s
}

func (s *ListAgentlessAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
