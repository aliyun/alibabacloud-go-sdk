// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiModelCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiModelCardResponse
	GetStatusCode() *int32
	SetBody(v *GetAiModelCardResponseBody) *GetAiModelCardResponse
	GetBody() *GetAiModelCardResponseBody
}

type GetAiModelCardResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiModelCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiModelCardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelCardResponse) GoString() string {
	return s.String()
}

func (s *GetAiModelCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiModelCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiModelCardResponse) GetBody() *GetAiModelCardResponseBody {
	return s.Body
}

func (s *GetAiModelCardResponse) SetHeaders(v map[string]*string) *GetAiModelCardResponse {
	s.Headers = v
	return s
}

func (s *GetAiModelCardResponse) SetStatusCode(v int32) *GetAiModelCardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiModelCardResponse) SetBody(v *GetAiModelCardResponseBody) *GetAiModelCardResponse {
	s.Body = v
	return s
}

func (s *GetAiModelCardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
