// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveDomainResponseBody) *DeleteLiveDomainResponse
	GetBody() *DeleteLiveDomainResponseBody
}

type DeleteLiveDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveDomainResponse) GetBody() *DeleteLiveDomainResponseBody {
	return s.Body
}

func (s *DeleteLiveDomainResponse) SetHeaders(v map[string]*string) *DeleteLiveDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveDomainResponse) SetStatusCode(v int32) *DeleteLiveDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveDomainResponse) SetBody(v *DeleteLiveDomainResponseBody) *DeleteLiveDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveDomainResponse) Validate() error {
	return dara.Validate(s)
}
