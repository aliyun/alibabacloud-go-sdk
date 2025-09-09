// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstancePublicAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableInstancePublicAccessResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DisableInstancePublicAccessResponseBody
	GetRequestId() *string
}

type DisableInstancePublicAccessResponseBody struct {
	// The ID of the bastion host whose Internet access is disabled.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 76FAAF15-D3A3-4099-9941-FC408D9FDB4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInstancePublicAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableInstancePublicAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableInstancePublicAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableInstancePublicAccessResponseBody) SetInstanceId(v string) *DisableInstancePublicAccessResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DisableInstancePublicAccessResponseBody) SetRequestId(v string) *DisableInstancePublicAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInstancePublicAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
