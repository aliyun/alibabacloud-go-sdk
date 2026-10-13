// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelCardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelCardsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelCardsResponseBody) *ListModelCardsResponse
	GetBody() *ListModelCardsResponseBody
}

type ListModelCardsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelCardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelCardsResponse) GoString() string {
	return s.String()
}

func (s *ListModelCardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelCardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelCardsResponse) GetBody() *ListModelCardsResponseBody {
	return s.Body
}

func (s *ListModelCardsResponse) SetHeaders(v map[string]*string) *ListModelCardsResponse {
	s.Headers = v
	return s
}

func (s *ListModelCardsResponse) SetStatusCode(v int32) *ListModelCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelCardsResponse) SetBody(v *ListModelCardsResponseBody) *ListModelCardsResponse {
	s.Body = v
	return s
}

func (s *ListModelCardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
