// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucketAcl(v string) *BucketInfo
	GetBucketAcl() *string
	SetBucketName(v string) *BucketInfo
	GetBucketName() *string
	SetComment(v string) *BucketInfo
	GetComment() *string
	SetCreateTime(v string) *BucketInfo
	GetCreateTime() *string
	SetDataRedundancyType(v string) *BucketInfo
	GetDataRedundancyType() *string
	SetDispatcherType(v string) *BucketInfo
	GetDispatcherType() *string
	SetEndpoint(v string) *BucketInfo
	GetEndpoint() *string
	SetEnsRegionId(v string) *BucketInfo
	GetEnsRegionId() *string
	SetModifyTime(v string) *BucketInfo
	GetModifyTime() *string
	SetResourceType(v string) *BucketInfo
	GetResourceType() *string
	SetStorageClass(v string) *BucketInfo
	GetStorageClass() *string
}

type BucketInfo struct {
	// example:
	//
	// private
	BucketAcl *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bucket001
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// this is a bucket
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2011-12-01T12:27:13.000Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	// example:
	//
	// global
	DispatcherType *string `json:"DispatcherType,omitempty" xml:"DispatcherType,omitempty"`
	// example:
	//
	// eos.aliyuncs.com
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// 2011-12-01T12:27:13.000Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// general
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s BucketInfo) String() string {
	return dara.Prettify(s)
}

func (s BucketInfo) GoString() string {
	return s.String()
}

func (s *BucketInfo) GetBucketAcl() *string {
	return s.BucketAcl
}

func (s *BucketInfo) GetBucketName() *string {
	return s.BucketName
}

func (s *BucketInfo) GetComment() *string {
	return s.Comment
}

func (s *BucketInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BucketInfo) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *BucketInfo) GetDispatcherType() *string {
	return s.DispatcherType
}

func (s *BucketInfo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *BucketInfo) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *BucketInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *BucketInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *BucketInfo) GetStorageClass() *string {
	return s.StorageClass
}

func (s *BucketInfo) SetBucketAcl(v string) *BucketInfo {
	s.BucketAcl = &v
	return s
}

func (s *BucketInfo) SetBucketName(v string) *BucketInfo {
	s.BucketName = &v
	return s
}

func (s *BucketInfo) SetComment(v string) *BucketInfo {
	s.Comment = &v
	return s
}

func (s *BucketInfo) SetCreateTime(v string) *BucketInfo {
	s.CreateTime = &v
	return s
}

func (s *BucketInfo) SetDataRedundancyType(v string) *BucketInfo {
	s.DataRedundancyType = &v
	return s
}

func (s *BucketInfo) SetDispatcherType(v string) *BucketInfo {
	s.DispatcherType = &v
	return s
}

func (s *BucketInfo) SetEndpoint(v string) *BucketInfo {
	s.Endpoint = &v
	return s
}

func (s *BucketInfo) SetEnsRegionId(v string) *BucketInfo {
	s.EnsRegionId = &v
	return s
}

func (s *BucketInfo) SetModifyTime(v string) *BucketInfo {
	s.ModifyTime = &v
	return s
}

func (s *BucketInfo) SetResourceType(v string) *BucketInfo {
	s.ResourceType = &v
	return s
}

func (s *BucketInfo) SetStorageClass(v string) *BucketInfo {
	s.StorageClass = &v
	return s
}

func (s *BucketInfo) Validate() error {
	return dara.Validate(s)
}
