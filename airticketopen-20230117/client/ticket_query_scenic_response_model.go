// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryScenicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketQueryScenicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketQueryScenicResponse
	GetStatusCode() *int32
	SetBody(v *TicketQueryScenicResponseBody) *TicketQueryScenicResponse
	GetBody() *TicketQueryScenicResponseBody
}

type TicketQueryScenicResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketQueryScenicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketQueryScenicResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryScenicResponse) GoString() string {
	return s.String()
}

func (s *TicketQueryScenicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketQueryScenicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketQueryScenicResponse) GetBody() *TicketQueryScenicResponseBody {
	return s.Body
}

func (s *TicketQueryScenicResponse) SetHeaders(v map[string]*string) *TicketQueryScenicResponse {
	s.Headers = v
	return s
}

func (s *TicketQueryScenicResponse) SetStatusCode(v int32) *TicketQueryScenicResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketQueryScenicResponse) SetBody(v *TicketQueryScenicResponseBody) *TicketQueryScenicResponse {
	s.Body = v
	return s
}

func (s *TicketQueryScenicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
