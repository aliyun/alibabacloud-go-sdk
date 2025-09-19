// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStrategyTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStrategyTargetResponseBody
	GetRequestId() *string
}

type ModifyStrategyTargetResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0E147337-5B0B-5776-B0B6-D569DBA8F60F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStrategyTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStrategyTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStrategyTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStrategyTargetResponseBody) SetRequestId(v string) *ModifyStrategyTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStrategyTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
