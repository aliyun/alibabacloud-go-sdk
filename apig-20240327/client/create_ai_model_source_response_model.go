// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAiModelSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAiModelSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAiModelSourceResponseBody) *CreateAiModelSourceResponse
	GetBody() *CreateAiModelSourceResponseBody
}

type CreateAiModelSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiModelSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAiModelSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateAiModelSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAiModelSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAiModelSourceResponse) GetBody() *CreateAiModelSourceResponseBody {
	return s.Body
}

func (s *CreateAiModelSourceResponse) SetHeaders(v map[string]*string) *CreateAiModelSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateAiModelSourceResponse) SetStatusCode(v int32) *CreateAiModelSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiModelSourceResponse) SetBody(v *CreateAiModelSourceResponseBody) *CreateAiModelSourceResponse {
	s.Body = v
	return s
}

func (s *CreateAiModelSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
