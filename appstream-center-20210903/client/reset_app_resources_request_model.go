// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ResetAppResourcesRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *ResetAppResourcesRequest
	GetBizRegionId() *string
	SetClientId(v string) *ResetAppResourcesRequest
	GetClientId() *string
	SetClientIp(v string) *ResetAppResourcesRequest
	GetClientIp() *string
	SetClientOS(v string) *ResetAppResourcesRequest
	GetClientOS() *string
	SetClientVersion(v string) *ResetAppResourcesRequest
	GetClientVersion() *string
	SetEndUserId(v string) *ResetAppResourcesRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *ResetAppResourcesRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *ResetAppResourcesRequest
	GetLoginToken() *string
	SetProductType(v string) *ResetAppResourcesRequest
	GetProductType() *string
	SetResourceIds(v []*string) *ResetAppResourcesRequest
	GetResourceIds() []*string
	SetSessionId(v string) *ResetAppResourcesRequest
	GetSessionId() *string
	SetUuid(v string) *ResetAppResourcesRequest
	GetUuid() *string
}

type ResetAppResourcesRequest struct {
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
	// f4a0dc8e-1702-4728-9a60-95b27a35****
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
	// 7.7.0-R-20241217.092056
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v285fdd7f6d39fa7861981639366085772e150a390a5bb7b43c4e62440d94fc392b945770e1596cebe90085ce0af4d****
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

func (s ResetAppResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAppResourcesRequest) GoString() string {
	return s.String()
}

func (s *ResetAppResourcesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ResetAppResourcesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ResetAppResourcesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ResetAppResourcesRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *ResetAppResourcesRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *ResetAppResourcesRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ResetAppResourcesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ResetAppResourcesRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *ResetAppResourcesRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ResetAppResourcesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ResetAppResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ResetAppResourcesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ResetAppResourcesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ResetAppResourcesRequest) SetAppInstanceGroupId(v string) *ResetAppResourcesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ResetAppResourcesRequest) SetBizRegionId(v string) *ResetAppResourcesRequest {
	s.BizRegionId = &v
	return s
}

func (s *ResetAppResourcesRequest) SetClientId(v string) *ResetAppResourcesRequest {
	s.ClientId = &v
	return s
}

func (s *ResetAppResourcesRequest) SetClientIp(v string) *ResetAppResourcesRequest {
	s.ClientIp = &v
	return s
}

func (s *ResetAppResourcesRequest) SetClientOS(v string) *ResetAppResourcesRequest {
	s.ClientOS = &v
	return s
}

func (s *ResetAppResourcesRequest) SetClientVersion(v string) *ResetAppResourcesRequest {
	s.ClientVersion = &v
	return s
}

func (s *ResetAppResourcesRequest) SetEndUserId(v string) *ResetAppResourcesRequest {
	s.EndUserId = &v
	return s
}

func (s *ResetAppResourcesRequest) SetLoginRegionId(v string) *ResetAppResourcesRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ResetAppResourcesRequest) SetLoginToken(v string) *ResetAppResourcesRequest {
	s.LoginToken = &v
	return s
}

func (s *ResetAppResourcesRequest) SetProductType(v string) *ResetAppResourcesRequest {
	s.ProductType = &v
	return s
}

func (s *ResetAppResourcesRequest) SetResourceIds(v []*string) *ResetAppResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ResetAppResourcesRequest) SetSessionId(v string) *ResetAppResourcesRequest {
	s.SessionId = &v
	return s
}

func (s *ResetAppResourcesRequest) SetUuid(v string) *ResetAppResourcesRequest {
	s.Uuid = &v
	return s
}

func (s *ResetAppResourcesRequest) Validate() error {
	return dara.Validate(s)
}
