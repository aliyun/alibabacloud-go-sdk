// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQwenConferenceSgTicketSearchPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryQwenConferenceSgTicketSearchPopRequest
	GetKeyword() *string
}

type QueryQwenConferenceSgTicketSearchPopRequest struct {
	// example:
	//
	// neal
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s QueryQwenConferenceSgTicketSearchPopRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketSearchPopRequest) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketSearchPopRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryQwenConferenceSgTicketSearchPopRequest) SetKeyword(v string) *QueryQwenConferenceSgTicketSearchPopRequest {
	s.Keyword = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopRequest) Validate() error {
	return dara.Validate(s)
}
