// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAiModelCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAiModelCardResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAiModelCardResponseBody) *UpdateAiModelCardResponse
	GetBody() *UpdateAiModelCardResponseBody
}

type UpdateAiModelCardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiModelCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAiModelCardResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAiModelCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAiModelCardResponse) GetBody() *UpdateAiModelCardResponseBody {
	return s.Body
}

func (s *UpdateAiModelCardResponse) SetHeaders(v map[string]*string) *UpdateAiModelCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiModelCardResponse) SetStatusCode(v int32) *UpdateAiModelCardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiModelCardResponse) SetBody(v *UpdateAiModelCardResponseBody) *UpdateAiModelCardResponse {
	s.Body = v
	return s
}

func (s *UpdateAiModelCardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
