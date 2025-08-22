// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSharedAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*int64) *AddSharedAccountsRequest
	GetAccountIds() []*int64
	SetResourceId(v string) *AddSharedAccountsRequest
	GetResourceId() *string
	SetResourceType(v string) *AddSharedAccountsRequest
	GetResourceType() *string
}

type AddSharedAccountsRequest struct {
	AccountIds []*int64 `json:"accountIds,omitempty" xml:"accountIds,omitempty" type:"Repeated"`
	// example:
	//
	// Public
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// RegistryModule
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s AddSharedAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *AddSharedAccountsRequest) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *AddSharedAccountsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddSharedAccountsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddSharedAccountsRequest) SetAccountIds(v []*int64) *AddSharedAccountsRequest {
	s.AccountIds = v
	return s
}

func (s *AddSharedAccountsRequest) SetResourceId(v string) *AddSharedAccountsRequest {
	s.ResourceId = &v
	return s
}

func (s *AddSharedAccountsRequest) SetResourceType(v string) *AddSharedAccountsRequest {
	s.ResourceType = &v
	return s
}

func (s *AddSharedAccountsRequest) Validate() error {
	return dara.Validate(s)
}
