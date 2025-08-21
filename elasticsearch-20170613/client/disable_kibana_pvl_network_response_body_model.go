// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableKibanaPvlNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableKibanaPvlNetworkResponseBody
	GetRequestId() *string
	SetResult(v bool) *DisableKibanaPvlNetworkResponseBody
	GetResult() *bool
}

type DisableKibanaPvlNetworkResponseBody struct {
	// request id
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1DERFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DisableKibanaPvlNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableKibanaPvlNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DisableKibanaPvlNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableKibanaPvlNetworkResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DisableKibanaPvlNetworkResponseBody) SetRequestId(v string) *DisableKibanaPvlNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableKibanaPvlNetworkResponseBody) SetResult(v bool) *DisableKibanaPvlNetworkResponseBody {
	s.Result = &v
	return s
}

func (s *DisableKibanaPvlNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
