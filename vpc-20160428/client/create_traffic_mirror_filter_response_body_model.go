// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrafficMirrorFilterResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateTrafficMirrorFilterResponseBody
	GetResourceGroupId() *string
	SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorFilterResponseBody
	GetTrafficMirrorFilterId() *string
}

type CreateTrafficMirrorFilterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 28DB147D-217B-43E8-9E94-A3F6837DDC8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the mirrored traffic belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the filter.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
}

func (s CreateTrafficMirrorFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrafficMirrorFilterResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTrafficMirrorFilterResponseBody) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *CreateTrafficMirrorFilterResponseBody) SetRequestId(v string) *CreateTrafficMirrorFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficMirrorFilterResponseBody) SetResourceGroupId(v string) *CreateTrafficMirrorFilterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTrafficMirrorFilterResponseBody) SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorFilterResponseBody {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *CreateTrafficMirrorFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
