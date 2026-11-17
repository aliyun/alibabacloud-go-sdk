// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAiModelSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAiModelSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAiModelSourceResponseBody) *DeleteAiModelSourceResponse
	GetBody() *DeleteAiModelSourceResponseBody
}

type DeleteAiModelSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAiModelSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAiModelSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAiModelSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAiModelSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAiModelSourceResponse) GetBody() *DeleteAiModelSourceResponseBody {
	return s.Body
}

func (s *DeleteAiModelSourceResponse) SetHeaders(v map[string]*string) *DeleteAiModelSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAiModelSourceResponse) SetStatusCode(v int32) *DeleteAiModelSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAiModelSourceResponse) SetBody(v *DeleteAiModelSourceResponseBody) *DeleteAiModelSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteAiModelSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
