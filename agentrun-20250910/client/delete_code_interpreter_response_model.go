// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeInterpreterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCodeInterpreterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCodeInterpreterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCodeInterpreterResult) *DeleteCodeInterpreterResponse
	GetBody() *DeleteCodeInterpreterResult
}

type DeleteCodeInterpreterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCodeInterpreterResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCodeInterpreterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeInterpreterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCodeInterpreterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCodeInterpreterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCodeInterpreterResponse) GetBody() *DeleteCodeInterpreterResult {
	return s.Body
}

func (s *DeleteCodeInterpreterResponse) SetHeaders(v map[string]*string) *DeleteCodeInterpreterResponse {
	s.Headers = v
	return s
}

func (s *DeleteCodeInterpreterResponse) SetStatusCode(v int32) *DeleteCodeInterpreterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCodeInterpreterResponse) SetBody(v *DeleteCodeInterpreterResult) *DeleteCodeInterpreterResponse {
	s.Body = v
	return s
}

func (s *DeleteCodeInterpreterResponse) Validate() error {
	return dara.Validate(s)
}
