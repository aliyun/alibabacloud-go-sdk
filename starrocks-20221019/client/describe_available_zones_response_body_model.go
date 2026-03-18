// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeAvailableZonesResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *DescribeAvailableZonesResponseBodyData) *DescribeAvailableZonesResponseBody
	GetData() *DescribeAvailableZonesResponseBodyData
	SetErrCode(v string) *DescribeAvailableZonesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeAvailableZonesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeAvailableZonesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeAvailableZonesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAvailableZonesResponseBody
	GetSuccess() *bool
}

type DescribeAvailableZonesResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *DescribeAvailableZonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeAvailableZonesResponseBody) GetData() *DescribeAvailableZonesResponseBodyData {
	return s.Data
}

func (s *DescribeAvailableZonesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeAvailableZonesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeAvailableZonesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableZonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAvailableZonesResponseBody) SetAccessDeniedDetail(v string) *DescribeAvailableZonesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeAvailableZonesResponseBody) SetData(v *DescribeAvailableZonesResponseBodyData) *DescribeAvailableZonesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAvailableZonesResponseBody) SetErrCode(v string) *DescribeAvailableZonesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeAvailableZonesResponseBody) SetErrMessage(v string) *DescribeAvailableZonesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeAvailableZonesResponseBody) SetHttpStatusCode(v int32) *DescribeAvailableZonesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAvailableZonesResponseBody) SetRequestId(v string) *DescribeAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableZonesResponseBody) SetSuccess(v bool) *DescribeAvailableZonesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAvailableZonesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableZonesResponseBodyData struct {
	OfficialAvailableZones []*string `json:"OfficialAvailableZones,omitempty" xml:"OfficialAvailableZones,omitempty" type:"Repeated"`
	TrialAvailableZones    []*string `json:"TrialAvailableZones,omitempty" xml:"TrialAvailableZones,omitempty" type:"Repeated"`
}

func (s DescribeAvailableZonesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBodyData) GetOfficialAvailableZones() []*string {
	return s.OfficialAvailableZones
}

func (s *DescribeAvailableZonesResponseBodyData) GetTrialAvailableZones() []*string {
	return s.TrialAvailableZones
}

func (s *DescribeAvailableZonesResponseBodyData) SetOfficialAvailableZones(v []*string) *DescribeAvailableZonesResponseBodyData {
	s.OfficialAvailableZones = v
	return s
}

func (s *DescribeAvailableZonesResponseBodyData) SetTrialAvailableZones(v []*string) *DescribeAvailableZonesResponseBodyData {
	s.TrialAvailableZones = v
	return s
}

func (s *DescribeAvailableZonesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
