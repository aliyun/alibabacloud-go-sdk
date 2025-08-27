// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketChangingDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketChangingDetailResponse
	GetStatusCode() *int32
	SetBody(v *TicketChangingDetailResponseBody) *TicketChangingDetailResponse
	GetBody() *TicketChangingDetailResponseBody
}

type TicketChangingDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketChangingDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketChangingDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingDetailResponse) GoString() string {
	return s.String()
}

func (s *TicketChangingDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketChangingDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketChangingDetailResponse) GetBody() *TicketChangingDetailResponseBody {
	return s.Body
}

func (s *TicketChangingDetailResponse) SetHeaders(v map[string]*string) *TicketChangingDetailResponse {
	s.Headers = v
	return s
}

func (s *TicketChangingDetailResponse) SetStatusCode(v int32) *TicketChangingDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketChangingDetailResponse) SetBody(v *TicketChangingDetailResponseBody) *TicketChangingDetailResponse {
	s.Body = v
	return s
}

func (s *TicketChangingDetailResponse) Validate() error {
	return dara.Validate(s)
}
