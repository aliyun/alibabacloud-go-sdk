// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceVpcAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCInstanceVpcAttributeResponseBody
	GetRequestId() *string
}

type ModifyRCInstanceVpcAttributeResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCInstanceVpcAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceVpcAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceVpcAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCInstanceVpcAttributeResponseBody) SetRequestId(v string) *ModifyRCInstanceVpcAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCInstanceVpcAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
