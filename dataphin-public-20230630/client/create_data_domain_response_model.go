// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataDomainResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataDomainResponseBody) *CreateDataDomainResponse
	GetBody() *CreateDataDomainResponseBody
}

type CreateDataDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDataDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataDomainResponse) GetBody() *CreateDataDomainResponseBody {
	return s.Body
}

func (s *CreateDataDomainResponse) SetHeaders(v map[string]*string) *CreateDataDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDataDomainResponse) SetStatusCode(v int32) *CreateDataDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataDomainResponse) SetBody(v *CreateDataDomainResponseBody) *CreateDataDomainResponse {
	s.Body = v
	return s
}

func (s *CreateDataDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
