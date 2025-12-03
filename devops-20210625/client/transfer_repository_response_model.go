// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *TransferRepositoryResponseBody) *TransferRepositoryResponse
	GetBody() *TransferRepositoryResponseBody
}

type TransferRepositoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferRepositoryResponse) GoString() string {
	return s.String()
}

func (s *TransferRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferRepositoryResponse) GetBody() *TransferRepositoryResponseBody {
	return s.Body
}

func (s *TransferRepositoryResponse) SetHeaders(v map[string]*string) *TransferRepositoryResponse {
	s.Headers = v
	return s
}

func (s *TransferRepositoryResponse) SetStatusCode(v int32) *TransferRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferRepositoryResponse) SetBody(v *TransferRepositoryResponseBody) *TransferRepositoryResponse {
	s.Body = v
	return s
}

func (s *TransferRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
