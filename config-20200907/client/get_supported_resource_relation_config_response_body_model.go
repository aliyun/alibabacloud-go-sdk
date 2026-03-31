// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupportedResourceRelationConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSupportedResourceRelationConfigResponseBody
	GetRequestId() *string
	SetResourceRelationConfigList(v []*GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) *GetSupportedResourceRelationConfigResponseBody
	GetResourceRelationConfigList() []*GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList
}

type GetSupportedResourceRelationConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 409D022F-394C-5AAB-A74A-2F1DC9F6375E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that contains the relationships.
	ResourceRelationConfigList []*GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList `json:"ResourceRelationConfigList,omitempty" xml:"ResourceRelationConfigList,omitempty" type:"Repeated"`
}

func (s GetSupportedResourceRelationConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedResourceRelationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupportedResourceRelationConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupportedResourceRelationConfigResponseBody) GetResourceRelationConfigList() []*GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList {
	return s.ResourceRelationConfigList
}

func (s *GetSupportedResourceRelationConfigResponseBody) SetRequestId(v string) *GetSupportedResourceRelationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupportedResourceRelationConfigResponseBody) SetResourceRelationConfigList(v []*GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) *GetSupportedResourceRelationConfigResponseBody {
	s.ResourceRelationConfigList = v
	return s
}

func (s *GetSupportedResourceRelationConfigResponseBody) Validate() error {
	if s.ResourceRelationConfigList != nil {
		for _, item := range s.ResourceRelationConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList struct {
	// The type of the relationship between the resource and the object. Valid values:
	//
	// 	- IsContained: The object is included as part of the resource.
	//
	// 	- IsAttachedTo: The object is added to the resource.
	//
	// 	- IsAssociatedIn: The object is associated with the resource.
	//
	// 	- Contains: The actual value contains the expected value.
	//
	// example:
	//
	// IsAttachedTo
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Disk
	TargetResourceType *string `json:"TargetResourceType,omitempty" xml:"TargetResourceType,omitempty"`
}

func (s GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) GoString() string {
	return s.String()
}

func (s *GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) GetRelationType() *string {
	return s.RelationType
}

func (s *GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) GetTargetResourceType() *string {
	return s.TargetResourceType
}

func (s *GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) SetRelationType(v string) *GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList {
	s.RelationType = &v
	return s
}

func (s *GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) SetTargetResourceType(v string) *GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList {
	s.TargetResourceType = &v
	return s
}

func (s *GetSupportedResourceRelationConfigResponseBodyResourceRelationConfigList) Validate() error {
	return dara.Validate(s)
}
