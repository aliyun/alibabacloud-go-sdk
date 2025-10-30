// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollections(v *ListCollectionsResponseBodyCollections) *ListCollectionsResponseBody
	GetCollections() *ListCollectionsResponseBodyCollections
	SetCount(v int32) *ListCollectionsResponseBody
	GetCount() *int32
	SetDBInstanceId(v string) *ListCollectionsResponseBody
	GetDBInstanceId() *string
	SetMessage(v string) *ListCollectionsResponseBody
	GetMessage() *string
	SetNamespace(v string) *ListCollectionsResponseBody
	GetNamespace() *string
	SetRegionId(v string) *ListCollectionsResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ListCollectionsResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListCollectionsResponseBody
	GetStatus() *string
}

type ListCollectionsResponseBody struct {
	// The queried vector collections.
	Collections *ListCollectionsResponseBodyCollections `json:"Collections,omitempty" xml:"Collections,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCollectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBody) GetCollections() *ListCollectionsResponseBodyCollections {
	return s.Collections
}

func (s *ListCollectionsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListCollectionsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListCollectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCollectionsResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *ListCollectionsResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCollectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCollectionsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListCollectionsResponseBody) SetCollections(v *ListCollectionsResponseBodyCollections) *ListCollectionsResponseBody {
	s.Collections = v
	return s
}

func (s *ListCollectionsResponseBody) SetCount(v int32) *ListCollectionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListCollectionsResponseBody) SetDBInstanceId(v string) *ListCollectionsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetMessage(v string) *ListCollectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCollectionsResponseBody) SetNamespace(v string) *ListCollectionsResponseBody {
	s.Namespace = &v
	return s
}

func (s *ListCollectionsResponseBody) SetRegionId(v string) *ListCollectionsResponseBody {
	s.RegionId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetRequestId(v string) *ListCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetStatus(v string) *ListCollectionsResponseBody {
	s.Status = &v
	return s
}

func (s *ListCollectionsResponseBody) Validate() error {
	if s.Collections != nil {
		if err := s.Collections.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCollectionsResponseBodyCollections struct {
	Collection []*string `json:"Collection,omitempty" xml:"Collection,omitempty" type:"Repeated"`
}

func (s ListCollectionsResponseBodyCollections) String() string {
	return dara.Prettify(s)
}

func (s ListCollectionsResponseBodyCollections) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBodyCollections) GetCollection() []*string {
	return s.Collection
}

func (s *ListCollectionsResponseBodyCollections) SetCollection(v []*string) *ListCollectionsResponseBodyCollections {
	s.Collection = v
	return s
}

func (s *ListCollectionsResponseBodyCollections) Validate() error {
	return dara.Validate(s)
}
