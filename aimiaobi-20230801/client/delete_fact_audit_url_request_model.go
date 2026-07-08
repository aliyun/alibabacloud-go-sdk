// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFactAuditUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *DeleteFactAuditUrlRequest
	GetUrl() *string
	SetWorkspaceId(v string) *DeleteFactAuditUrlRequest
	GetWorkspaceId() *string
}

type DeleteFactAuditUrlRequest struct {
	// The URL of the audit information source you want to delete. The provided URL must match the result from GetFactAuditUrl for successful deletion.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The unique identifier of the Alibaba Cloud Model Studio workspace. Get the [workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteFactAuditUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFactAuditUrlRequest) GoString() string {
	return s.String()
}

func (s *DeleteFactAuditUrlRequest) GetUrl() *string {
	return s.Url
}

func (s *DeleteFactAuditUrlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteFactAuditUrlRequest) SetUrl(v string) *DeleteFactAuditUrlRequest {
	s.Url = &v
	return s
}

func (s *DeleteFactAuditUrlRequest) SetWorkspaceId(v string) *DeleteFactAuditUrlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteFactAuditUrlRequest) Validate() error {
	return dara.Validate(s)
}
