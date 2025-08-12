// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayTransferStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAlipayTransferStatusRequest
	GetCode() *string
	SetWorkspaceId(v string) *GetAlipayTransferStatusRequest
	GetWorkspaceId() *string
}

type GetAlipayTransferStatusRequest struct {
	// example:
	//
	// xxx-xxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// llm-cxxxxxxb8d47ks
	WorkspaceId *string `json:"workspace_id,omitempty" xml:"workspace_id,omitempty"`
}

func (s GetAlipayTransferStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayTransferStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAlipayTransferStatusRequest) GetCode() *string {
	return s.Code
}

func (s *GetAlipayTransferStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAlipayTransferStatusRequest) SetCode(v string) *GetAlipayTransferStatusRequest {
	s.Code = &v
	return s
}

func (s *GetAlipayTransferStatusRequest) SetWorkspaceId(v string) *GetAlipayTransferStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAlipayTransferStatusRequest) Validate() error {
	return dara.Validate(s)
}
