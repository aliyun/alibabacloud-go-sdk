// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStageModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStageModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStageModelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStageModelResponseBody) *DeleteStageModelResponse
	GetBody() *DeleteStageModelResponseBody
}

type DeleteStageModelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStageModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStageModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStageModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteStageModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStageModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStageModelResponse) GetBody() *DeleteStageModelResponseBody {
	return s.Body
}

func (s *DeleteStageModelResponse) SetHeaders(v map[string]*string) *DeleteStageModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteStageModelResponse) SetStatusCode(v int32) *DeleteStageModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStageModelResponse) SetBody(v *DeleteStageModelResponseBody) *DeleteStageModelResponse {
	s.Body = v
	return s
}

func (s *DeleteStageModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
