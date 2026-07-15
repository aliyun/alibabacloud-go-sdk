// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAiModelCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAiModelCardResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAiModelCardResponseBody) *DeleteAiModelCardResponse
	GetBody() *DeleteAiModelCardResponseBody
}

type DeleteAiModelCardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAiModelCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAiModelCardResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelCardResponse) GoString() string {
	return s.String()
}

func (s *DeleteAiModelCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAiModelCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAiModelCardResponse) GetBody() *DeleteAiModelCardResponseBody {
	return s.Body
}

func (s *DeleteAiModelCardResponse) SetHeaders(v map[string]*string) *DeleteAiModelCardResponse {
	s.Headers = v
	return s
}

func (s *DeleteAiModelCardResponse) SetStatusCode(v int32) *DeleteAiModelCardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAiModelCardResponse) SetBody(v *DeleteAiModelCardResponseBody) *DeleteAiModelCardResponse {
	s.Body = v
	return s
}

func (s *DeleteAiModelCardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
