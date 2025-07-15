// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAttendedTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAttendedTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAttendedTransferResponse
	GetStatusCode() *int32
	SetBody(v *CancelAttendedTransferResponseBody) *CancelAttendedTransferResponse
	GetBody() *CancelAttendedTransferResponseBody
}

type CancelAttendedTransferResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAttendedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAttendedTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAttendedTransferResponse) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAttendedTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAttendedTransferResponse) GetBody() *CancelAttendedTransferResponseBody {
	return s.Body
}

func (s *CancelAttendedTransferResponse) SetHeaders(v map[string]*string) *CancelAttendedTransferResponse {
	s.Headers = v
	return s
}

func (s *CancelAttendedTransferResponse) SetStatusCode(v int32) *CancelAttendedTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAttendedTransferResponse) SetBody(v *CancelAttendedTransferResponseBody) *CancelAttendedTransferResponse {
	s.Body = v
	return s
}

func (s *CancelAttendedTransferResponse) Validate() error {
	return dara.Validate(s)
}
