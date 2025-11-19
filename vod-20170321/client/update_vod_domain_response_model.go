// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVodDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVodDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVodDomainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVodDomainResponseBody) *UpdateVodDomainResponse
	GetBody() *UpdateVodDomainResponseBody
}

type UpdateVodDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVodDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVodDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateVodDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVodDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVodDomainResponse) GetBody() *UpdateVodDomainResponseBody {
	return s.Body
}

func (s *UpdateVodDomainResponse) SetHeaders(v map[string]*string) *UpdateVodDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateVodDomainResponse) SetStatusCode(v int32) *UpdateVodDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVodDomainResponse) SetBody(v *UpdateVodDomainResponseBody) *UpdateVodDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateVodDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
