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
	SetResourceGroupID(v string) *CreateResourceGroupResponseBody
	GetResourceGroupID() *string
}

type CreateResourceGroupResponseBody struct {
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ResourceGroup ID。
	//
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
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

func (s *CreateResourceGroupResponseBody) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroupID(v string) *CreateResourceGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *CreateResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
