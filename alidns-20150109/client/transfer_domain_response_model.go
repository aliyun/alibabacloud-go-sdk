// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferDomainResponse
	GetStatusCode() *int32
	SetBody(v *TransferDomainResponseBody) *TransferDomainResponse
	GetBody() *TransferDomainResponseBody
}

type TransferDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferDomainResponse) GoString() string {
	return s.String()
}

func (s *TransferDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferDomainResponse) GetBody() *TransferDomainResponseBody {
	return s.Body
}

func (s *TransferDomainResponse) SetHeaders(v map[string]*string) *TransferDomainResponse {
	s.Headers = v
	return s
}

func (s *TransferDomainResponse) SetStatusCode(v int32) *TransferDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferDomainResponse) SetBody(v *TransferDomainResponseBody) *TransferDomainResponse {
	s.Body = v
	return s
}

func (s *TransferDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
