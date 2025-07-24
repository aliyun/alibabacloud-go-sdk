// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScriptResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScriptResponseBody) *DeleteScriptResponse
	GetBody() *DeleteScriptResponseBody
}

type DeleteScriptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptResponse) GoString() string {
	return s.String()
}

func (s *DeleteScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScriptResponse) GetBody() *DeleteScriptResponseBody {
	return s.Body
}

func (s *DeleteScriptResponse) SetHeaders(v map[string]*string) *DeleteScriptResponse {
	s.Headers = v
	return s
}

func (s *DeleteScriptResponse) SetStatusCode(v int32) *DeleteScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScriptResponse) SetBody(v *DeleteScriptResponseBody) *DeleteScriptResponse {
	s.Body = v
	return s
}

func (s *DeleteScriptResponse) Validate() error {
	return dara.Validate(s)
}
