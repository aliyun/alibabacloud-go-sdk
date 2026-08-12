// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeRegionsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody
	GetData() []*DescribeRegionsResponseBodyData
	SetErrCode(v string) *DescribeRegionsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeRegionsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeRegionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRegionsResponseBody
	GetSuccess() *bool
}

type DescribeRegionsResponseBody struct {
	// Details about access denial.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The instance information.
	Data []*DescribeRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP request status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeRegionsResponseBody) GetData() []*DescribeRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeRegionsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeRegionsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeRegionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRegionsResponseBody) SetAccessDeniedDetail(v string) *DescribeRegionsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetData(v []*DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrCode(v string) *DescribeRegionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrMessage(v string) *DescribeRegionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetHttpStatusCode(v int32) *DescribeRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyData struct {
	// The description of the region.
	//
	// example:
	//
	// Hangzhou
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The English description.
	//
	// example:
	//
	// cn-hangzhou
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// 华东1（杭州）
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeRegionsResponseBodyData) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *DescribeRegionsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyData) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeRegionsResponseBodyData) SetDescription(v string) *DescribeRegionsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeRegionsResponseBodyData) SetDescriptionEn(v string) *DescribeRegionsResponseBodyData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeRegionsResponseBodyData) SetRegionId(v string) *DescribeRegionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyData) SetRegionName(v string) *DescribeRegionsResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
