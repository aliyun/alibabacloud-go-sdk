// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribePbcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubscribePbcResponseBody
	GetRequestId() *string
	SetResult(v []*Pbc) *SubscribePbcResponseBody
	GetResult() []*Pbc
}

type SubscribePbcResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*Pbc  `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s SubscribePbcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubscribePbcResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribePbcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubscribePbcResponseBody) GetResult() []*Pbc {
	return s.Result
}

func (s *SubscribePbcResponseBody) SetRequestId(v string) *SubscribePbcResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribePbcResponseBody) SetResult(v []*Pbc) *SubscribePbcResponseBody {
	s.Result = v
	return s
}

func (s *SubscribePbcResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
