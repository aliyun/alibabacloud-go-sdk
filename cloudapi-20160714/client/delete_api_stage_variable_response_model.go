// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiStageVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiStageVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiStageVariableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiStageVariableResponseBody) *DeleteApiStageVariableResponse
	GetBody() *DeleteApiStageVariableResponseBody
}

type DeleteApiStageVariableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiStageVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiStageVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiStageVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiStageVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiStageVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiStageVariableResponse) GetBody() *DeleteApiStageVariableResponseBody {
	return s.Body
}

func (s *DeleteApiStageVariableResponse) SetHeaders(v map[string]*string) *DeleteApiStageVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiStageVariableResponse) SetStatusCode(v int32) *DeleteApiStageVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiStageVariableResponse) SetBody(v *DeleteApiStageVariableResponseBody) *DeleteApiStageVariableResponse {
	s.Body = v
	return s
}

func (s *DeleteApiStageVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
