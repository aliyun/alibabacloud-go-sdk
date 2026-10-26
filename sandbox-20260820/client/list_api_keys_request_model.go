// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyName(v string) *ListApiKeysRequest
	GetApiKeyName() *string
	SetPageNumber(v int32) *ListApiKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApiKeysRequest
	GetPageSize() *int32
	SetResourceGroupID(v string) *ListApiKeysRequest
	GetResourceGroupID() *string
	SetStatus(v string) *ListApiKeysRequest
	GetStatus() *string
	SetTeamID(v string) *ListApiKeysRequest
	GetTeamID() *string
	SetUserID(v string) *ListApiKeysRequest
	GetUserID() *string
}

type ListApiKeysRequest struct {
	// The name of the API key.
	//
	// example:
	//
	// dev
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The page number. Default value: 1. Valid values: 1 to 10000.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of API keys to return on each page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmwxqyrgwabcd
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// Filters API keys by status. Valid values:
	//
	// - active: Enabled.
	//
	// - inactive: Disabled.
	//
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 123456789
	UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s ListApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListApiKeysRequest) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *ListApiKeysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApiKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiKeysRequest) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *ListApiKeysRequest) GetStatus() *string {
	return s.Status
}

func (s *ListApiKeysRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *ListApiKeysRequest) GetUserID() *string {
	return s.UserID
}

func (s *ListApiKeysRequest) SetApiKeyName(v string) *ListApiKeysRequest {
	s.ApiKeyName = &v
	return s
}

func (s *ListApiKeysRequest) SetPageNumber(v int32) *ListApiKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApiKeysRequest) SetPageSize(v int32) *ListApiKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListApiKeysRequest) SetResourceGroupID(v string) *ListApiKeysRequest {
	s.ResourceGroupID = &v
	return s
}

func (s *ListApiKeysRequest) SetStatus(v string) *ListApiKeysRequest {
	s.Status = &v
	return s
}

func (s *ListApiKeysRequest) SetTeamID(v string) *ListApiKeysRequest {
	s.TeamID = &v
	return s
}

func (s *ListApiKeysRequest) SetUserID(v string) *ListApiKeysRequest {
	s.UserID = &v
	return s
}

func (s *ListApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
