// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerAliUid(v string) *DescribeClientsRequest
	GetCallerAliUid() *string
	SetClientType(v int32) *DescribeClientsRequest
	GetClientType() *int32
	SetCustomResourceId(v string) *DescribeClientsRequest
	GetCustomResourceId() *string
	SetCustomResourceStatus(v bool) *DescribeClientsRequest
	GetCustomResourceStatus() *bool
	SetInManage(v bool) *DescribeClientsRequest
	GetInManage() *bool
	SetIncludeSubGroups(v bool) *DescribeClientsRequest
	GetIncludeSubGroups() *bool
	SetMaxResults(v int32) *DescribeClientsRequest
	GetMaxResults() *int32
	SetModel(v string) *DescribeClientsRequest
	GetModel() *string
	SetNextToken(v string) *DescribeClientsRequest
	GetNextToken() *string
	SetOnlineStatus(v bool) *DescribeClientsRequest
	GetOnlineStatus() *bool
	SetPlatform(v string) *DescribeClientsRequest
	GetPlatform() *string
	SetSearchKeyword(v string) *DescribeClientsRequest
	GetSearchKeyword() *string
	SetTerminalGroupId(v string) *DescribeClientsRequest
	GetTerminalGroupId() *string
	SetUuids(v []*string) *DescribeClientsRequest
	GetUuids() []*string
	SetWithBindUser(v bool) *DescribeClientsRequest
	GetWithBindUser() *bool
}

type DescribeClientsRequest struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// ***
	CallerAliUid *string `json:"CallerAliUid,omitempty" xml:"CallerAliUid,omitempty"`
	// The client type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The custom task ID.
	//
	// example:
	//
	// cr-***
	CustomResourceId *string `json:"CustomResourceId,omitempty" xml:"CustomResourceId,omitempty"`
	// The effective status of the custom task. Valid values:
	//
	// - true: effective.
	//
	// - false: ineffective.
	//
	// - null: all.
	CustomResourceStatus *bool `json:"CustomResourceStatus,omitempty" xml:"CustomResourceStatus,omitempty"`
	// Specifies whether the client is managed.
	//
	// example:
	//
	// True
	InManage *bool `json:"InManage,omitempty" xml:"InManage,omitempty"`
	// Specifies whether to include subgroups.
	//
	// example:
	//
	// false
	IncludeSubGroups *bool `json:"IncludeSubGroups,omitempty" xml:"IncludeSubGroups,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The device model.
	//
	// example:
	//
	// US02-2BFXG
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// AAAAAdEdsXbwG2ZlbWCzN4wTTg6wQvfp7u1BJl4bxCAby41POSaYAlCvfULQpkAnb0ff****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The soft client status.
	//
	// example:
	//
	// True
	OnlineStatus *bool `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The soft client platform.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The fuzzy search parameter. Supports fuzzy search by SN, alias, or IP address. This parameter is incompatible with exact search parameters.
	//
	// example:
	//
	// 061
	SearchKeyword *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	// The group ID.
	//
	// example:
	//
	// tg-bp103v8x70nasykdjrd1
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	// The UUIDs of the servers to query. Separate multiple UUIDs with commas (,).
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
	// Specifies whether to return bound users.
	WithBindUser *bool `json:"WithBindUser,omitempty" xml:"WithBindUser,omitempty"`
}

func (s DescribeClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientsRequest) GetCallerAliUid() *string {
	return s.CallerAliUid
}

func (s *DescribeClientsRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *DescribeClientsRequest) GetCustomResourceId() *string {
	return s.CustomResourceId
}

func (s *DescribeClientsRequest) GetCustomResourceStatus() *bool {
	return s.CustomResourceStatus
}

func (s *DescribeClientsRequest) GetInManage() *bool {
	return s.InManage
}

func (s *DescribeClientsRequest) GetIncludeSubGroups() *bool {
	return s.IncludeSubGroups
}

func (s *DescribeClientsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeClientsRequest) GetModel() *string {
	return s.Model
}

func (s *DescribeClientsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeClientsRequest) GetOnlineStatus() *bool {
	return s.OnlineStatus
}

func (s *DescribeClientsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeClientsRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *DescribeClientsRequest) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *DescribeClientsRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *DescribeClientsRequest) GetWithBindUser() *bool {
	return s.WithBindUser
}

func (s *DescribeClientsRequest) SetCallerAliUid(v string) *DescribeClientsRequest {
	s.CallerAliUid = &v
	return s
}

func (s *DescribeClientsRequest) SetClientType(v int32) *DescribeClientsRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeClientsRequest) SetCustomResourceId(v string) *DescribeClientsRequest {
	s.CustomResourceId = &v
	return s
}

func (s *DescribeClientsRequest) SetCustomResourceStatus(v bool) *DescribeClientsRequest {
	s.CustomResourceStatus = &v
	return s
}

func (s *DescribeClientsRequest) SetInManage(v bool) *DescribeClientsRequest {
	s.InManage = &v
	return s
}

func (s *DescribeClientsRequest) SetIncludeSubGroups(v bool) *DescribeClientsRequest {
	s.IncludeSubGroups = &v
	return s
}

func (s *DescribeClientsRequest) SetMaxResults(v int32) *DescribeClientsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeClientsRequest) SetModel(v string) *DescribeClientsRequest {
	s.Model = &v
	return s
}

func (s *DescribeClientsRequest) SetNextToken(v string) *DescribeClientsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeClientsRequest) SetOnlineStatus(v bool) *DescribeClientsRequest {
	s.OnlineStatus = &v
	return s
}

func (s *DescribeClientsRequest) SetPlatform(v string) *DescribeClientsRequest {
	s.Platform = &v
	return s
}

func (s *DescribeClientsRequest) SetSearchKeyword(v string) *DescribeClientsRequest {
	s.SearchKeyword = &v
	return s
}

func (s *DescribeClientsRequest) SetTerminalGroupId(v string) *DescribeClientsRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *DescribeClientsRequest) SetUuids(v []*string) *DescribeClientsRequest {
	s.Uuids = v
	return s
}

func (s *DescribeClientsRequest) SetWithBindUser(v bool) *DescribeClientsRequest {
	s.WithBindUser = &v
	return s
}

func (s *DescribeClientsRequest) Validate() error {
	return dara.Validate(s)
}
