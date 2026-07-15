// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAiModelCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAiModelCardResponse
	GetStatusCode() *int32
	SetBody(v *CreateAiModelCardResponseBody) *CreateAiModelCardResponse
	GetBody() *CreateAiModelCardResponseBody
}

type CreateAiModelCardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiModelCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAiModelCardResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardResponse) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAiModelCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAiModelCardResponse) GetBody() *CreateAiModelCardResponseBody {
	return s.Body
}

func (s *CreateAiModelCardResponse) SetHeaders(v map[string]*string) *CreateAiModelCardResponse {
	s.Headers = v
	return s
}

func (s *CreateAiModelCardResponse) SetStatusCode(v int32) *CreateAiModelCardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiModelCardResponse) SetBody(v *CreateAiModelCardResponseBody) *CreateAiModelCardResponse {
	s.Body = v
	return s
}

func (s *CreateAiModelCardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
