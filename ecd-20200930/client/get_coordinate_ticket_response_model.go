// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoordinateTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCoordinateTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCoordinateTicketResponse
	GetStatusCode() *int32
	SetBody(v *GetCoordinateTicketResponseBody) *GetCoordinateTicketResponse
	GetBody() *GetCoordinateTicketResponseBody
}

type GetCoordinateTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCoordinateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCoordinateTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCoordinateTicketResponse) GoString() string {
	return s.String()
}

func (s *GetCoordinateTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCoordinateTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCoordinateTicketResponse) GetBody() *GetCoordinateTicketResponseBody {
	return s.Body
}

func (s *GetCoordinateTicketResponse) SetHeaders(v map[string]*string) *GetCoordinateTicketResponse {
	s.Headers = v
	return s
}

func (s *GetCoordinateTicketResponse) SetStatusCode(v int32) *GetCoordinateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCoordinateTicketResponse) SetBody(v *GetCoordinateTicketResponseBody) *GetCoordinateTicketResponse {
	s.Body = v
	return s
}

func (s *GetCoordinateTicketResponse) Validate() error {
	return dara.Validate(s)
}
