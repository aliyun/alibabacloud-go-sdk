// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscribePbcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSubscribePbcsResponseBody
	GetRequestId() *string
	SetResult(v []*Pbc) *ListSubscribePbcsResponseBody
	GetResult() []*Pbc
}

type ListSubscribePbcsResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*Pbc  `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSubscribePbcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubscribePbcsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscribePbcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubscribePbcsResponseBody) GetResult() []*Pbc {
	return s.Result
}

func (s *ListSubscribePbcsResponseBody) SetRequestId(v string) *ListSubscribePbcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscribePbcsResponseBody) SetResult(v []*Pbc) *ListSubscribePbcsResponseBody {
	s.Result = v
	return s
}

func (s *ListSubscribePbcsResponseBody) Validate() error {
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
