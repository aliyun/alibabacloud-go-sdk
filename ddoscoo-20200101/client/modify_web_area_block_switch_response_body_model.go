// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAreaBlockSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebAreaBlockSwitchResponseBody
	GetRequestId() *string
}

type ModifyWebAreaBlockSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAreaBlockSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAreaBlockSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebAreaBlockSwitchResponseBody) SetRequestId(v string) *ModifyWebAreaBlockSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
