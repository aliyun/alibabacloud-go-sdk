// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelSourceModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAiModelSourceModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAiModelSourceModelResponse
	GetStatusCode() *int32
	SetBody(v *CreateAiModelSourceModelResponseBody) *CreateAiModelSourceModelResponse
	GetBody() *CreateAiModelSourceModelResponseBody
}

type CreateAiModelSourceModelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiModelSourceModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAiModelSourceModelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelSourceModelResponse) GoString() string {
	return s.String()
}

func (s *CreateAiModelSourceModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAiModelSourceModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAiModelSourceModelResponse) GetBody() *CreateAiModelSourceModelResponseBody {
	return s.Body
}

func (s *CreateAiModelSourceModelResponse) SetHeaders(v map[string]*string) *CreateAiModelSourceModelResponse {
	s.Headers = v
	return s
}

func (s *CreateAiModelSourceModelResponse) SetStatusCode(v int32) *CreateAiModelSourceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiModelSourceModelResponse) SetBody(v *CreateAiModelSourceModelResponseBody) *CreateAiModelSourceModelResponse {
	s.Body = v
	return s
}

func (s *CreateAiModelSourceModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
