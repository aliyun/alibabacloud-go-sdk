// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveCenterTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveCenterTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveCenterTransferResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveCenterTransferResponseBody) *AddLiveCenterTransferResponse
	GetBody() *AddLiveCenterTransferResponseBody
}

type AddLiveCenterTransferResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveCenterTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveCenterTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveCenterTransferResponse) GoString() string {
	return s.String()
}

func (s *AddLiveCenterTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveCenterTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveCenterTransferResponse) GetBody() *AddLiveCenterTransferResponseBody {
	return s.Body
}

func (s *AddLiveCenterTransferResponse) SetHeaders(v map[string]*string) *AddLiveCenterTransferResponse {
	s.Headers = v
	return s
}

func (s *AddLiveCenterTransferResponse) SetStatusCode(v int32) *AddLiveCenterTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveCenterTransferResponse) SetBody(v *AddLiveCenterTransferResponseBody) *AddLiveCenterTransferResponse {
	s.Body = v
	return s
}

func (s *AddLiveCenterTransferResponse) Validate() error {
	return dara.Validate(s)
}
