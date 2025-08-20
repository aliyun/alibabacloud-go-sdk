// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChainResponse
	GetStatusCode() *int32
	SetBody(v *CreateChainResponseBody) *CreateChainResponse
	GetBody() *CreateChainResponseBody
}

type CreateChainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChainResponse) GoString() string {
	return s.String()
}

func (s *CreateChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChainResponse) GetBody() *CreateChainResponseBody {
	return s.Body
}

func (s *CreateChainResponse) SetHeaders(v map[string]*string) *CreateChainResponse {
	s.Headers = v
	return s
}

func (s *CreateChainResponse) SetStatusCode(v int32) *CreateChainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChainResponse) SetBody(v *CreateChainResponseBody) *CreateChainResponse {
	s.Body = v
	return s
}

func (s *CreateChainResponse) Validate() error {
	return dara.Validate(s)
}
