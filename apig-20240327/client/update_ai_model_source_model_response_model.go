// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelSourceModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAiModelSourceModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAiModelSourceModelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAiModelSourceModelResponseBody) *UpdateAiModelSourceModelResponse
	GetBody() *UpdateAiModelSourceModelResponseBody
}

type UpdateAiModelSourceModelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiModelSourceModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAiModelSourceModelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelSourceModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiModelSourceModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAiModelSourceModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAiModelSourceModelResponse) GetBody() *UpdateAiModelSourceModelResponseBody {
	return s.Body
}

func (s *UpdateAiModelSourceModelResponse) SetHeaders(v map[string]*string) *UpdateAiModelSourceModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiModelSourceModelResponse) SetStatusCode(v int32) *UpdateAiModelSourceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiModelSourceModelResponse) SetBody(v *UpdateAiModelSourceModelResponseBody) *UpdateAiModelSourceModelResponse {
	s.Body = v
	return s
}

func (s *UpdateAiModelSourceModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
