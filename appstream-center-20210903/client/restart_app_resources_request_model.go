// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartAppResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *RestartAppResourcesRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *RestartAppResourcesRequest
	GetBizRegionId() *string
	SetClientId(v string) *RestartAppResourcesRequest
	GetClientId() *string
	SetClientIp(v string) *RestartAppResourcesRequest
	GetClientIp() *string
	SetClientOS(v string) *RestartAppResourcesRequest
	GetClientOS() *string
	SetClientVersion(v string) *RestartAppResourcesRequest
	GetClientVersion() *string
	SetEndUserId(v string) *RestartAppResourcesRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *RestartAppResourcesRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *RestartAppResourcesRequest
	GetLoginToken() *string
	SetProductType(v string) *RestartAppResourcesRequest
	GetProductType() *string
	SetResourceIds(v []*string) *RestartAppResourcesRequest
	GetResourceIds() []*string
	SetSessionId(v string) *RestartAppResourcesRequest
	GetSessionId() *string
	SetUuid(v string) *RestartAppResourcesRequest
	GetUuid() *string
}

type RestartAppResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aig-53fvrq1oanz6c****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eac19bef-1e45-4190-a03a-4ea74b699ca7
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 1.2.3.4
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Enterprise\\" 10.0 (Build 14393)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 7.5.3-RS-20241127.131156
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// user001
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1124500957832f30b3e716406562071655aa43b2a723ed2be0837815483d54e025db13ba5469f06f2410d0efc4d302e36
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AndroidCloud
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// c261a6a1-e242-4f4b-813c-5fe807e49f03
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RestartAppResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartAppResourcesRequest) GoString() string {
	return s.String()
}

func (s *RestartAppResourcesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *RestartAppResourcesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *RestartAppResourcesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RestartAppResourcesRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *RestartAppResourcesRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *RestartAppResourcesRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *RestartAppResourcesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RestartAppResourcesRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *RestartAppResourcesRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *RestartAppResourcesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RestartAppResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *RestartAppResourcesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RestartAppResourcesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *RestartAppResourcesRequest) SetAppInstanceGroupId(v string) *RestartAppResourcesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *RestartAppResourcesRequest) SetBizRegionId(v string) *RestartAppResourcesRequest {
	s.BizRegionId = &v
	return s
}

func (s *RestartAppResourcesRequest) SetClientId(v string) *RestartAppResourcesRequest {
	s.ClientId = &v
	return s
}

func (s *RestartAppResourcesRequest) SetClientIp(v string) *RestartAppResourcesRequest {
	s.ClientIp = &v
	return s
}

func (s *RestartAppResourcesRequest) SetClientOS(v string) *RestartAppResourcesRequest {
	s.ClientOS = &v
	return s
}

func (s *RestartAppResourcesRequest) SetClientVersion(v string) *RestartAppResourcesRequest {
	s.ClientVersion = &v
	return s
}

func (s *RestartAppResourcesRequest) SetEndUserId(v string) *RestartAppResourcesRequest {
	s.EndUserId = &v
	return s
}

func (s *RestartAppResourcesRequest) SetLoginRegionId(v string) *RestartAppResourcesRequest {
	s.LoginRegionId = &v
	return s
}

func (s *RestartAppResourcesRequest) SetLoginToken(v string) *RestartAppResourcesRequest {
	s.LoginToken = &v
	return s
}

func (s *RestartAppResourcesRequest) SetProductType(v string) *RestartAppResourcesRequest {
	s.ProductType = &v
	return s
}

func (s *RestartAppResourcesRequest) SetResourceIds(v []*string) *RestartAppResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *RestartAppResourcesRequest) SetSessionId(v string) *RestartAppResourcesRequest {
	s.SessionId = &v
	return s
}

func (s *RestartAppResourcesRequest) SetUuid(v string) *RestartAppResourcesRequest {
	s.Uuid = &v
	return s
}

func (s *RestartAppResourcesRequest) Validate() error {
	return dara.Validate(s)
}
