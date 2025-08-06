// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodDomainResponseBody) *DeleteVodDomainResponse
	GetBody() *DeleteVodDomainResponseBody
}

type DeleteVodDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodDomainResponse) GetBody() *DeleteVodDomainResponseBody {
	return s.Body
}

func (s *DeleteVodDomainResponse) SetHeaders(v map[string]*string) *DeleteVodDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodDomainResponse) SetStatusCode(v int32) *DeleteVodDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodDomainResponse) SetBody(v *DeleteVodDomainResponseBody) *DeleteVodDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteVodDomainResponse) Validate() error {
	return dara.Validate(s)
}
