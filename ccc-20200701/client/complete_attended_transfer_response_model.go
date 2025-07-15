// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteAttendedTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteAttendedTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteAttendedTransferResponse
	GetStatusCode() *int32
	SetBody(v *CompleteAttendedTransferResponseBody) *CompleteAttendedTransferResponse
	GetBody() *CompleteAttendedTransferResponseBody
}

type CompleteAttendedTransferResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteAttendedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteAttendedTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferResponse) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteAttendedTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteAttendedTransferResponse) GetBody() *CompleteAttendedTransferResponseBody {
	return s.Body
}

func (s *CompleteAttendedTransferResponse) SetHeaders(v map[string]*string) *CompleteAttendedTransferResponse {
	s.Headers = v
	return s
}

func (s *CompleteAttendedTransferResponse) SetStatusCode(v int32) *CompleteAttendedTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteAttendedTransferResponse) SetBody(v *CompleteAttendedTransferResponseBody) *CompleteAttendedTransferResponse {
	s.Body = v
	return s
}

func (s *CompleteAttendedTransferResponse) Validate() error {
	return dara.Validate(s)
}
