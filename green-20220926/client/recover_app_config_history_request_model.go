// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAppConfigHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RecoverAppConfigHistoryRequest
	GetAppId() *string
	SetAppVersion(v int64) *RecoverAppConfigHistoryRequest
	GetAppVersion() *int64
	SetRegionId(v string) *RecoverAppConfigHistoryRequest
	GetRegionId() *string
	SetResourceType(v string) *RecoverAppConfigHistoryRequest
	GetResourceType() *string
}

type RecoverAppConfigHistoryRequest struct {
	// App ID。
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1785898163
	AppVersion *int64 `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
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

func (s RecoverAppConfigHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverAppConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *RecoverAppConfigHistoryRequest) GetAppId() *string {
	return s.AppId
}

func (s *RecoverAppConfigHistoryRequest) GetAppVersion() *int64 {
	return s.AppVersion
}

func (s *RecoverAppConfigHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RecoverAppConfigHistoryRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RecoverAppConfigHistoryRequest) SetAppId(v string) *RecoverAppConfigHistoryRequest {
	s.AppId = &v
	return s
}

func (s *RecoverAppConfigHistoryRequest) SetAppVersion(v int64) *RecoverAppConfigHistoryRequest {
	s.AppVersion = &v
	return s
}

func (s *RecoverAppConfigHistoryRequest) SetRegionId(v string) *RecoverAppConfigHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *RecoverAppConfigHistoryRequest) SetResourceType(v string) *RecoverAppConfigHistoryRequest {
	s.ResourceType = &v
	return s
}

func (s *RecoverAppConfigHistoryRequest) Validate() error {
	return dara.Validate(s)
}
