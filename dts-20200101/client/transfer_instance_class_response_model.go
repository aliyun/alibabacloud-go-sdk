// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInstanceClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferInstanceClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferInstanceClassResponse
	GetStatusCode() *int32
	SetBody(v *TransferInstanceClassResponseBody) *TransferInstanceClassResponse
	GetBody() *TransferInstanceClassResponseBody
}

type TransferInstanceClassResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferInstanceClassResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferInstanceClassResponse) GoString() string {
	return s.String()
}

func (s *TransferInstanceClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferInstanceClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferInstanceClassResponse) GetBody() *TransferInstanceClassResponseBody {
	return s.Body
}

func (s *TransferInstanceClassResponse) SetHeaders(v map[string]*string) *TransferInstanceClassResponse {
	s.Headers = v
	return s
}

func (s *TransferInstanceClassResponse) SetStatusCode(v int32) *TransferInstanceClassResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInstanceClassResponse) SetBody(v *TransferInstanceClassResponseBody) *TransferInstanceClassResponse {
	s.Body = v
	return s
}

func (s *TransferInstanceClassResponse) Validate() error {
	return dara.Validate(s)
}
