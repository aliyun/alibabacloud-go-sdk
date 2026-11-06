// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKVCacheInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKVCacheInstanceId(v string) *DescribeKVCacheInstanceAttributeRequest
	GetKVCacheInstanceId() *string
}

type DescribeKVCacheInstanceAttributeRequest struct {
	// The ID of the KVCache instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkv-123456
	KVCacheInstanceId *string `json:"KVCacheInstanceId,omitempty" xml:"KVCacheInstanceId,omitempty"`
}

func (s DescribeKVCacheInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstanceAttributeRequest) GetKVCacheInstanceId() *string {
	return s.KVCacheInstanceId
}

func (s *DescribeKVCacheInstanceAttributeRequest) SetKVCacheInstanceId(v string) *DescribeKVCacheInstanceAttributeRequest {
	s.KVCacheInstanceId = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
