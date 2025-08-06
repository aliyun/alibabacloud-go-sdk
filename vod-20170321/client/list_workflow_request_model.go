// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListWorkflowRequest
	GetAppId() *string
	SetBizVersion(v string) *ListWorkflowRequest
	GetBizVersion() *string
	SetOwnerId(v int64) *ListWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListWorkflowRequest
	GetResourceOwnerId() *int64
	SetState(v string) *ListWorkflowRequest
	GetState() *string
}

type ListWorkflowRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizVersion           *string `json:"BizVersion,omitempty" xml:"BizVersion,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListWorkflowRequest) GetBizVersion() *string {
	return s.BizVersion
}

func (s *ListWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListWorkflowRequest) GetState() *string {
	return s.State
}

func (s *ListWorkflowRequest) SetAppId(v string) *ListWorkflowRequest {
	s.AppId = &v
	return s
}

func (s *ListWorkflowRequest) SetBizVersion(v string) *ListWorkflowRequest {
	s.BizVersion = &v
	return s
}

func (s *ListWorkflowRequest) SetOwnerId(v int64) *ListWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *ListWorkflowRequest) SetResourceOwnerAccount(v string) *ListWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListWorkflowRequest) SetResourceOwnerId(v int64) *ListWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListWorkflowRequest) SetState(v string) *ListWorkflowRequest {
	s.State = &v
	return s
}

func (s *ListWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
