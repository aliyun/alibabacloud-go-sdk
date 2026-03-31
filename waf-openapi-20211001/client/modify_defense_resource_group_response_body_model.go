// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseResourceGroupResponseBody
	GetRequestId() *string
}

type ModifyDefenseResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2CC1AFDE-BB31-5A2F-906E-92FCBDDE6B75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseResourceGroupResponseBody) SetRequestId(v string) *ModifyDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
