// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindGuestTicketRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindGuestTicketRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindGuestTicketRecordResponse
	GetStatusCode() *int32
	SetBody(v *FindGuestTicketRecordResponseBody) *FindGuestTicketRecordResponse
	GetBody() *FindGuestTicketRecordResponseBody
}

type FindGuestTicketRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindGuestTicketRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindGuestTicketRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordResponse) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindGuestTicketRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindGuestTicketRecordResponse) GetBody() *FindGuestTicketRecordResponseBody {
	return s.Body
}

func (s *FindGuestTicketRecordResponse) SetHeaders(v map[string]*string) *FindGuestTicketRecordResponse {
	s.Headers = v
	return s
}

func (s *FindGuestTicketRecordResponse) SetStatusCode(v int32) *FindGuestTicketRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *FindGuestTicketRecordResponse) SetBody(v *FindGuestTicketRecordResponseBody) *FindGuestTicketRecordResponse {
	s.Body = v
	return s
}

func (s *FindGuestTicketRecordResponse) Validate() error {
	return dara.Validate(s)
}
