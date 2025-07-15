// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlindTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BlindTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BlindTransferResponse
	GetStatusCode() *int32
	SetBody(v *BlindTransferResponseBody) *BlindTransferResponse
	GetBody() *BlindTransferResponseBody
}

type BlindTransferResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BlindTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BlindTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s BlindTransferResponse) GoString() string {
	return s.String()
}

func (s *BlindTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BlindTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BlindTransferResponse) GetBody() *BlindTransferResponseBody {
	return s.Body
}

func (s *BlindTransferResponse) SetHeaders(v map[string]*string) *BlindTransferResponse {
	s.Headers = v
	return s
}

func (s *BlindTransferResponse) SetStatusCode(v int32) *BlindTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *BlindTransferResponse) SetBody(v *BlindTransferResponseBody) *BlindTransferResponse {
	s.Body = v
	return s
}

func (s *BlindTransferResponse) Validate() error {
	return dara.Validate(s)
}
