// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartNodeResponse
	GetStatusCode() *int32
	SetBody(v *RestartNodeResponseBody) *RestartNodeResponse
	GetBody() *RestartNodeResponseBody
}

type RestartNodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartNodeResponse) GoString() string {
	return s.String()
}

func (s *RestartNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartNodeResponse) GetBody() *RestartNodeResponseBody {
	return s.Body
}

func (s *RestartNodeResponse) SetHeaders(v map[string]*string) *RestartNodeResponse {
	s.Headers = v
	return s
}

func (s *RestartNodeResponse) SetStatusCode(v int32) *RestartNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartNodeResponse) SetBody(v *RestartNodeResponseBody) *RestartNodeResponse {
	s.Body = v
	return s
}

func (s *RestartNodeResponse) Validate() error {
	return dara.Validate(s)
}
