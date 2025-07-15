// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateAttendedTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitiateAttendedTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitiateAttendedTransferResponse
	GetStatusCode() *int32
	SetBody(v *InitiateAttendedTransferResponseBody) *InitiateAttendedTransferResponse
	GetBody() *InitiateAttendedTransferResponseBody
}

type InitiateAttendedTransferResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitiateAttendedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitiateAttendedTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferResponse) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitiateAttendedTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitiateAttendedTransferResponse) GetBody() *InitiateAttendedTransferResponseBody {
	return s.Body
}

func (s *InitiateAttendedTransferResponse) SetHeaders(v map[string]*string) *InitiateAttendedTransferResponse {
	s.Headers = v
	return s
}

func (s *InitiateAttendedTransferResponse) SetStatusCode(v int32) *InitiateAttendedTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *InitiateAttendedTransferResponse) SetBody(v *InitiateAttendedTransferResponseBody) *InitiateAttendedTransferResponse {
	s.Body = v
	return s
}

func (s *InitiateAttendedTransferResponse) Validate() error {
	return dara.Validate(s)
}
