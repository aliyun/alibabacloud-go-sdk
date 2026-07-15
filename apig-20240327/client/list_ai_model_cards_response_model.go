// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelCardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiModelCardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiModelCardsResponse
	GetStatusCode() *int32
	SetBody(v *ListAiModelCardsResponseBody) *ListAiModelCardsResponse
	GetBody() *ListAiModelCardsResponseBody
}

type ListAiModelCardsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiModelCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiModelCardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsResponse) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiModelCardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiModelCardsResponse) GetBody() *ListAiModelCardsResponseBody {
	return s.Body
}

func (s *ListAiModelCardsResponse) SetHeaders(v map[string]*string) *ListAiModelCardsResponse {
	s.Headers = v
	return s
}

func (s *ListAiModelCardsResponse) SetStatusCode(v int32) *ListAiModelCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiModelCardsResponse) SetBody(v *ListAiModelCardsResponseBody) *ListAiModelCardsResponse {
	s.Body = v
	return s
}

func (s *ListAiModelCardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
