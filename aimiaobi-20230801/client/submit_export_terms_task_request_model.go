// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitExportTermsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTermsName(v string) *SubmitExportTermsTaskRequest
	GetTermsName() *string
	SetWorkspaceId(v string) *SubmitExportTermsTaskRequest
	GetWorkspaceId() *string
}

type SubmitExportTermsTaskRequest struct {
	TermsName *string `json:"TermsName,omitempty" xml:"TermsName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitExportTermsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitExportTermsTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitExportTermsTaskRequest) GetTermsName() *string {
	return s.TermsName
}

func (s *SubmitExportTermsTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitExportTermsTaskRequest) SetTermsName(v string) *SubmitExportTermsTaskRequest {
	s.TermsName = &v
	return s
}

func (s *SubmitExportTermsTaskRequest) SetWorkspaceId(v string) *SubmitExportTermsTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitExportTermsTaskRequest) Validate() error {
	return dara.Validate(s)
}
