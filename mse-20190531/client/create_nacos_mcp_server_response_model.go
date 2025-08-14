// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNacosMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNacosMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateNacosMcpServerResponseBody) *CreateNacosMcpServerResponse
	GetBody() *CreateNacosMcpServerResponseBody
}

type CreateNacosMcpServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNacosMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNacosMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosMcpServerResponse) GoString() string {
	return s.String()
}

func (s *CreateNacosMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNacosMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNacosMcpServerResponse) GetBody() *CreateNacosMcpServerResponseBody {
	return s.Body
}

func (s *CreateNacosMcpServerResponse) SetHeaders(v map[string]*string) *CreateNacosMcpServerResponse {
	s.Headers = v
	return s
}

func (s *CreateNacosMcpServerResponse) SetStatusCode(v int32) *CreateNacosMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNacosMcpServerResponse) SetBody(v *CreateNacosMcpServerResponseBody) *CreateNacosMcpServerResponse {
	s.Body = v
	return s
}

func (s *CreateNacosMcpServerResponse) Validate() error {
	return dara.Validate(s)
}
