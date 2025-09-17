// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWorkgroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWorkgroupAttributeResponseBody
	GetRequestId() *string
}

type ModifyWorkgroupAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3E8B9ABB-289A-44E6-942D-8AA9E493****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWorkgroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWorkgroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWorkgroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWorkgroupAttributeResponseBody) SetRequestId(v string) *ModifyWorkgroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWorkgroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
