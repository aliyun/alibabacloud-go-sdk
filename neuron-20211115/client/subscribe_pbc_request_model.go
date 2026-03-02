// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribePbcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcSubscribe) *SubscribePbcRequest
	GetBody() *PbcSubscribe
}

type SubscribePbcRequest struct {
	Body *PbcSubscribe `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribePbcRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribePbcRequest) GoString() string {
	return s.String()
}

func (s *SubscribePbcRequest) GetBody() *PbcSubscribe {
	return s.Body
}

func (s *SubscribePbcRequest) SetBody(v *PbcSubscribe) *SubscribePbcRequest {
	s.Body = v
	return s
}

func (s *SubscribePbcRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
