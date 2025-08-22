// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnDomainResponseBody) *DeleteDcdnDomainResponse
	GetBody() *DeleteDcdnDomainResponseBody
}

type DeleteDcdnDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnDomainResponse) GetBody() *DeleteDcdnDomainResponseBody {
	return s.Body
}

func (s *DeleteDcdnDomainResponse) SetHeaders(v map[string]*string) *DeleteDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnDomainResponse) SetStatusCode(v int32) *DeleteDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnDomainResponse) SetBody(v *DeleteDcdnDomainResponseBody) *DeleteDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
