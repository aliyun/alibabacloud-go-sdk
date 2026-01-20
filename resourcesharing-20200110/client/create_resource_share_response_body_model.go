// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceShareResponseBody
	GetRequestId() *string
	SetResourceShare(v *CreateResourceShareResponseBodyResourceShare) *CreateResourceShareResponseBody
	GetResourceShare() *CreateResourceShareResponseBodyResourceShare
}

type CreateResourceShareResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2C3FA051-61DC-4F3E-81E9-E4830524DF4B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource share.
	ResourceShare *CreateResourceShareResponseBodyResourceShare `json:"ResourceShare,omitempty" xml:"ResourceShare,omitempty" type:"Struct"`
}

func (s CreateResourceShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceShareResponseBody) GetResourceShare() *CreateResourceShareResponseBodyResourceShare {
	return s.ResourceShare
}

func (s *CreateResourceShareResponseBody) SetRequestId(v string) *CreateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceShareResponseBody) SetResourceShare(v *CreateResourceShareResponseBodyResourceShare) *CreateResourceShareResponseBody {
	s.ResourceShare = v
	return s
}

func (s *CreateResourceShareResponseBody) Validate() error {
	if s.ResourceShare != nil {
		if err := s.ResourceShare.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourceShareResponseBodyResourceShare struct {
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
	// test
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
	// >  The system automatically deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	//
	// example:
	//
	// Active
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	// The time when the resource share was updated.
	//
	// example:
	//
	// 2020-12-03T08:02:22.413Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateResourceShareResponseBodyResourceShare) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareResponseBodyResourceShare) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponseBodyResourceShare) GetAllowExternalTargets() *bool {
	return s.AllowExternalTargets
}

func (s *CreateResourceShareResponseBodyResourceShare) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateResourceShareResponseBodyResourceShare) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *CreateResourceShareResponseBodyResourceShare) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *CreateResourceShareResponseBodyResourceShare) GetResourceShareOwner() *string {
	return s.ResourceShareOwner
}

func (s *CreateResourceShareResponseBodyResourceShare) GetResourceShareStatus() *string {
	return s.ResourceShareStatus
}

func (s *CreateResourceShareResponseBodyResourceShare) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateResourceShareResponseBodyResourceShare) SetAllowExternalTargets(v bool) *CreateResourceShareResponseBodyResourceShare {
	s.AllowExternalTargets = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetCreateTime(v string) *CreateResourceShareResponseBodyResourceShare {
	s.CreateTime = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareId(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareId = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareName(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareName = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareOwner(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareOwner = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareStatus(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareStatus = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetUpdateTime(v string) *CreateResourceShareResponseBodyResourceShare {
	s.UpdateTime = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) Validate() error {
	return dara.Validate(s)
}
