// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketAcl(v string) *PutBucketRequest
	GetBucketAcl() *string
	SetBucketName(v string) *PutBucketRequest
	GetBucketName() *string
	SetComment(v string) *PutBucketRequest
	GetComment() *string
	SetDispatchScope(v string) *PutBucketRequest
	GetDispatchScope() *string
	SetEnsRegionId(v string) *PutBucketRequest
	GetEnsRegionId() *string
	SetLogicalBucketType(v string) *PutBucketRequest
	GetLogicalBucketType() *string
}

type PutBucketRequest struct {
	BucketAcl *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	// This parameter is required.
	BucketName        *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Comment           *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DispatchScope     *string `json:"DispatchScope,omitempty" xml:"DispatchScope,omitempty"`
	EnsRegionId       *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	LogicalBucketType *string `json:"LogicalBucketType,omitempty" xml:"LogicalBucketType,omitempty"`
}

func (s PutBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRequest) GetBucketAcl() *string {
	return s.BucketAcl
}

func (s *PutBucketRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *PutBucketRequest) GetComment() *string {
	return s.Comment
}

func (s *PutBucketRequest) GetDispatchScope() *string {
	return s.DispatchScope
}

func (s *PutBucketRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *PutBucketRequest) GetLogicalBucketType() *string {
	return s.LogicalBucketType
}

func (s *PutBucketRequest) SetBucketAcl(v string) *PutBucketRequest {
	s.BucketAcl = &v
	return s
}

func (s *PutBucketRequest) SetBucketName(v string) *PutBucketRequest {
	s.BucketName = &v
	return s
}

func (s *PutBucketRequest) SetComment(v string) *PutBucketRequest {
	s.Comment = &v
	return s
}

func (s *PutBucketRequest) SetDispatchScope(v string) *PutBucketRequest {
	s.DispatchScope = &v
	return s
}

func (s *PutBucketRequest) SetEnsRegionId(v string) *PutBucketRequest {
	s.EnsRegionId = &v
	return s
}

func (s *PutBucketRequest) SetLogicalBucketType(v string) *PutBucketRequest {
	s.LogicalBucketType = &v
	return s
}

func (s *PutBucketRequest) Validate() error {
	return dara.Validate(s)
}
