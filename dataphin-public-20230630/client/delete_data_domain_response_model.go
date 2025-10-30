// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataDomainResponseBody) *DeleteDataDomainResponse
	GetBody() *DeleteDataDomainResponseBody
}

type DeleteDataDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataDomainResponse) GetBody() *DeleteDataDomainResponseBody {
	return s.Body
}

func (s *DeleteDataDomainResponse) SetHeaders(v map[string]*string) *DeleteDataDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataDomainResponse) SetStatusCode(v int32) *DeleteDataDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataDomainResponse) SetBody(v *DeleteDataDomainResponseBody) *DeleteDataDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteDataDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
