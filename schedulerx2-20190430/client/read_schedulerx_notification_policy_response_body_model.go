// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxNotificationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *ReadSchedulerxNotificationPolicyResponseBody
	GetAccessDeniedDetail() *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail
	SetCode(v int32) *ReadSchedulerxNotificationPolicyResponseBody
	GetCode() *int32
	SetData(v *ReadSchedulerxNotificationPolicyResponseBodyData) *ReadSchedulerxNotificationPolicyResponseBody
	GetData() *ReadSchedulerxNotificationPolicyResponseBodyData
	SetMessage(v string) *ReadSchedulerxNotificationPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadSchedulerxNotificationPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadSchedulerxNotificationPolicyResponseBody
	GetSuccess() *bool
}

type ReadSchedulerxNotificationPolicyResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// *
	Data *ReadSchedulerxNotificationPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// unknown exception occurred
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadSchedulerxNotificationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) GetAccessDeniedDetail() *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) GetData() *ReadSchedulerxNotificationPolicyResponseBodyData {
	return s.Data
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) SetAccessDeniedDetail(v *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *ReadSchedulerxNotificationPolicyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) SetCode(v int32) *ReadSchedulerxNotificationPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) SetData(v *ReadSchedulerxNotificationPolicyResponseBodyData) *ReadSchedulerxNotificationPolicyResponseBody {
	s.Data = v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) SetMessage(v string) *ReadSchedulerxNotificationPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) SetRequestId(v string) *ReadSchedulerxNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) SetSuccess(v bool) *ReadSchedulerxNotificationPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail struct {
	// The authentication operation.
	//
	// example:
	//
	// edas:ReadSchedulerxNotificationPolicy
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The principal name.
	//
	// example:
	//
	// 209312833131416xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The account of the principal.
	//
	// example:
	//
	// 1827811800526xxx
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The principal type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQFn/cLPZ/3Cz0YxQkZBMjVGLTY0REUtNTlGNS05NzUwLTgyMUE4M0MwMTFDRQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The permission denial type.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ReadSchedulerxNotificationPolicyResponseBodyData struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// When data that matches the query conditions has not been fully retrieved, the server returns nextToken. You can then use nextToken to continue retrieving the remaining data.
	//
	// example:
	//
	// O39nXKu5XafATl3/cJjSJw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The data records.
	Records []*ReadSchedulerxNotificationPolicyResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 42
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ReadSchedulerxNotificationPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxNotificationPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) GetRecords() []*ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	return s.Records
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) SetMaxResults(v int32) *ReadSchedulerxNotificationPolicyResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) SetNextToken(v string) *ReadSchedulerxNotificationPolicyResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) SetRecords(v []*ReadSchedulerxNotificationPolicyResponseBodyDataRecords) *ReadSchedulerxNotificationPolicyResponseBodyData {
	s.Records = v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) SetTotal(v int64) *ReadSchedulerxNotificationPolicyResponseBodyData {
	s.Total = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadSchedulerxNotificationPolicyResponseBodyDataRecords struct {
	// The configuration of the notification policy.
	//
	// example:
	//
	// {
	//
	//   "isUnifiedSetting": false,
	//
	//   "timezone": "GMT+8",
	//
	//   "webhookIsAtAll": "false",
	//
	//   "timeRanges": {
	//
	//     "webhook": [
	//
	//       {
	//
	//         "startTime": "08:00",
	//
	//         "endTime": "18:00",
	//
	//         "daysOfWeek": [
	//
	//           1,
	//
	//           2,
	//
	//           3,
	//
	//           4,
	//
	//           5
	//
	//         ]
	//
	//       }
	//
	//     ],
	//
	//     "sms": [
	//
	//       {
	//
	//         "startTime": "08:00",
	//
	//         "endTime": "18:00",
	//
	//         "daysOfWeek": [
	//
	//           1,
	//
	//           2,
	//
	//           3,
	//
	//           4,
	//
	//           5
	//
	//         ]
	//
	//       }
	//
	//     ],
	//
	//     "mail": [
	//
	//       {
	//
	//         "startTime": "08:00",
	//
	//         "endTime": "18:00",
	//
	//         "daysOfWeek": [
	//
	//           1,
	//
	//           2,
	//
	//           3,
	//
	//           4,
	//
	//           5
	//
	//         ]
	//
	//       }
	//
	//     ],
	//
	//     "phone": [
	//
	//       {
	//
	//         "startTime": "08:00",
	//
	//         "endTime": "18:00",
	//
	//         "daysOfWeek": [
	//
	//           1,
	//
	//           2,
	//
	//           3,
	//
	//           4,
	//
	//           5
	//
	//         ]
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	// }
	ChannelTimeRange *string `json:"ChannelTimeRange,omitempty" xml:"ChannelTimeRange,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-09-17 11:21:01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 201576653956616970
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the notification policy.
	//
	// example:
	//
	// Monday-Friday only
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// test-weekdays
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The list of applications associated with the notification policy.
	ReferenceApps []*ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps `json:"ReferenceApps,omitempty" xml:"ReferenceApps,omitempty" type:"Repeated"`
	// The update time.
	//
	// example:
	//
	// 2025-09-17 11:21:01
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The updater.
	//
	// example:
	//
	// 1144881807903942
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
}

func (s ReadSchedulerxNotificationPolicyResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetChannelTimeRange() *string {
	return s.ChannelTimeRange
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetReferenceApps() []*ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps {
	return s.ReferenceApps
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) GetUpdater() *string {
	return s.Updater
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetChannelTimeRange(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.ChannelTimeRange = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetCreateTime(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.CreateTime = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetCreator(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetDescription(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetPolicyName(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.PolicyName = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetReferenceApps(v []*ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.ReferenceApps = v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetUpdateTime(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.UpdateTime = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) SetUpdater(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecords {
	s.Updater = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecords) Validate() error {
	if s.ReferenceApps != nil {
		for _, item := range s.ReferenceApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps struct {
	// The ID of the task group.
	//
	// example:
	//
	// 123
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// test-app
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// Notification strategy testing namespace
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The unique identifier of the namespace.
	//
	// example:
	//
	// 18271388-aa16-4eab-9a6f-55f65d7e4391
	NamespaceUid *string `json:"NamespaceUid,omitempty" xml:"NamespaceUid,omitempty"`
}

func (s ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) GetGroupId() *string {
	return s.GroupId
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) GetNamespaceUid() *string {
	return s.NamespaceUid
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) SetAppGroupId(v int64) *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps {
	s.AppGroupId = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) SetGroupId(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps {
	s.GroupId = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) SetNamespaceName(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps {
	s.NamespaceName = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) SetNamespaceUid(v string) *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps {
	s.NamespaceUid = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponseBodyDataRecordsReferenceApps) Validate() error {
	return dara.Validate(s)
}
