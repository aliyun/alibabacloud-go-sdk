// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChainResponse
	GetStatusCode() *int32
	SetBody(v *GetChainResponseBody) *GetChainResponse
	GetBody() *GetChainResponseBody
}

type GetChainResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChainResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponse) GoString() string {
	return s.String()
}

func (s *GetChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChainResponse) GetBody() *GetChainResponseBody {
	return s.Body
}

func (s *GetChainResponse) SetHeaders(v map[string]*string) *GetChainResponse {
	s.Headers = v
	return s
}

func (s *GetChainResponse) SetStatusCode(v int32) *GetChainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChainResponse) SetBody(v *GetChainResponseBody) *GetChainResponse {
	s.Body = v
	return s
}

func (s *GetChainResponse) Validate() error {
	return dara.Validate(s)
}
