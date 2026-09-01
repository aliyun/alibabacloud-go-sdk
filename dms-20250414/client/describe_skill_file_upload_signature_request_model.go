// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSkillFileUploadSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeSkillFileUploadSignatureRequest
	GetWorkspaceId() *string
}

type DescribeSkillFileUploadSignatureRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// aci5e5yd***********0crv
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeSkillFileUploadSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillFileUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeSkillFileUploadSignatureRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeSkillFileUploadSignatureRequest) SetWorkspaceId(v string) *DescribeSkillFileUploadSignatureRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureRequest) Validate() error {
	return dara.Validate(s)
}
