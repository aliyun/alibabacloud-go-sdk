// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataRedistributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseDataRedistributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseDataRedistributeResponse
	GetStatusCode() *int32
	SetBody(v *PauseDataRedistributeResponseBody) *PauseDataRedistributeResponse
	GetBody() *PauseDataRedistributeResponseBody
}

type PauseDataRedistributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseDataRedistributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseDataRedistributeResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseDataRedistributeResponse) GoString() string {
	return s.String()
}

func (s *PauseDataRedistributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseDataRedistributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseDataRedistributeResponse) GetBody() *PauseDataRedistributeResponseBody {
	return s.Body
}

func (s *PauseDataRedistributeResponse) SetHeaders(v map[string]*string) *PauseDataRedistributeResponse {
	s.Headers = v
	return s
}

func (s *PauseDataRedistributeResponse) SetStatusCode(v int32) *PauseDataRedistributeResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseDataRedistributeResponse) SetBody(v *PauseDataRedistributeResponseBody) *PauseDataRedistributeResponse {
	s.Body = v
	return s
}

func (s *PauseDataRedistributeResponse) Validate() error {
	return dara.Validate(s)
}
