// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupID(v string) *GetResourceGroupTotalRequest
	GetResourceGroupID() *string
}

type GetResourceGroupTotalRequest struct {
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s GetResourceGroupTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupTotalRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupTotalRequest) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *GetResourceGroupTotalRequest) SetResourceGroupID(v string) *GetResourceGroupTotalRequest {
	s.ResourceGroupID = &v
	return s
}

func (s *GetResourceGroupTotalRequest) Validate() error {
	return dara.Validate(s)
}
