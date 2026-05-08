// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQwenConferenceSgTicketPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicketToken(v string) *QueryQwenConferenceSgTicketPopRequest
	GetTicketToken() *string
}

type QueryQwenConferenceSgTicketPopRequest struct {
	// example:
	//
	// bPbXgB8nSzI9UIbdqAWaOhtr7T3p1Ryr
	TicketToken *string `json:"TicketToken,omitempty" xml:"TicketToken,omitempty"`
}

func (s QueryQwenConferenceSgTicketPopRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketPopRequest) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketPopRequest) GetTicketToken() *string {
	return s.TicketToken
}

func (s *QueryQwenConferenceSgTicketPopRequest) SetTicketToken(v string) *QueryQwenConferenceSgTicketPopRequest {
	s.TicketToken = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopRequest) Validate() error {
	return dara.Validate(s)
}
