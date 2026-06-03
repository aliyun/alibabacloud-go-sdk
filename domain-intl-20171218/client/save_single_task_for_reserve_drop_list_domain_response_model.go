// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForReserveDropListDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForReserveDropListDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForReserveDropListDomainResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForReserveDropListDomainResponseBody) *SaveSingleTaskForReserveDropListDomainResponse
	GetBody() *SaveSingleTaskForReserveDropListDomainResponseBody
}

type SaveSingleTaskForReserveDropListDomainResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForReserveDropListDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForReserveDropListDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForReserveDropListDomainResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForReserveDropListDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForReserveDropListDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForReserveDropListDomainResponse) GetBody() *SaveSingleTaskForReserveDropListDomainResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForReserveDropListDomainResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForReserveDropListDomainResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainResponse) SetStatusCode(v int32) *SaveSingleTaskForReserveDropListDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainResponse) SetBody(v *SaveSingleTaskForReserveDropListDomainResponseBody) *SaveSingleTaskForReserveDropListDomainResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
