// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiModelProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiModelProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListAiModelProvidersResponseBody) *ListAiModelProvidersResponse
	GetBody() *ListAiModelProvidersResponseBody
}

type ListAiModelProvidersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiModelProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiModelProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiModelProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiModelProvidersResponse) GetBody() *ListAiModelProvidersResponseBody {
	return s.Body
}

func (s *ListAiModelProvidersResponse) SetHeaders(v map[string]*string) *ListAiModelProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListAiModelProvidersResponse) SetStatusCode(v int32) *ListAiModelProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiModelProvidersResponse) SetBody(v *ListAiModelProvidersResponseBody) *ListAiModelProvidersResponse {
	s.Body = v
	return s
}

func (s *ListAiModelProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
