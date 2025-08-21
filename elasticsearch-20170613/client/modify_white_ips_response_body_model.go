// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhiteIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWhiteIpsResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyWhiteIpsResponseBody
	GetResult() *bool
}

type ModifyWhiteIpsResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1DERFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyWhiteIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWhiteIpsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyWhiteIpsResponseBody) SetRequestId(v string) *ModifyWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWhiteIpsResponseBody) SetResult(v bool) *ModifyWhiteIpsResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyWhiteIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
