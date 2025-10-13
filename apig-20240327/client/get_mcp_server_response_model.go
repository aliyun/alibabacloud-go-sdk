// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpServerResponseBody) *GetMcpServerResponse
	GetBody() *GetMcpServerResponseBody
}

type GetMcpServerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponse) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpServerResponse) GetBody() *GetMcpServerResponseBody {
	return s.Body
}

func (s *GetMcpServerResponse) SetHeaders(v map[string]*string) *GetMcpServerResponse {
	s.Headers = v
	return s
}

func (s *GetMcpServerResponse) SetStatusCode(v int32) *GetMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpServerResponse) SetBody(v *GetMcpServerResponseBody) *GetMcpServerResponse {
	s.Body = v
	return s
}

func (s *GetMcpServerResponse) Validate() error {
	return dara.Validate(s)
}
