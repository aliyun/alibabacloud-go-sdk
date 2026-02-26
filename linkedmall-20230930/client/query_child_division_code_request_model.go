// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChildDivisionCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DivisionQuery) *QueryChildDivisionCodeRequest
	GetBody() *DivisionQuery
}

type QueryChildDivisionCodeRequest struct {
	// This parameter is required.
	Body *DivisionQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChildDivisionCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryChildDivisionCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeRequest) GetBody() *DivisionQuery {
	return s.Body
}

func (s *QueryChildDivisionCodeRequest) SetBody(v *DivisionQuery) *QueryChildDivisionCodeRequest {
	s.Body = v
	return s
}

func (s *QueryChildDivisionCodeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
