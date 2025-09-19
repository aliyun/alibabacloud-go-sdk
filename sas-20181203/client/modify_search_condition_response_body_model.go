// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearchConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySearchConditionResponseBody
	GetRequestId() *string
}

type ModifySearchConditionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DFFCF83-4F7B-5E05-B82D-3B619D5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySearchConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySearchConditionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySearchConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySearchConditionResponseBody) SetRequestId(v string) *ModifySearchConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySearchConditionResponseBody) Validate() error {
	return dara.Validate(s)
}
