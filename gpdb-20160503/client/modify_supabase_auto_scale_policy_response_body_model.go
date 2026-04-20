// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseAutoScalePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySupabaseAutoScalePolicyResponseBody
	GetRequestId() *string
}

type ModifySupabaseAutoScalePolicyResponseBody struct {
	// example:
	//
	// 07F6177E-6DE4-408A-BB4F-0723301340F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySupabaseAutoScalePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseAutoScalePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySupabaseAutoScalePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySupabaseAutoScalePolicyResponseBody) SetRequestId(v string) *ModifySupabaseAutoScalePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySupabaseAutoScalePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
