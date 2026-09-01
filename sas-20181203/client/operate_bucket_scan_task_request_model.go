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
	SetSource(v string) *OperateBucketScanTaskRequest
	GetSource() *string
}

type OperateBucketScanTaskRequest struct {
	// The bucket name.
	//
	// example:
	//
	// iboxpublic****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The operation to perform on the bucket. Valid values:
	//
	// - **1**: Cancel detection.
	//
	// example:
	//
	// 1
	OperateCode *int32 `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
	// The business source. Valid values:
	//
	// - **OSS**: OSS
	//
	// - **NAS**: NAS
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
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

func (s *OperateBucketScanTaskRequest) GetSource() *string {
	return s.Source
}

func (s *OperateBucketScanTaskRequest) SetBucketName(v string) *OperateBucketScanTaskRequest {
	s.BucketName = &v
	return s
}

func (s *OperateBucketScanTaskRequest) SetOperateCode(v int32) *OperateBucketScanTaskRequest {
	s.OperateCode = &v
	return s
}

func (s *OperateBucketScanTaskRequest) SetSource(v string) *OperateBucketScanTaskRequest {
	s.Source = &v
	return s
}

func (s *OperateBucketScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
