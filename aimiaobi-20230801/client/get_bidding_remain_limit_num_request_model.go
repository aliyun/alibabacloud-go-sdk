// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBiddingRemainLimitNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *GetBiddingRemainLimitNumRequest
	GetApiName() *string
	SetWorkspaceId(v string) *GetBiddingRemainLimitNumRequest
	GetWorkspaceId() *string
}

type GetBiddingRemainLimitNumRequest struct {
	// example:
	//
	// asyncUploadTenderDoc
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetBiddingRemainLimitNumRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingRemainLimitNumRequest) GoString() string {
	return s.String()
}

func (s *GetBiddingRemainLimitNumRequest) GetApiName() *string {
	return s.ApiName
}

func (s *GetBiddingRemainLimitNumRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetBiddingRemainLimitNumRequest) SetApiName(v string) *GetBiddingRemainLimitNumRequest {
	s.ApiName = &v
	return s
}

func (s *GetBiddingRemainLimitNumRequest) SetWorkspaceId(v string) *GetBiddingRemainLimitNumRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetBiddingRemainLimitNumRequest) Validate() error {
	return dara.Validate(s)
}
