// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPodStatus(v string) *GetResourceGroupRequestRequest
	GetPodStatus() *string
	SetResourceGroupID(v string) *GetResourceGroupRequestRequest
	GetResourceGroupID() *string
}

type GetResourceGroupRequestRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// Running
	PodStatus *string `json:"PodStatus,omitempty" xml:"PodStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s GetResourceGroupRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupRequestRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestRequest) GetPodStatus() *string {
	return s.PodStatus
}

func (s *GetResourceGroupRequestRequest) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *GetResourceGroupRequestRequest) SetPodStatus(v string) *GetResourceGroupRequestRequest {
	s.PodStatus = &v
	return s
}

func (s *GetResourceGroupRequestRequest) SetResourceGroupID(v string) *GetResourceGroupRequestRequest {
	s.ResourceGroupID = &v
	return s
}

func (s *GetResourceGroupRequestRequest) Validate() error {
	return dara.Validate(s)
}
