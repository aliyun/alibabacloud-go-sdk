// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQosId(v string) *CreateQosResponseBody
	GetQosId() *string
	SetRequestId(v string) *CreateQosResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateQosResponseBody
	GetResourceGroupId() *string
}

type CreateQosResponseBody struct {
	// The ID of the QoS policy.
	//
	// example:
	//
	// rg-acfm2iu4fnc****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AACF5140-783D-48F0-9E4F-E59D716F7D08
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the QoS policy belongs.
	//
	// example:
	//
	// qos-oek3r2cmvk7m8q****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQosResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQosResponseBody) GetQosId() *string {
	return s.QosId
}

func (s *CreateQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQosResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateQosResponseBody) SetQosId(v string) *CreateQosResponseBody {
	s.QosId = &v
	return s
}

func (s *CreateQosResponseBody) SetRequestId(v string) *CreateQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQosResponseBody) SetResourceGroupId(v string) *CreateQosResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateQosResponseBody) Validate() error {
	return dara.Validate(s)
}
