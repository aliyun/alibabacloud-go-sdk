// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeOpEntitiesRequest
	GetEndTime() *int64
	SetEntityObject(v string) *DescribeOpEntitiesRequest
	GetEntityObject() *string
	SetEntityType(v int32) *DescribeOpEntitiesRequest
	GetEntityType() *int32
	SetOpAction(v int32) *DescribeOpEntitiesRequest
	GetOpAction() *int32
	SetPageNo(v int32) *DescribeOpEntitiesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeOpEntitiesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeOpEntitiesRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeOpEntitiesRequest
	GetSourceIp() *string
	SetStartTime(v int64) *DescribeOpEntitiesRequest
	GetStartTime() *int64
}

type DescribeOpEntitiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1536715558000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xx
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	OpAction   *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1534123558000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeOpEntitiesRequest) GetEntityObject() *string {
	return s.EntityObject
}

func (s *DescribeOpEntitiesRequest) GetEntityType() *int32 {
	return s.EntityType
}

func (s *DescribeOpEntitiesRequest) GetOpAction() *int32 {
	return s.OpAction
}

func (s *DescribeOpEntitiesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeOpEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOpEntitiesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeOpEntitiesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOpEntitiesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityObject(v string) *DescribeOpEntitiesRequest {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityType(v int32) *DescribeOpEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOpAction(v int32) *DescribeOpEntitiesRequest {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageNo(v int32) *DescribeOpEntitiesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
