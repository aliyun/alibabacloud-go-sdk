// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStageModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStageModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStageModelResponse
	GetStatusCode() *int32
	SetBody(v *CreateStageModelResponseBody) *CreateStageModelResponse
	GetBody() *CreateStageModelResponseBody
}

type CreateStageModelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStageModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStageModelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStageModelResponse) GoString() string {
	return s.String()
}

func (s *CreateStageModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStageModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStageModelResponse) GetBody() *CreateStageModelResponseBody {
	return s.Body
}

func (s *CreateStageModelResponse) SetHeaders(v map[string]*string) *CreateStageModelResponse {
	s.Headers = v
	return s
}

func (s *CreateStageModelResponse) SetStatusCode(v int32) *CreateStageModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStageModelResponse) SetBody(v *CreateStageModelResponseBody) *CreateStageModelResponse {
	s.Body = v
	return s
}

func (s *CreateStageModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
