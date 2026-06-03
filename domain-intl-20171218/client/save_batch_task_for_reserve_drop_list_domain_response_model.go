// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForReserveDropListDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForReserveDropListDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForReserveDropListDomainResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForReserveDropListDomainResponseBody) *SaveBatchTaskForReserveDropListDomainResponse
	GetBody() *SaveBatchTaskForReserveDropListDomainResponseBody
}

type SaveBatchTaskForReserveDropListDomainResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForReserveDropListDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForReserveDropListDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) GetBody() *SaveBatchTaskForReserveDropListDomainResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForReserveDropListDomainResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) SetStatusCode(v int32) *SaveBatchTaskForReserveDropListDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) SetBody(v *SaveBatchTaskForReserveDropListDomainResponseBody) *SaveBatchTaskForReserveDropListDomainResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
