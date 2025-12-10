// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroup(v *UpdateResourceGroupResponseBodyResourceGroup) *UpdateResourceGroupResponseBody
	GetResourceGroup() *UpdateResourceGroupResponseBodyResourceGroup
}

type UpdateResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource group.
	ResourceGroup *UpdateResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s UpdateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceGroupResponseBody) GetResourceGroup() *UpdateResourceGroupResponseBodyResourceGroup {
	return s.ResourceGroup
}

func (s *UpdateResourceGroupResponseBody) SetRequestId(v string) *UpdateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) SetResourceGroup(v *UpdateResourceGroupResponseBodyResourceGroup) *UpdateResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *UpdateResourceGroupResponseBody) Validate() error {
	if s.ResourceGroup != nil {
		if err := s.ResourceGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// example:
	//
	// 123456789****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18+08:00
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	//
	// example:
	//
	// project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-9gLOoK****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateResourceGroupResponseBodyResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) GetId() *string {
	return s.Id
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) GetName() *string {
	return s.Name
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetId(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetName(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) Validate() error {
	return dara.Validate(s)
}
