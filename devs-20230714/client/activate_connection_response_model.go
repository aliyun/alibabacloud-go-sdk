// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateConnectionResponse
	GetStatusCode() *int32
	SetBody(v *Connection) *ActivateConnectionResponse
	GetBody() *Connection
}

type ActivateConnectionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Connection        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateConnectionResponse) GoString() string {
	return s.String()
}

func (s *ActivateConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateConnectionResponse) GetBody() *Connection {
	return s.Body
}

func (s *ActivateConnectionResponse) SetHeaders(v map[string]*string) *ActivateConnectionResponse {
	s.Headers = v
	return s
}

func (s *ActivateConnectionResponse) SetStatusCode(v int32) *ActivateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateConnectionResponse) SetBody(v *Connection) *ActivateConnectionResponse {
	s.Body = v
	return s
}

func (s *ActivateConnectionResponse) Validate() error {
	return dara.Validate(s)
}
