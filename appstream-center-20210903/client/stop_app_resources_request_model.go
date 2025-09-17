// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAppResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *StopAppResourcesRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *StopAppResourcesRequest
	GetBizRegionId() *string
	SetClientId(v string) *StopAppResourcesRequest
	GetClientId() *string
	SetClientIp(v string) *StopAppResourcesRequest
	GetClientIp() *string
	SetClientOS(v string) *StopAppResourcesRequest
	GetClientOS() *string
	SetClientVersion(v string) *StopAppResourcesRequest
	GetClientVersion() *string
	SetEndUserId(v string) *StopAppResourcesRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *StopAppResourcesRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *StopAppResourcesRequest
	GetLoginToken() *string
	SetProductType(v string) *StopAppResourcesRequest
	GetProductType() *string
	SetResourceIds(v []*string) *StopAppResourcesRequest
	GetResourceIds() []*string
	SetSessionId(v string) *StopAppResourcesRequest
	GetSessionId() *string
	SetUuid(v string) *StopAppResourcesRequest
	GetUuid() *string
}

type StopAppResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
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
	// windows_"Windows 10 Enterprise" 10.0 (Build 19042)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 7.2.0-R-20241008.110000
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// testUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v12369636c721ba6b3ddb1683341016775c3f63e4d0e78f120f9a0544ed826b7af7daf747c402f0d0730b52f451b70****
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
	// 597e869d-ea14-4b83-9490-714f68bfe935
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// C50973691A6D2BE23F2CDD73B85B****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s StopAppResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAppResourcesRequest) GoString() string {
	return s.String()
}

func (s *StopAppResourcesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *StopAppResourcesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *StopAppResourcesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StopAppResourcesRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *StopAppResourcesRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *StopAppResourcesRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *StopAppResourcesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *StopAppResourcesRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *StopAppResourcesRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *StopAppResourcesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *StopAppResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *StopAppResourcesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopAppResourcesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *StopAppResourcesRequest) SetAppInstanceGroupId(v string) *StopAppResourcesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *StopAppResourcesRequest) SetBizRegionId(v string) *StopAppResourcesRequest {
	s.BizRegionId = &v
	return s
}

func (s *StopAppResourcesRequest) SetClientId(v string) *StopAppResourcesRequest {
	s.ClientId = &v
	return s
}

func (s *StopAppResourcesRequest) SetClientIp(v string) *StopAppResourcesRequest {
	s.ClientIp = &v
	return s
}

func (s *StopAppResourcesRequest) SetClientOS(v string) *StopAppResourcesRequest {
	s.ClientOS = &v
	return s
}

func (s *StopAppResourcesRequest) SetClientVersion(v string) *StopAppResourcesRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopAppResourcesRequest) SetEndUserId(v string) *StopAppResourcesRequest {
	s.EndUserId = &v
	return s
}

func (s *StopAppResourcesRequest) SetLoginRegionId(v string) *StopAppResourcesRequest {
	s.LoginRegionId = &v
	return s
}

func (s *StopAppResourcesRequest) SetLoginToken(v string) *StopAppResourcesRequest {
	s.LoginToken = &v
	return s
}

func (s *StopAppResourcesRequest) SetProductType(v string) *StopAppResourcesRequest {
	s.ProductType = &v
	return s
}

func (s *StopAppResourcesRequest) SetResourceIds(v []*string) *StopAppResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *StopAppResourcesRequest) SetSessionId(v string) *StopAppResourcesRequest {
	s.SessionId = &v
	return s
}

func (s *StopAppResourcesRequest) SetUuid(v string) *StopAppResourcesRequest {
	s.Uuid = &v
	return s
}

func (s *StopAppResourcesRequest) Validate() error {
	return dara.Validate(s)
}
