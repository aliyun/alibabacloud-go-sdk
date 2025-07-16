// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketMemoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTicketMemoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTicketMemoResponse
	GetStatusCode() *int32
	SetBody(v *AddTicketMemoResponseBody) *AddTicketMemoResponse
	GetBody() *AddTicketMemoResponseBody
}

type AddTicketMemoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTicketMemoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTicketMemoResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoResponse) GoString() string {
	return s.String()
}

func (s *AddTicketMemoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTicketMemoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTicketMemoResponse) GetBody() *AddTicketMemoResponseBody {
	return s.Body
}

func (s *AddTicketMemoResponse) SetHeaders(v map[string]*string) *AddTicketMemoResponse {
	s.Headers = v
	return s
}

func (s *AddTicketMemoResponse) SetStatusCode(v int32) *AddTicketMemoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTicketMemoResponse) SetBody(v *AddTicketMemoResponseBody) *AddTicketMemoResponse {
	s.Body = v
	return s
}

func (s *AddTicketMemoResponse) Validate() error {
	return dara.Validate(s)
}
