// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSelfBindVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSelfBindVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSelfBindVariableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSelfBindVariableResponseBody) *DeleteSelfBindVariableResponse
	GetBody() *DeleteSelfBindVariableResponseBody
}

type DeleteSelfBindVariableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSelfBindVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSelfBindVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSelfBindVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteSelfBindVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSelfBindVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSelfBindVariableResponse) GetBody() *DeleteSelfBindVariableResponseBody {
	return s.Body
}

func (s *DeleteSelfBindVariableResponse) SetHeaders(v map[string]*string) *DeleteSelfBindVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteSelfBindVariableResponse) SetStatusCode(v int32) *DeleteSelfBindVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSelfBindVariableResponse) SetBody(v *DeleteSelfBindVariableResponseBody) *DeleteSelfBindVariableResponse {
	s.Body = v
	return s
}

func (s *DeleteSelfBindVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
