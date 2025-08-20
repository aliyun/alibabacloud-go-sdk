// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChainResponseBody) *UpdateChainResponse
	GetBody() *UpdateChainResponseBody
}

type UpdateChainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChainResponse) GoString() string {
	return s.String()
}

func (s *UpdateChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChainResponse) GetBody() *UpdateChainResponseBody {
	return s.Body
}

func (s *UpdateChainResponse) SetHeaders(v map[string]*string) *UpdateChainResponse {
	s.Headers = v
	return s
}

func (s *UpdateChainResponse) SetStatusCode(v int32) *UpdateChainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChainResponse) SetBody(v *UpdateChainResponseBody) *UpdateChainResponse {
	s.Body = v
	return s
}

func (s *UpdateChainResponse) Validate() error {
	return dara.Validate(s)
}
