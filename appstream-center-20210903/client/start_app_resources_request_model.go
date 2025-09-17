// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAppResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *StartAppResourcesRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *StartAppResourcesRequest
	GetBizRegionId() *string
	SetClientId(v string) *StartAppResourcesRequest
	GetClientId() *string
	SetClientIp(v string) *StartAppResourcesRequest
	GetClientIp() *string
	SetClientOS(v string) *StartAppResourcesRequest
	GetClientOS() *string
	SetClientVersion(v string) *StartAppResourcesRequest
	GetClientVersion() *string
	SetEndUserId(v string) *StartAppResourcesRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *StartAppResourcesRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *StartAppResourcesRequest
	GetLoginToken() *string
	SetProductType(v string) *StartAppResourcesRequest
	GetProductType() *string
	SetResourceIds(v []*string) *StartAppResourcesRequest
	GetResourceIds() []*string
	SetSessionId(v string) *StartAppResourcesRequest
	GetSessionId() *string
	SetUuid(v string) *StartAppResourcesRequest
	GetUuid() *string
}

type StartAppResourcesRequest struct {
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
	// eac19bef-1e45-4190-a03a-4ea74b69****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 1.2.3.4
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Pro\\" 10.0 (Build 22631)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 7.5.1-R-20240903.163046
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// testUser01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v185fdd7f6d39fa7861981639366085772e150a390a5bb7b43c4e62440d94fc392b945770e1596cebe90085ce0af4d****
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
	// a863f4c3-2f1d-4971-8cf7-e2b92ae9****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s StartAppResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAppResourcesRequest) GoString() string {
	return s.String()
}

func (s *StartAppResourcesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *StartAppResourcesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *StartAppResourcesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StartAppResourcesRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *StartAppResourcesRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *StartAppResourcesRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *StartAppResourcesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *StartAppResourcesRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *StartAppResourcesRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *StartAppResourcesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *StartAppResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *StartAppResourcesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartAppResourcesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *StartAppResourcesRequest) SetAppInstanceGroupId(v string) *StartAppResourcesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *StartAppResourcesRequest) SetBizRegionId(v string) *StartAppResourcesRequest {
	s.BizRegionId = &v
	return s
}

func (s *StartAppResourcesRequest) SetClientId(v string) *StartAppResourcesRequest {
	s.ClientId = &v
	return s
}

func (s *StartAppResourcesRequest) SetClientIp(v string) *StartAppResourcesRequest {
	s.ClientIp = &v
	return s
}

func (s *StartAppResourcesRequest) SetClientOS(v string) *StartAppResourcesRequest {
	s.ClientOS = &v
	return s
}

func (s *StartAppResourcesRequest) SetClientVersion(v string) *StartAppResourcesRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartAppResourcesRequest) SetEndUserId(v string) *StartAppResourcesRequest {
	s.EndUserId = &v
	return s
}

func (s *StartAppResourcesRequest) SetLoginRegionId(v string) *StartAppResourcesRequest {
	s.LoginRegionId = &v
	return s
}

func (s *StartAppResourcesRequest) SetLoginToken(v string) *StartAppResourcesRequest {
	s.LoginToken = &v
	return s
}

func (s *StartAppResourcesRequest) SetProductType(v string) *StartAppResourcesRequest {
	s.ProductType = &v
	return s
}

func (s *StartAppResourcesRequest) SetResourceIds(v []*string) *StartAppResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *StartAppResourcesRequest) SetSessionId(v string) *StartAppResourcesRequest {
	s.SessionId = &v
	return s
}

func (s *StartAppResourcesRequest) SetUuid(v string) *StartAppResourcesRequest {
	s.Uuid = &v
	return s
}

func (s *StartAppResourcesRequest) Validate() error {
	return dara.Validate(s)
}
