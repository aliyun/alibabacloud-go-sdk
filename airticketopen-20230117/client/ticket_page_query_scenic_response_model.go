// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPageQueryScenicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketPageQueryScenicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketPageQueryScenicResponse
	GetStatusCode() *int32
	SetBody(v *TicketPageQueryScenicResponseBody) *TicketPageQueryScenicResponse
	GetBody() *TicketPageQueryScenicResponseBody
}

type TicketPageQueryScenicResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketPageQueryScenicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketPageQueryScenicResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryScenicResponse) GoString() string {
	return s.String()
}

func (s *TicketPageQueryScenicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketPageQueryScenicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketPageQueryScenicResponse) GetBody() *TicketPageQueryScenicResponseBody {
	return s.Body
}

func (s *TicketPageQueryScenicResponse) SetHeaders(v map[string]*string) *TicketPageQueryScenicResponse {
	s.Headers = v
	return s
}

func (s *TicketPageQueryScenicResponse) SetStatusCode(v int32) *TicketPageQueryScenicResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketPageQueryScenicResponse) SetBody(v *TicketPageQueryScenicResponseBody) *TicketPageQueryScenicResponse {
	s.Body = v
	return s
}

func (s *TicketPageQueryScenicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
