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
	// The ID of the enterprise knowledge base folder. Failed resources in the folder and its subfolders are included recursively.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
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
