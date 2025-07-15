// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubCodes(v []*string) *SubmitSmartAuditRequest
	GetSubCodes() []*string
	SetText(v string) *SubmitSmartAuditRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitSmartAuditRequest
	GetWorkspaceId() *string
}

type SubmitSmartAuditRequest struct {
	SubCodes []*string `json:"SubCodes,omitempty" xml:"SubCodes,omitempty" type:"Repeated"`
	Text     *string   `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitSmartAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditRequest) GetSubCodes() []*string {
	return s.SubCodes
}

func (s *SubmitSmartAuditRequest) GetText() *string {
	return s.Text
}

func (s *SubmitSmartAuditRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitSmartAuditRequest) SetSubCodes(v []*string) *SubmitSmartAuditRequest {
	s.SubCodes = v
	return s
}

func (s *SubmitSmartAuditRequest) SetText(v string) *SubmitSmartAuditRequest {
	s.Text = &v
	return s
}

func (s *SubmitSmartAuditRequest) SetWorkspaceId(v string) *SubmitSmartAuditRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitSmartAuditRequest) Validate() error {
	return dara.Validate(s)
}
