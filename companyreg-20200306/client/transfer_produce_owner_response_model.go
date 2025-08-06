// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferProduceOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferProduceOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferProduceOwnerResponse
	GetStatusCode() *int32
	SetBody(v *TransferProduceOwnerResponseBody) *TransferProduceOwnerResponse
	GetBody() *TransferProduceOwnerResponseBody
}

type TransferProduceOwnerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferProduceOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferProduceOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferProduceOwnerResponse) GoString() string {
	return s.String()
}

func (s *TransferProduceOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferProduceOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferProduceOwnerResponse) GetBody() *TransferProduceOwnerResponseBody {
	return s.Body
}

func (s *TransferProduceOwnerResponse) SetHeaders(v map[string]*string) *TransferProduceOwnerResponse {
	s.Headers = v
	return s
}

func (s *TransferProduceOwnerResponse) SetStatusCode(v int32) *TransferProduceOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferProduceOwnerResponse) SetBody(v *TransferProduceOwnerResponseBody) *TransferProduceOwnerResponse {
	s.Body = v
	return s
}

func (s *TransferProduceOwnerResponse) Validate() error {
	return dara.Validate(s)
}
