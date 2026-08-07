// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConfigHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppConfigHistoryRequest
	GetAppId() *string
	SetRegionId(v string) *ListAppConfigHistoryRequest
	GetRegionId() *string
	SetResourceType(v string) *ListAppConfigHistoryRequest
	GetResourceType() *string
}

type ListAppConfigHistoryRequest struct {
	// App ID。
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAppConfigHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListAppConfigHistoryRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppConfigHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAppConfigHistoryRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppConfigHistoryRequest) SetAppId(v string) *ListAppConfigHistoryRequest {
	s.AppId = &v
	return s
}

func (s *ListAppConfigHistoryRequest) SetRegionId(v string) *ListAppConfigHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListAppConfigHistoryRequest) SetResourceType(v string) *ListAppConfigHistoryRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAppConfigHistoryRequest) Validate() error {
	return dara.Validate(s)
}
