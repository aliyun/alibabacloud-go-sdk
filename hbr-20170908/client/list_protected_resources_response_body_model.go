// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProtectedResourcesResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListProtectedResourcesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListProtectedResourcesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListProtectedResourcesResponseBody
	GetNextToken() *string
	SetProtectedResources(v []*ListProtectedResourcesResponseBodyProtectedResources) *ListProtectedResourcesResponseBody
	GetProtectedResources() []*ListProtectedResourcesResponseBodyProtectedResources
	SetRequestId(v string) *ListProtectedResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProtectedResourcesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListProtectedResourcesResponseBody
	GetTotalCount() *int32
}

type ListProtectedResourcesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// eyJJ************MX0=
	NextToken          *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtectedResources []*ListProtectedResourcesResponseBodyProtectedResources `json:"ProtectedResources,omitempty" xml:"ProtectedResources,omitempty" type:"Repeated"`
	// example:
	//
	// EB09****-****-****-****-********6C38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProtectedResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProtectedResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProtectedResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectedResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProtectedResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectedResourcesResponseBody) GetProtectedResources() []*ListProtectedResourcesResponseBodyProtectedResources {
	return s.ProtectedResources
}

func (s *ListProtectedResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProtectedResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProtectedResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProtectedResourcesResponseBody) SetCode(v string) *ListProtectedResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetMaxResults(v int32) *ListProtectedResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetMessage(v string) *ListProtectedResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetNextToken(v string) *ListProtectedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetProtectedResources(v []*ListProtectedResourcesResponseBodyProtectedResources) *ListProtectedResourcesResponseBody {
	s.ProtectedResources = v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetRequestId(v string) *ListProtectedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetSuccess(v bool) *ListProtectedResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetTotalCount(v int32) *ListProtectedResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) Validate() error {
	if s.ProtectedResources != nil {
		for _, item := range s.ProtectedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProtectedResourcesResponseBodyProtectedResources struct {
	// example:
	//
	// BASIC
	CreatedByProduct *string `json:"CreatedByProduct,omitempty" xml:"CreatedByProduct,omitempty"`
	// example:
	//
	// 107374182400
	ProtectedDataSize *int64 `json:"ProtectedDataSize,omitempty" xml:"ProtectedDataSize,omitempty"`
	// example:
	//
	// pr-0004************gs61
	ProtectedResourceId *string `json:"ProtectedResourceId,omitempty" xml:"ProtectedResourceId,omitempty"`
	// example:
	//
	// i-wz95************7zrd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1024********0703
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 30
	SnapshotCount *int64 `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListProtectedResourcesResponseBodyProtectedResources) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedResourcesResponseBodyProtectedResources) GoString() string {
	return s.String()
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetCreatedByProduct() *string {
	return s.CreatedByProduct
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetProtectedDataSize() *int64 {
	return s.ProtectedDataSize
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetProtectedResourceId() *string {
	return s.ProtectedResourceId
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetSnapshotCount() *int64 {
	return s.SnapshotCount
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetSourceType() *string {
	return s.SourceType
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetCreatedByProduct(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.CreatedByProduct = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetProtectedDataSize(v int64) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ProtectedDataSize = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetProtectedResourceId(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ProtectedResourceId = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetResourceId(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ResourceId = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetResourceOwnerId(v int64) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetSnapshotCount(v int64) *ListProtectedResourcesResponseBodyProtectedResources {
	s.SnapshotCount = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetSourceType(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.SourceType = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) Validate() error {
	return dara.Validate(s)
}
