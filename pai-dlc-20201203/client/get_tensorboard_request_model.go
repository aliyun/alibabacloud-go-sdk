// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTensorboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJodId(v string) *GetTensorboardRequest
	GetJodId() *string
	SetToken(v string) *GetTensorboardRequest
	GetToken() *string
	SetWorkspaceId(v string) *GetTensorboardRequest
	GetWorkspaceId() *string
}

type GetTensorboardRequest struct {
	// The job ID. Refer to [ListJobs](https://help.aliyun.com/document_detail/459676.html) to obtain the job ID.
	//
	// example:
	//
	// dlc-xxxxxxxx
	JodId *string `json:"JodId,omitempty" xml:"JodId,omitempty"`
	// The sharing token. Specify this parameter to use the sharing token to obtain the permission to view a specific Tensorboard job. You can extract the token from the URL returned by calling [GetTensorboardSharedUrl](https://help.aliyun.com/document_detail/2557813.html).
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e
	//
	// yJleHAiOjE2OTUyODA0NTMsImlhdCI6MTY5NTE5NDA1MywidXNlcl9pZCI6IjExN
	//
	// Tc3MDMyNzA5OTQ5MDEiLCJ0YXJnZXRfaWQiOiJ0YjRrOGxjNXhmdTM2b3B0Iiw
	//
	// idGFyZ2V0X3R5cGUiOiJ0ZW5zb3Jib2FyZCJ9.6eT68J-KMBwwfN2d7fj7u6vyPcf0erfqYeizd2N****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The workspace ID. <props="china">Refer to [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID..
	//
	// example:
	//
	// 46099
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTensorboardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTensorboardRequest) GoString() string {
	return s.String()
}

func (s *GetTensorboardRequest) GetJodId() *string {
	return s.JodId
}

func (s *GetTensorboardRequest) GetToken() *string {
	return s.Token
}

func (s *GetTensorboardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTensorboardRequest) SetJodId(v string) *GetTensorboardRequest {
	s.JodId = &v
	return s
}

func (s *GetTensorboardRequest) SetToken(v string) *GetTensorboardRequest {
	s.Token = &v
	return s
}

func (s *GetTensorboardRequest) SetWorkspaceId(v string) *GetTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetTensorboardRequest) Validate() error {
	return dara.Validate(s)
}
