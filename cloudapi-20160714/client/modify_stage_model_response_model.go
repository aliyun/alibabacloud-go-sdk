// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStageModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStageModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStageModelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStageModelResponseBody) *ModifyStageModelResponse
	GetBody() *ModifyStageModelResponseBody
}

type ModifyStageModelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStageModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStageModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStageModelResponse) GoString() string {
	return s.String()
}

func (s *ModifyStageModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStageModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStageModelResponse) GetBody() *ModifyStageModelResponseBody {
	return s.Body
}

func (s *ModifyStageModelResponse) SetHeaders(v map[string]*string) *ModifyStageModelResponse {
	s.Headers = v
	return s
}

func (s *ModifyStageModelResponse) SetStatusCode(v int32) *ModifyStageModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStageModelResponse) SetBody(v *ModifyStageModelResponseBody) *ModifyStageModelResponse {
	s.Body = v
	return s
}

func (s *ModifyStageModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
