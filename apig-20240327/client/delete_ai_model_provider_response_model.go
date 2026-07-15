// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAiModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAiModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAiModelProviderResponseBody) *DeleteAiModelProviderResponse
	GetBody() *DeleteAiModelProviderResponseBody
}

type DeleteAiModelProviderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAiModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAiModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteAiModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAiModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAiModelProviderResponse) GetBody() *DeleteAiModelProviderResponseBody {
	return s.Body
}

func (s *DeleteAiModelProviderResponse) SetHeaders(v map[string]*string) *DeleteAiModelProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteAiModelProviderResponse) SetStatusCode(v int32) *DeleteAiModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAiModelProviderResponse) SetBody(v *DeleteAiModelProviderResponseBody) *DeleteAiModelProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteAiModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
