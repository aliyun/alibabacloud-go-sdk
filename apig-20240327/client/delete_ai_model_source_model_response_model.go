// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelSourceModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAiModelSourceModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAiModelSourceModelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAiModelSourceModelResponseBody) *DeleteAiModelSourceModelResponse
	GetBody() *DeleteAiModelSourceModelResponseBody
}

type DeleteAiModelSourceModelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAiModelSourceModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAiModelSourceModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelSourceModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteAiModelSourceModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAiModelSourceModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAiModelSourceModelResponse) GetBody() *DeleteAiModelSourceModelResponseBody {
	return s.Body
}

func (s *DeleteAiModelSourceModelResponse) SetHeaders(v map[string]*string) *DeleteAiModelSourceModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteAiModelSourceModelResponse) SetStatusCode(v int32) *DeleteAiModelSourceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAiModelSourceModelResponse) SetBody(v *DeleteAiModelSourceModelResponseBody) *DeleteAiModelSourceModelResponse {
	s.Body = v
	return s
}

func (s *DeleteAiModelSourceModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
