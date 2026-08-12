// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainMetaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainMetaResponseBody) *UpdateDomainMetaResponse
	GetBody() *UpdateDomainMetaResponseBody
}

type UpdateDomainMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainMetaResponse) GetBody() *UpdateDomainMetaResponseBody {
	return s.Body
}

func (s *UpdateDomainMetaResponse) SetHeaders(v map[string]*string) *UpdateDomainMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainMetaResponse) SetStatusCode(v int32) *UpdateDomainMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainMetaResponse) SetBody(v *UpdateDomainMetaResponseBody) *UpdateDomainMetaResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
