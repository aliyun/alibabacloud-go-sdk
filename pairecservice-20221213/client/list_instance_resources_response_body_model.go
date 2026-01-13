// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstanceResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*ListInstanceResourcesResponseBodyResources) *ListInstanceResourcesResponseBody
	GetResources() []*ListInstanceResourcesResponseBodyResources
	SetTotalCount(v int64) *ListInstanceResourcesResponseBody
	GetTotalCount() *int64
}

type ListInstanceResourcesResponseBody struct {
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*ListInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResourcesResponseBody) GetResources() []*ListInstanceResourcesResponseBodyResources {
	return s.Resources
}

func (s *ListInstanceResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceResourcesResponseBody) SetRequestId(v string) *ListInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResourcesResponseBody) SetResources(v []*ListInstanceResourcesResponseBodyResources) *ListInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListInstanceResourcesResponseBody) SetTotalCount(v int64) *ListInstanceResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResourcesResponseBodyResources struct {
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtCreateAt *string `json:"GmtCreateAt,omitempty" xml:"GmtCreateAt,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtModifiedAt *string `json:"GmtModifiedAt,omitempty" xml:"GmtModifiedAt,omitempty"`
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListInstanceResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesResponseBodyResources) GetCategory() *string {
	return s.Category
}

func (s *ListInstanceResourcesResponseBodyResources) GetConfig() *string {
	return s.Config
}

func (s *ListInstanceResourcesResponseBodyResources) GetGmtCreateAt() *string {
	return s.GmtCreateAt
}

func (s *ListInstanceResourcesResponseBodyResources) GetGmtModifiedAt() *string {
	return s.GmtModifiedAt
}

func (s *ListInstanceResourcesResponseBodyResources) GetGroup() *string {
	return s.Group
}

func (s *ListInstanceResourcesResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListInstanceResourcesResponseBodyResources) GetType() *string {
	return s.Type
}

func (s *ListInstanceResourcesResponseBodyResources) GetUri() *string {
	return s.Uri
}

func (s *ListInstanceResourcesResponseBodyResources) SetCategory(v string) *ListInstanceResourcesResponseBodyResources {
	s.Category = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetConfig(v string) *ListInstanceResourcesResponseBodyResources {
	s.Config = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetGmtCreateAt(v string) *ListInstanceResourcesResponseBodyResources {
	s.GmtCreateAt = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetGmtModifiedAt(v string) *ListInstanceResourcesResponseBodyResources {
	s.GmtModifiedAt = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetGroup(v string) *ListInstanceResourcesResponseBodyResources {
	s.Group = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetResourceId(v string) *ListInstanceResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetType(v string) *ListInstanceResourcesResponseBodyResources {
	s.Type = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetUri(v string) *ListInstanceResourcesResponseBodyResources {
	s.Uri = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
