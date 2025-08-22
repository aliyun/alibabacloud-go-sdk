// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSharedAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*int64) *RemoveSharedAccountsRequest
	GetAccountIds() []*int64
	SetResourceId(v string) *RemoveSharedAccountsRequest
	GetResourceId() *string
	SetResourceType(v string) *RemoveSharedAccountsRequest
	GetResourceType() *string
}

type RemoveSharedAccountsRequest struct {
	// This parameter is required.
	AccountIds []*int64 `json:"accountIds,omitempty" xml:"accountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Namespace
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s RemoveSharedAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *RemoveSharedAccountsRequest) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *RemoveSharedAccountsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *RemoveSharedAccountsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RemoveSharedAccountsRequest) SetAccountIds(v []*int64) *RemoveSharedAccountsRequest {
	s.AccountIds = v
	return s
}

func (s *RemoveSharedAccountsRequest) SetResourceId(v string) *RemoveSharedAccountsRequest {
	s.ResourceId = &v
	return s
}

func (s *RemoveSharedAccountsRequest) SetResourceType(v string) *RemoveSharedAccountsRequest {
	s.ResourceType = &v
	return s
}

func (s *RemoveSharedAccountsRequest) Validate() error {
	return dara.Validate(s)
}
