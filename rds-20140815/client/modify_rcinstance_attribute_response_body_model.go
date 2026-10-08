// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCInstanceAttributeResponseBody
	GetRequestId() *string
}

type ModifyRCInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCInstanceAttributeResponseBody) SetRequestId(v string) *ModifyRCInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
