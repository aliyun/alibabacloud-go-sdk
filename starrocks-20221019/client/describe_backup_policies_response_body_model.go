// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeBackupPoliciesResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*DescribeBackupPoliciesResponseBodyData) *DescribeBackupPoliciesResponseBody
	GetData() []*DescribeBackupPoliciesResponseBodyData
	SetErrCode(v string) *DescribeBackupPoliciesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupPoliciesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeBackupPoliciesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeBackupPoliciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupPoliciesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeBackupPoliciesResponseBody
	GetTotal() *int32
}

type DescribeBackupPoliciesResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*DescribeBackupPoliciesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
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
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeBackupPoliciesResponseBody) GetData() []*DescribeBackupPoliciesResponseBodyData {
	return s.Data
}

func (s *DescribeBackupPoliciesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupPoliciesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupPoliciesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBackupPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPoliciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupPoliciesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeBackupPoliciesResponseBody) SetAccessDeniedDetail(v string) *DescribeBackupPoliciesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetData(v []*DescribeBackupPoliciesResponseBodyData) *DescribeBackupPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetErrCode(v string) *DescribeBackupPoliciesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetErrMessage(v string) *DescribeBackupPoliciesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetHttpStatusCode(v int32) *DescribeBackupPoliciesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetRequestId(v string) *DescribeBackupPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetSuccess(v bool) *DescribeBackupPoliciesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetTotal(v int32) *DescribeBackupPoliciesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) Validate() error {
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

type DescribeBackupPoliciesResponseBodyData struct {
	// example:
	//
	// 7
	ExpireDays *int32 `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
	// example:
	//
	// 2
	Hour *string `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// c-96f3bc7f04b2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 30
	Minute *string `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 1
	PolicyId         *string  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s DescribeBackupPoliciesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesResponseBodyData) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *DescribeBackupPoliciesResponseBodyData) GetHour() *string {
	return s.Hour
}

func (s *DescribeBackupPoliciesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupPoliciesResponseBodyData) GetMinute() *string {
	return s.Minute
}

func (s *DescribeBackupPoliciesResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeBackupPoliciesResponseBodyData) GetRecurrenceValues() []*int32 {
	return s.RecurrenceValues
}

func (s *DescribeBackupPoliciesResponseBodyData) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *DescribeBackupPoliciesResponseBodyData) SetExpireDays(v int32) *DescribeBackupPoliciesResponseBodyData {
	s.ExpireDays = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyData) SetHour(v string) *DescribeBackupPoliciesResponseBodyData {
	s.Hour = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyData) SetInstanceId(v string) *DescribeBackupPoliciesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyData) SetMinute(v string) *DescribeBackupPoliciesResponseBodyData {
	s.Minute = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyData) SetPolicyId(v string) *DescribeBackupPoliciesResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyData) SetRecurrenceValues(v []*int32) *DescribeBackupPoliciesResponseBodyData {
	s.RecurrenceValues = v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyData) SetTimeoutSeconds(v int32) *DescribeBackupPoliciesResponseBodyData {
	s.TimeoutSeconds = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
