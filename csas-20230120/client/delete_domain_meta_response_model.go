// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainMetaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainMetaResponseBody) *DeleteDomainMetaResponse
	GetBody() *DeleteDomainMetaResponseBody
}

type DeleteDomainMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainMetaResponse) GetBody() *DeleteDomainMetaResponseBody {
	return s.Body
}

func (s *DeleteDomainMetaResponse) SetHeaders(v map[string]*string) *DeleteDomainMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainMetaResponse) SetStatusCode(v int32) *DeleteDomainMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainMetaResponse) SetBody(v *DeleteDomainMetaResponseBody) *DeleteDomainMetaResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
