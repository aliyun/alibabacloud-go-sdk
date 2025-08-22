// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModuleResponseBody) *DeleteModuleResponse
	GetBody() *DeleteModuleResponseBody
}

type DeleteModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModuleResponse) GetBody() *DeleteModuleResponseBody {
	return s.Body
}

func (s *DeleteModuleResponse) SetHeaders(v map[string]*string) *DeleteModuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteModuleResponse) SetStatusCode(v int32) *DeleteModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModuleResponse) SetBody(v *DeleteModuleResponseBody) *DeleteModuleResponse {
	s.Body = v
	return s
}

func (s *DeleteModuleResponse) Validate() error {
	return dara.Validate(s)
}
