// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelSourceModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiModelSourceModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiModelSourceModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListAiModelSourceModelsResponseBody) *ListAiModelSourceModelsResponse
	GetBody() *ListAiModelSourceModelsResponseBody
}

type ListAiModelSourceModelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiModelSourceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiModelSourceModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelSourceModelsResponse) GoString() string {
	return s.String()
}

func (s *ListAiModelSourceModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiModelSourceModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiModelSourceModelsResponse) GetBody() *ListAiModelSourceModelsResponseBody {
	return s.Body
}

func (s *ListAiModelSourceModelsResponse) SetHeaders(v map[string]*string) *ListAiModelSourceModelsResponse {
	s.Headers = v
	return s
}

func (s *ListAiModelSourceModelsResponse) SetStatusCode(v int32) *ListAiModelSourceModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiModelSourceModelsResponse) SetBody(v *ListAiModelSourceModelsResponseBody) *ListAiModelSourceModelsResponse {
	s.Body = v
	return s
}

func (s *ListAiModelSourceModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
