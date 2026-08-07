// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CopyAppConfigRequest
	GetAppId() *string
	SetName(v string) *CopyAppConfigRequest
	GetName() *string
	SetRegionId(v string) *CopyAppConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *CopyAppConfigRequest
	GetResourceType() *string
}

type CopyAppConfigRequest struct {
	// The ID of the source App to copy from.
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the new App to be created from the copy.
	//
	// example:
	//
	// CustomTextModeration
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s CopyAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyAppConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyAppConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *CopyAppConfigRequest) GetName() *string {
	return s.Name
}

func (s *CopyAppConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyAppConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CopyAppConfigRequest) SetAppId(v string) *CopyAppConfigRequest {
	s.AppId = &v
	return s
}

func (s *CopyAppConfigRequest) SetName(v string) *CopyAppConfigRequest {
	s.Name = &v
	return s
}

func (s *CopyAppConfigRequest) SetRegionId(v string) *CopyAppConfigRequest {
	s.RegionId = &v
	return s
}

func (s *CopyAppConfigRequest) SetResourceType(v string) *CopyAppConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *CopyAppConfigRequest) Validate() error {
	return dara.Validate(s)
}
