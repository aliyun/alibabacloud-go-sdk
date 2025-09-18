// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCommandRequest
	GetAppId() *string
	SetDomainCode(v string) *DescribeCommandRequest
	GetDomainCode() *string
	SetToolId(v string) *DescribeCommandRequest
	GetToolId() *string
	SetWorkspaceId(v string) *DescribeCommandRequest
	GetWorkspaceId() *string
}

type DescribeCommandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm-xxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 343894343
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4864343453445
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommandRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCommandRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *DescribeCommandRequest) GetToolId() *string {
	return s.ToolId
}

func (s *DescribeCommandRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeCommandRequest) SetAppId(v string) *DescribeCommandRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCommandRequest) SetDomainCode(v string) *DescribeCommandRequest {
	s.DomainCode = &v
	return s
}

func (s *DescribeCommandRequest) SetToolId(v string) *DescribeCommandRequest {
	s.ToolId = &v
	return s
}

func (s *DescribeCommandRequest) SetWorkspaceId(v string) *DescribeCommandRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCommandRequest) Validate() error {
	return dara.Validate(s)
}
