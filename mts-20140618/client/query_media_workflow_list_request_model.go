// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaWorkflowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowIds(v string) *QueryMediaWorkflowListRequest
	GetMediaWorkflowIds() *string
	SetOwnerAccount(v string) *QueryMediaWorkflowListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryMediaWorkflowListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryMediaWorkflowListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMediaWorkflowListRequest
	GetResourceOwnerId() *int64
}

type QueryMediaWorkflowListRequest struct {
	// The IDs of the media workflows that you want to query. To obtain the IDs of the media workflows, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Workflow Settings*	- in the left-side navigation pane. You can query up to 10 media workflows at a time. Separate multiple IDs of media workflows with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e9181d4****,72dfa5e679ab4be9a3ed9974c736****
	MediaWorkflowIds     *string `json:"MediaWorkflowIds,omitempty" xml:"MediaWorkflowIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMediaWorkflowListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowListRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowListRequest) GetMediaWorkflowIds() *string {
	return s.MediaWorkflowIds
}

func (s *QueryMediaWorkflowListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryMediaWorkflowListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMediaWorkflowListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMediaWorkflowListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMediaWorkflowListRequest) SetMediaWorkflowIds(v string) *QueryMediaWorkflowListRequest {
	s.MediaWorkflowIds = &v
	return s
}

func (s *QueryMediaWorkflowListRequest) SetOwnerAccount(v string) *QueryMediaWorkflowListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryMediaWorkflowListRequest) SetOwnerId(v int64) *QueryMediaWorkflowListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMediaWorkflowListRequest) SetResourceOwnerAccount(v string) *QueryMediaWorkflowListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMediaWorkflowListRequest) SetResourceOwnerId(v int64) *QueryMediaWorkflowListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMediaWorkflowListRequest) Validate() error {
	return dara.Validate(s)
}
