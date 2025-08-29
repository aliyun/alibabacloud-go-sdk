// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransferInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTransferInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTransferInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTransferInfoResponseBody) *GetTransferInfoResponse
	GetBody() *GetTransferInfoResponseBody
}

type GetTransferInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTransferInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTransferInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTransferInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTransferInfoResponse) GetBody() *GetTransferInfoResponseBody {
	return s.Body
}

func (s *GetTransferInfoResponse) SetHeaders(v map[string]*string) *GetTransferInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTransferInfoResponse) SetStatusCode(v int32) *GetTransferInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTransferInfoResponse) SetBody(v *GetTransferInfoResponseBody) *GetTransferInfoResponse {
	s.Body = v
	return s
}

func (s *GetTransferInfoResponse) Validate() error {
	return dara.Validate(s)
}
