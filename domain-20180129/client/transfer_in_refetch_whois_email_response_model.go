// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInRefetchWhoisEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferInRefetchWhoisEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferInRefetchWhoisEmailResponse
	GetStatusCode() *int32
	SetBody(v *TransferInRefetchWhoisEmailResponseBody) *TransferInRefetchWhoisEmailResponse
	GetBody() *TransferInRefetchWhoisEmailResponseBody
}

type TransferInRefetchWhoisEmailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferInRefetchWhoisEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferInRefetchWhoisEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferInRefetchWhoisEmailResponse) GoString() string {
	return s.String()
}

func (s *TransferInRefetchWhoisEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferInRefetchWhoisEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferInRefetchWhoisEmailResponse) GetBody() *TransferInRefetchWhoisEmailResponseBody {
	return s.Body
}

func (s *TransferInRefetchWhoisEmailResponse) SetHeaders(v map[string]*string) *TransferInRefetchWhoisEmailResponse {
	s.Headers = v
	return s
}

func (s *TransferInRefetchWhoisEmailResponse) SetStatusCode(v int32) *TransferInRefetchWhoisEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInRefetchWhoisEmailResponse) SetBody(v *TransferInRefetchWhoisEmailResponseBody) *TransferInRefetchWhoisEmailResponse {
	s.Body = v
	return s
}

func (s *TransferInRefetchWhoisEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
