// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnIpaDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnIpaDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnIpaDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnIpaDomainResponseBody) *DeleteDcdnIpaDomainResponse
	GetBody() *DeleteDcdnIpaDomainResponseBody
}

type DeleteDcdnIpaDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnIpaDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnIpaDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnIpaDomainResponse) GetBody() *DeleteDcdnIpaDomainResponseBody {
	return s.Body
}

func (s *DeleteDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *DeleteDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnIpaDomainResponse) SetStatusCode(v int32) *DeleteDcdnIpaDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnIpaDomainResponse) SetBody(v *DeleteDcdnIpaDomainResponseBody) *DeleteDcdnIpaDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnIpaDomainResponse) Validate() error {
	return dara.Validate(s)
}
