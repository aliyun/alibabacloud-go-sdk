// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAnalysisJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisJobIds(v string) *QueryAnalysisJobListRequest
	GetAnalysisJobIds() *string
	SetOwnerAccount(v string) *QueryAnalysisJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryAnalysisJobListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryAnalysisJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAnalysisJobListRequest
	GetResourceOwnerId() *int64
}

type QueryAnalysisJobListRequest struct {
	// The template analysis job ID list.
	//
	// This parameter is required.
	//
	// example:
	//
	// bb558c1cc25b45309aab5be44d19****
	AnalysisJobIds       *string `json:"AnalysisJobIds,omitempty" xml:"AnalysisJobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryAnalysisJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListRequest) GetAnalysisJobIds() *string {
	return s.AnalysisJobIds
}

func (s *QueryAnalysisJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryAnalysisJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAnalysisJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAnalysisJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAnalysisJobListRequest) SetAnalysisJobIds(v string) *QueryAnalysisJobListRequest {
	s.AnalysisJobIds = &v
	return s
}

func (s *QueryAnalysisJobListRequest) SetOwnerAccount(v string) *QueryAnalysisJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryAnalysisJobListRequest) SetOwnerId(v int64) *QueryAnalysisJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAnalysisJobListRequest) SetResourceOwnerAccount(v string) *QueryAnalysisJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAnalysisJobListRequest) SetResourceOwnerId(v int64) *QueryAnalysisJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAnalysisJobListRequest) Validate() error {
	return dara.Validate(s)
}
