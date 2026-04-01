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
	// The access control list (ACL) of the bucket. private public-read-write public-read
	//
	// example:
	//
	// private
	BucketAcl *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	// The name of the bucket. This parameter can contain 3 to 50 characters in length and can contain only lowercase letters, digits, and hyphens (-). The name cannot start or end with a hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The description. The description can be 0 to 128 characters in length, and can contain Chinese characters and emoticons.
	//
	// example:
	//
	// numb
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The scheduling scope. This parameter takes effect only when the value of the LogicalBucketType parameter is standard. Valid values:
	//
	// 	- **domestic**: Chinese mainland.
	//
	// 	- **oversea**: outside the Chinese mainland.
	//
	// example:
	//
	// domestic
	DispatchScope *string `json:"DispatchScope,omitempty" xml:"DispatchScope,omitempty"`
	// The ID of the region where the node is located. If this parameter is not specified, the node is the global default node.
	//
	// example:
	//
	// cn-beijing-15
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The type of the logical bucket. Valid values: sink standard (default)
	//
	// example:
	//
	// sink
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
