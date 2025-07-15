// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveCenterTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveCenterTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveCenterTransferResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveCenterTransferResponseBody) *UpdateLiveCenterTransferResponse
	GetBody() *UpdateLiveCenterTransferResponseBody
}

type UpdateLiveCenterTransferResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveCenterTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveCenterTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveCenterTransferResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveCenterTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveCenterTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveCenterTransferResponse) GetBody() *UpdateLiveCenterTransferResponseBody {
	return s.Body
}

func (s *UpdateLiveCenterTransferResponse) SetHeaders(v map[string]*string) *UpdateLiveCenterTransferResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveCenterTransferResponse) SetStatusCode(v int32) *UpdateLiveCenterTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveCenterTransferResponse) SetBody(v *UpdateLiveCenterTransferResponseBody) *UpdateLiveCenterTransferResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveCenterTransferResponse) Validate() error {
	return dara.Validate(s)
}
