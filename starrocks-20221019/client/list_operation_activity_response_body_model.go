// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationActivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListOperationActivityResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*ListOperationActivityResponseBodyData) *ListOperationActivityResponseBody
	GetData() []*ListOperationActivityResponseBodyData
	SetErrCode(v string) *ListOperationActivityResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListOperationActivityResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListOperationActivityResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListOperationActivityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOperationActivityResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListOperationActivityResponseBody
	GetTotal() *int32
}

type ListOperationActivityResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*ListOperationActivityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
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
	// 440
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOperationActivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationActivityResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationActivityResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListOperationActivityResponseBody) GetData() []*ListOperationActivityResponseBodyData {
	return s.Data
}

func (s *ListOperationActivityResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListOperationActivityResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListOperationActivityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOperationActivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationActivityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOperationActivityResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListOperationActivityResponseBody) SetAccessDeniedDetail(v string) *ListOperationActivityResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListOperationActivityResponseBody) SetData(v []*ListOperationActivityResponseBodyData) *ListOperationActivityResponseBody {
	s.Data = v
	return s
}

func (s *ListOperationActivityResponseBody) SetErrCode(v string) *ListOperationActivityResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListOperationActivityResponseBody) SetErrMessage(v string) *ListOperationActivityResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListOperationActivityResponseBody) SetHttpStatusCode(v int32) *ListOperationActivityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOperationActivityResponseBody) SetRequestId(v string) *ListOperationActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationActivityResponseBody) SetSuccess(v bool) *ListOperationActivityResponseBody {
	s.Success = &v
	return s
}

func (s *ListOperationActivityResponseBody) SetTotal(v int32) *ListOperationActivityResponseBody {
	s.Total = &v
	return s
}

func (s *ListOperationActivityResponseBody) Validate() error {
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

type ListOperationActivityResponseBodyData struct {
	// example:
	//
	// ac-8f3f7c4026e3****
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// COMPLETED
	ActivityStatus *string `json:"ActivityStatus,omitempty" xml:"ActivityStatus,omitempty"`
	// example:
	//
	// 1742178604000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// UpdateClusterStatusToModifyingConfigStatus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1742178604000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListOperationActivityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOperationActivityResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOperationActivityResponseBodyData) GetActivityId() *string {
	return s.ActivityId
}

func (s *ListOperationActivityResponseBodyData) GetActivityStatus() *string {
	return s.ActivityStatus
}

func (s *ListOperationActivityResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationActivityResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListOperationActivityResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationActivityResponseBodyData) SetActivityId(v string) *ListOperationActivityResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *ListOperationActivityResponseBodyData) SetActivityStatus(v string) *ListOperationActivityResponseBodyData {
	s.ActivityStatus = &v
	return s
}

func (s *ListOperationActivityResponseBodyData) SetEndTime(v int64) *ListOperationActivityResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListOperationActivityResponseBodyData) SetName(v string) *ListOperationActivityResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListOperationActivityResponseBodyData) SetStartTime(v int64) *ListOperationActivityResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListOperationActivityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
