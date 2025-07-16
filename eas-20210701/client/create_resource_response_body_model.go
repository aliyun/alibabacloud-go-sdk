// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateResourceResponseBody
	GetClusterId() *string
	SetInstanceIds(v []*string) *CreateResourceResponseBody
	GetInstanceIds() []*string
	SetOwnerUid(v string) *CreateResourceResponseBody
	GetOwnerUid() *string
	SetRequestId(v string) *CreateResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *CreateResourceResponseBody
	GetResourceId() *string
	SetResourceName(v string) *CreateResourceResponseBody
	GetResourceName() *string
}

type CreateResourceResponseBody struct {
	// The ID of the cluster to which the resource group belongs.
	//
	// example:
	//
	// cn-shanghai
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The user ID (UID) of the resource group owner.
	//
	// example:
	//
	// 14401087478****
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// eas-r-h7lcw24dyqztwxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// MyResource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateResourceResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateResourceResponseBody) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *CreateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateResourceResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateResourceResponseBody) SetClusterId(v string) *CreateResourceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateResourceResponseBody) SetInstanceIds(v []*string) *CreateResourceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateResourceResponseBody) SetOwnerUid(v string) *CreateResourceResponseBody {
	s.OwnerUid = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceId(v string) *CreateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceName(v string) *CreateResourceResponseBody {
	s.ResourceName = &v
	return s
}

func (s *CreateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
