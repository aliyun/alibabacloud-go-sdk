// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewMCTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *PreviewMCTableRequest
	GetEndpoint() *string
	SetLimit(v int32) *PreviewMCTableRequest
	GetLimit() *int32
	SetPartition(v string) *PreviewMCTableRequest
	GetPartition() *string
	SetWorkspaceId(v string) *PreviewMCTableRequest
	GetWorkspaceId() *string
}

type PreviewMCTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://service.cn-hangzhou-vpc.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 1000
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// pt=20240805
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PreviewMCTableRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewMCTableRequest) GoString() string {
	return s.String()
}

func (s *PreviewMCTableRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *PreviewMCTableRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *PreviewMCTableRequest) GetPartition() *string {
	return s.Partition
}

func (s *PreviewMCTableRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PreviewMCTableRequest) SetEndpoint(v string) *PreviewMCTableRequest {
	s.Endpoint = &v
	return s
}

func (s *PreviewMCTableRequest) SetLimit(v int32) *PreviewMCTableRequest {
	s.Limit = &v
	return s
}

func (s *PreviewMCTableRequest) SetPartition(v string) *PreviewMCTableRequest {
	s.Partition = &v
	return s
}

func (s *PreviewMCTableRequest) SetWorkspaceId(v string) *PreviewMCTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PreviewMCTableRequest) Validate() error {
	return dara.Validate(s)
}
