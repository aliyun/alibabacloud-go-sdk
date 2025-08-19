// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCloudAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *RemoveCloudAccountRequest
	GetAccountId() *string
}

type RemoveCloudAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s RemoveCloudAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *RemoveCloudAccountRequest) SetAccountId(v string) *RemoveCloudAccountRequest {
	s.AccountId = &v
	return s
}

func (s *RemoveCloudAccountRequest) Validate() error {
	return dara.Validate(s)
}
