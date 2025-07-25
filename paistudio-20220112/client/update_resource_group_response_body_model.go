// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupID(v string) *UpdateResourceGroupResponseBody
	GetResourceGroupID() *string
	SetRequestId(v string) *UpdateResourceGroupResponseBody
	GetRequestId() *string
}

type UpdateResourceGroupResponseBody struct {
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	// example:
	//
	// FFB1D4B4-B253-540A-9B3B-AA711C48A1B7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBody) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *UpdateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceGroupResponseBody) SetResourceGroupID(v string) *UpdateResourceGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) SetRequestId(v string) *UpdateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
