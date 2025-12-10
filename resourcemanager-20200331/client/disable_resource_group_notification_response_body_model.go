// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableResourceGroupNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableResourceGroupNotificationResponseBody
	GetRequestId() *string
}

type DisableResourceGroupNotificationResponseBody struct {
	// example:
	//
	// 898FAB24-7509-43EE-A287-086FE4C44394
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableResourceGroupNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableResourceGroupNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableResourceGroupNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableResourceGroupNotificationResponseBody) SetRequestId(v string) *DisableResourceGroupNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableResourceGroupNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}
