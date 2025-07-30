// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateResourceGroupResponseBody
	GetResourceGroupId() *string
}

type CreateResourceGroupResponseBody struct {
	// example:
	//
	// 868B8926-2E7A-5BE7-9897-E811E994****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-ckf3cx7isinhk***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceGroupResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroupId(v string) *CreateResourceGroupResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
