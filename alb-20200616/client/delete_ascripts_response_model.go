// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAScriptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAScriptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAScriptsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAScriptsResponseBody) *DeleteAScriptsResponse
	GetBody() *DeleteAScriptsResponseBody
}

type DeleteAScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAScriptsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAScriptsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAScriptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAScriptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAScriptsResponse) GetBody() *DeleteAScriptsResponseBody {
	return s.Body
}

func (s *DeleteAScriptsResponse) SetHeaders(v map[string]*string) *DeleteAScriptsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAScriptsResponse) SetStatusCode(v int32) *DeleteAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAScriptsResponse) SetBody(v *DeleteAScriptsResponseBody) *DeleteAScriptsResponse {
	s.Body = v
	return s
}

func (s *DeleteAScriptsResponse) Validate() error {
	return dara.Validate(s)
}
