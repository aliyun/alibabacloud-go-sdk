// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayTransferStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlipayTransferStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlipayTransferStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAlipayTransferStatusResponseBody) *GetAlipayTransferStatusResponse
	GetBody() *GetAlipayTransferStatusResponseBody
}

type GetAlipayTransferStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlipayTransferStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlipayTransferStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayTransferStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAlipayTransferStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlipayTransferStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlipayTransferStatusResponse) GetBody() *GetAlipayTransferStatusResponseBody {
	return s.Body
}

func (s *GetAlipayTransferStatusResponse) SetHeaders(v map[string]*string) *GetAlipayTransferStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAlipayTransferStatusResponse) SetStatusCode(v int32) *GetAlipayTransferStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlipayTransferStatusResponse) SetBody(v *GetAlipayTransferStatusResponseBody) *GetAlipayTransferStatusResponse {
	s.Body = v
	return s
}

func (s *GetAlipayTransferStatusResponse) Validate() error {
	return dara.Validate(s)
}
