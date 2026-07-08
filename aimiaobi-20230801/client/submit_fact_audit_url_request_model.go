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
	// The URL of the information source you want to use for factuality audit. After you add a source URL using this operation, MiaoBi retrieves audit information only from your configured list of URLs. If you do not add any URL, MiaoBi searches the entire web for audit information. You can configure up to 10 source URLs.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Unique identifier of the Alibaba Cloud Model Studio workspace. To get the workspace ID, see [Get the workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
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
