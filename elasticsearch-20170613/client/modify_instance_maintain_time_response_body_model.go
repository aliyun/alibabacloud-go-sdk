// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintainTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyInstanceMaintainTimeResponseBody
	GetResult() *bool
}

type ModifyInstanceMaintainTimeResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyInstanceMaintainTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceMaintainTimeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetResult(v bool) *ModifyInstanceMaintainTimeResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyInstanceMaintainTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
