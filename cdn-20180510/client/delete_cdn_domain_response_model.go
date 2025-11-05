// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCdnDomainResponseBody) *DeleteCdnDomainResponse
	GetBody() *DeleteCdnDomainResponseBody
}

type DeleteCdnDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCdnDomainResponse) GetBody() *DeleteCdnDomainResponseBody {
	return s.Body
}

func (s *DeleteCdnDomainResponse) SetHeaders(v map[string]*string) *DeleteCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteCdnDomainResponse) SetStatusCode(v int32) *DeleteCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCdnDomainResponse) SetBody(v *DeleteCdnDomainResponseBody) *DeleteCdnDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteCdnDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
