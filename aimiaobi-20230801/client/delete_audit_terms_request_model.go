// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditTermsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdList(v []*int64) *DeleteAuditTermsRequest
	GetIdList() []*int64
	SetWorkspaceId(v string) *DeleteAuditTermsRequest
	GetWorkspaceId() *string
}

type DeleteAuditTermsRequest struct {
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteAuditTermsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuditTermsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuditTermsRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *DeleteAuditTermsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteAuditTermsRequest) SetIdList(v []*int64) *DeleteAuditTermsRequest {
	s.IdList = v
	return s
}

func (s *DeleteAuditTermsRequest) SetWorkspaceId(v string) *DeleteAuditTermsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteAuditTermsRequest) Validate() error {
	return dara.Validate(s)
}
