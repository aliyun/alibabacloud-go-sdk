// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAppInstanceTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshAppInstanceTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshAppInstanceTicketResponse
	GetStatusCode() *int32
	SetBody(v *RefreshAppInstanceTicketResponseBody) *RefreshAppInstanceTicketResponse
	GetBody() *RefreshAppInstanceTicketResponseBody
}

type RefreshAppInstanceTicketResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAppInstanceTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAppInstanceTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshAppInstanceTicketResponse) GoString() string {
	return s.String()
}

func (s *RefreshAppInstanceTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshAppInstanceTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshAppInstanceTicketResponse) GetBody() *RefreshAppInstanceTicketResponseBody {
	return s.Body
}

func (s *RefreshAppInstanceTicketResponse) SetHeaders(v map[string]*string) *RefreshAppInstanceTicketResponse {
	s.Headers = v
	return s
}

func (s *RefreshAppInstanceTicketResponse) SetStatusCode(v int32) *RefreshAppInstanceTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAppInstanceTicketResponse) SetBody(v *RefreshAppInstanceTicketResponseBody) *RefreshAppInstanceTicketResponse {
	s.Body = v
	return s
}

func (s *RefreshAppInstanceTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
