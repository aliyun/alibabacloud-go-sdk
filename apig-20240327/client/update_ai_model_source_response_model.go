// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAiModelSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAiModelSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAiModelSourceResponseBody) *UpdateAiModelSourceResponse
	GetBody() *UpdateAiModelSourceResponseBody
}

type UpdateAiModelSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiModelSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAiModelSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiModelSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAiModelSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAiModelSourceResponse) GetBody() *UpdateAiModelSourceResponseBody {
	return s.Body
}

func (s *UpdateAiModelSourceResponse) SetHeaders(v map[string]*string) *UpdateAiModelSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiModelSourceResponse) SetStatusCode(v int32) *UpdateAiModelSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiModelSourceResponse) SetBody(v *UpdateAiModelSourceResponseBody) *UpdateAiModelSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateAiModelSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
