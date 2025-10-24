// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaWorkflowExecutionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *QueryMediaWorkflowExecutionListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryMediaWorkflowExecutionListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryMediaWorkflowExecutionListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMediaWorkflowExecutionListRequest
	GetResourceOwnerId() *int64
	SetRunIds(v string) *QueryMediaWorkflowExecutionListRequest
	GetRunIds() *string
}

type QueryMediaWorkflowExecutionListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the media workflow execution instances. To obtain the instance ID, log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Execution Instances*	- in the left-side navigation pane. Separate multiple IDs with commas (,). You can query a maximum of 10 media workflow execution instances at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 48e33690ac19445488c706924321****
	RunIds *string `json:"RunIds,omitempty" xml:"RunIds,omitempty"`
}

func (s QueryMediaWorkflowExecutionListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryMediaWorkflowExecutionListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMediaWorkflowExecutionListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMediaWorkflowExecutionListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMediaWorkflowExecutionListRequest) GetRunIds() *string {
	return s.RunIds
}

func (s *QueryMediaWorkflowExecutionListRequest) SetOwnerAccount(v string) *QueryMediaWorkflowExecutionListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListRequest) SetOwnerId(v int64) *QueryMediaWorkflowExecutionListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListRequest) SetResourceOwnerAccount(v string) *QueryMediaWorkflowExecutionListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListRequest) SetResourceOwnerId(v int64) *QueryMediaWorkflowExecutionListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListRequest) SetRunIds(v string) *QueryMediaWorkflowExecutionListRequest {
	s.RunIds = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListRequest) Validate() error {
	return dara.Validate(s)
}
