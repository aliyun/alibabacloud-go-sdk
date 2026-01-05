// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductPlanApproversResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApprovers(v []*ListProvisionedProductPlanApproversResponseBodyApprovers) *ListProvisionedProductPlanApproversResponseBody
	GetApprovers() []*ListProvisionedProductPlanApproversResponseBodyApprovers
	SetRequestId(v string) *ListProvisionedProductPlanApproversResponseBody
	GetRequestId() *string
}

type ListProvisionedProductPlanApproversResponseBody struct {
	// An array that consists of reviewers.
	Approvers []*ListProvisionedProductPlanApproversResponseBodyApprovers `json:"Approvers,omitempty" xml:"Approvers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProvisionedProductPlanApproversResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlanApproversResponseBody) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversResponseBody) GetApprovers() []*ListProvisionedProductPlanApproversResponseBodyApprovers {
	return s.Approvers
}

func (s *ListProvisionedProductPlanApproversResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProvisionedProductPlanApproversResponseBody) SetApprovers(v []*ListProvisionedProductPlanApproversResponseBodyApprovers) *ListProvisionedProductPlanApproversResponseBody {
	s.Approvers = v
	return s
}

func (s *ListProvisionedProductPlanApproversResponseBody) SetRequestId(v string) *ListProvisionedProductPlanApproversResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProvisionedProductPlanApproversResponseBody) Validate() error {
	if s.Approvers != nil {
		for _, item := range s.Approvers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProvisionedProductPlanApproversResponseBodyApprovers struct {
	// The name of the reviewer.
	//
	// example:
	//
	// approver-1
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the Resource Access Management (RAM) entity of the reviewer. Valid values:
	//
	// 	- RamUser: a RAM user
	//
	// 	- RamRole: a RAM role
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListProvisionedProductPlanApproversResponseBodyApprovers) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlanApproversResponseBodyApprovers) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversResponseBodyApprovers) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListProvisionedProductPlanApproversResponseBodyApprovers) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListProvisionedProductPlanApproversResponseBodyApprovers) SetPrincipalName(v string) *ListProvisionedProductPlanApproversResponseBodyApprovers {
	s.PrincipalName = &v
	return s
}

func (s *ListProvisionedProductPlanApproversResponseBodyApprovers) SetPrincipalType(v string) *ListProvisionedProductPlanApproversResponseBodyApprovers {
	s.PrincipalType = &v
	return s
}

func (s *ListProvisionedProductPlanApproversResponseBodyApprovers) Validate() error {
	return dara.Validate(s)
}
