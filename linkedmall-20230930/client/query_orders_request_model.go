// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OrderPageQuery) *QueryOrdersRequest
	GetBody() *OrderPageQuery
}

type QueryOrdersRequest struct {
	// This parameter is required.
	Body *OrderPageQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrdersRequest) GoString() string {
	return s.String()
}

func (s *QueryOrdersRequest) GetBody() *OrderPageQuery {
	return s.Body
}

func (s *QueryOrdersRequest) SetBody(v *OrderPageQuery) *QueryOrdersRequest {
	s.Body = v
	return s
}

func (s *QueryOrdersRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
