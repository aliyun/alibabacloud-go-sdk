// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBucketScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *OperateBucketScanTaskRequest
	GetBucketName() *string
	SetOperateCode(v int32) *OperateBucketScanTaskRequest
	GetOperateCode() *int32
}

type OperateBucketScanTaskRequest struct {
	// example:
	//
	// testBucket******
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 1
	OperateCode *int32 `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
}

func (s OperateBucketScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateBucketScanTaskRequest) GoString() string {
	return s.String()
}

func (s *OperateBucketScanTaskRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *OperateBucketScanTaskRequest) GetOperateCode() *int32 {
	return s.OperateCode
}

func (s *OperateBucketScanTaskRequest) SetBucketName(v string) *OperateBucketScanTaskRequest {
	s.BucketName = &v
	return s
}

func (s *OperateBucketScanTaskRequest) SetOperateCode(v int32) *OperateBucketScanTaskRequest {
	s.OperateCode = &v
	return s
}

func (s *OperateBucketScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
