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
	// The request ID.
	//
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of resource objects.
	Resources []*ListInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of resources.
	//
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
	// The category of the resource. Valid values:
	//
	// - DataManagement
	//
	// - Engine
	//
	// - Monitor
	//
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The configuration of the resource.
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-13 17:34:52
	GmtCreateAt *string `json:"GmtCreateAt,omitempty" xml:"GmtCreateAt,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-13 17:34:52
	GmtModifiedAt *string `json:"GmtModifiedAt,omitempty" xml:"GmtModifiedAt,omitempty"`
	// The group of the resource.
	//
	// If `Category` is `DataManagement`, valid values are:
	//
	// - storage
	//
	// - modelpipeline
	//
	// - datastorage
	//
	// - modeltrain
	//
	// If `Category` is `Engine`, valid values are:
	//
	// - feature
	//
	// - predict
	//
	// - recall
	//
	// - recengine
	//
	// If `Category` is `Monitor`, valid values are:
	//
	// - logs
	//
	// - logsback
	//
	// - coldstart
	//
	// - deploy
	//
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// - Hologres
	//
	// - EAS
	//
	// - BE
	//
	// - Rec
	//
	// - Platform
	//
	// - SLS
	//
	// - DataHub
	//
	// - ApsaraMQ for Kafka
	//
	// - Realtime Compute for Apache Flink
	//
	// - ACR
	//
	// - OSS
	//
	// - DataWorks
	//
	// - PAI
	//
	// - MaxCompute
	//
	// - Graph Compute Service
	//
	// - ApsaraDB for Redis
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The resource URI.
	//
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
