// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDirectoryFailedSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *RetryDirectoryFailedSourcesRequest
	GetDirectoryId() *string
	SetTenantId(v string) *RetryDirectoryFailedSourcesRequest
	GetTenantId() *string
}

type RetryDirectoryFailedSourcesRequest struct {
	// 目录 ID（递归包含子目录下的失败资源）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RetryDirectoryFailedSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryDirectoryFailedSourcesRequest) GoString() string {
	return s.String()
}

func (s *RetryDirectoryFailedSourcesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RetryDirectoryFailedSourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RetryDirectoryFailedSourcesRequest) SetDirectoryId(v string) *RetryDirectoryFailedSourcesRequest {
	s.DirectoryId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesRequest) SetTenantId(v string) *RetryDirectoryFailedSourcesRequest {
	s.TenantId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesRequest) Validate() error {
	return dara.Validate(s)
}
