// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceShareResponseBody
	GetRequestId() *string
	SetResourceShare(v *UpdateResourceShareResponseBodyResourceShare) *UpdateResourceShareResponseBody
	GetResourceShare() *UpdateResourceShareResponseBodyResourceShare
}

type UpdateResourceShareResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2860A3A4-D8C1-4EF4-954E-84A3945E26E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource share.
	ResourceShare *UpdateResourceShareResponseBodyResourceShare `json:"ResourceShare,omitempty" xml:"ResourceShare,omitempty" type:"Struct"`
}

func (s UpdateResourceShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceShareResponseBody) GetResourceShare() *UpdateResourceShareResponseBodyResourceShare {
	return s.ResourceShare
}

func (s *UpdateResourceShareResponseBody) SetRequestId(v string) *UpdateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceShareResponseBody) SetResourceShare(v *UpdateResourceShareResponseBodyResourceShare) *UpdateResourceShareResponseBody {
	s.ResourceShare = v
	return s
}

func (s *UpdateResourceShareResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateResourceShareResponseBodyResourceShare struct {
	// Indicates whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// 	- false: Resources in the resource share can be shared only with accounts in the resource directory.
	//
	// 	- true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	//
	// example:
	//
	// false
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The time when the resource share was created.
	//
	// example:
	//
	// 2020-12-03T08:02:22.413Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource share.
	//
	// example:
	//
	// rs-qSkW1HBY****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	//
	// example:
	//
	// new
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The owner of the resource share.
	//
	// example:
	//
	// 151266687691****
	ResourceShareOwner *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	// The status of the resource share. Valid values:
	//
	// 	- Active: The resource share is enabled.
	//
	// 	- Pending: The resource share is associated with one or more resource sharing invitations that are waiting for confirmation.
	//
	// 	- Deleting: The resource share is being deleted.
	//
	// 	- Deleted: The resource share is deleted.
	//
	// >  The system deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	//
	// example:
	//
	// Active
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	// The time when the resource share was updated.
	//
	// example:
	//
	// 2020-12-04T08:55:25.382Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateResourceShareResponseBodyResourceShare) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceShareResponseBodyResourceShare) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponseBodyResourceShare) GetAllowExternalTargets() *bool {
	return s.AllowExternalTargets
}

func (s *UpdateResourceShareResponseBodyResourceShare) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateResourceShareResponseBodyResourceShare) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *UpdateResourceShareResponseBodyResourceShare) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *UpdateResourceShareResponseBodyResourceShare) GetResourceShareOwner() *string {
	return s.ResourceShareOwner
}

func (s *UpdateResourceShareResponseBodyResourceShare) GetResourceShareStatus() *string {
	return s.ResourceShareStatus
}

func (s *UpdateResourceShareResponseBodyResourceShare) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetAllowExternalTargets(v bool) *UpdateResourceShareResponseBodyResourceShare {
	s.AllowExternalTargets = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetCreateTime(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.CreateTime = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareId(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareId = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareName(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareName = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareOwner(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareOwner = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareStatus(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareStatus = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetUpdateTime(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.UpdateTime = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) Validate() error {
	return dara.Validate(s)
}
