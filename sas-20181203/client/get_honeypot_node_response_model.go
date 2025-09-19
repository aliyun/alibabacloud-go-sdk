// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneypotNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneypotNodeResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneypotNodeResponseBody) *GetHoneypotNodeResponse
	GetBody() *GetHoneypotNodeResponseBody
}

type GetHoneypotNodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneypotNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneypotNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotNodeResponse) GoString() string {
	return s.String()
}

func (s *GetHoneypotNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneypotNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneypotNodeResponse) GetBody() *GetHoneypotNodeResponseBody {
	return s.Body
}

func (s *GetHoneypotNodeResponse) SetHeaders(v map[string]*string) *GetHoneypotNodeResponse {
	s.Headers = v
	return s
}

func (s *GetHoneypotNodeResponse) SetStatusCode(v int32) *GetHoneypotNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneypotNodeResponse) SetBody(v *GetHoneypotNodeResponseBody) *GetHoneypotNodeResponse {
	s.Body = v
	return s
}

func (s *GetHoneypotNodeResponse) Validate() error {
	return dara.Validate(s)
}
