// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrafficMirrorSessionResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateTrafficMirrorSessionResponseBody
	GetResourceGroupId() *string
	SetTrafficMirrorSessionId(v string) *CreateTrafficMirrorSessionResponseBody
	GetTrafficMirrorSessionId() *string
}

type CreateTrafficMirrorSessionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the mirrored traffic belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the traffic mirror session.
	//
	// example:
	//
	// tms-j6ce5di4w7nvigfjz****
	TrafficMirrorSessionId *string `json:"TrafficMirrorSessionId,omitempty" xml:"TrafficMirrorSessionId,omitempty"`
}

func (s CreateTrafficMirrorSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrafficMirrorSessionResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTrafficMirrorSessionResponseBody) GetTrafficMirrorSessionId() *string {
	return s.TrafficMirrorSessionId
}

func (s *CreateTrafficMirrorSessionResponseBody) SetRequestId(v string) *CreateTrafficMirrorSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficMirrorSessionResponseBody) SetResourceGroupId(v string) *CreateTrafficMirrorSessionResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTrafficMirrorSessionResponseBody) SetTrafficMirrorSessionId(v string) *CreateTrafficMirrorSessionResponseBody {
	s.TrafficMirrorSessionId = &v
	return s
}

func (s *CreateTrafficMirrorSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
