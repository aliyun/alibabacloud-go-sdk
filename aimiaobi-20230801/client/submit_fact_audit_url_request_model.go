// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFactAuditUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *SubmitFactAuditUrlRequest
	GetUrl() *string
	SetWorkspaceId(v string) *SubmitFactAuditUrlRequest
	GetWorkspaceId() *string
}

type SubmitFactAuditUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitFactAuditUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitFactAuditUrlRequest) GoString() string {
	return s.String()
}

func (s *SubmitFactAuditUrlRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitFactAuditUrlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitFactAuditUrlRequest) SetUrl(v string) *SubmitFactAuditUrlRequest {
	s.Url = &v
	return s
}

func (s *SubmitFactAuditUrlRequest) SetWorkspaceId(v string) *SubmitFactAuditUrlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitFactAuditUrlRequest) Validate() error {
	return dara.Validate(s)
}
