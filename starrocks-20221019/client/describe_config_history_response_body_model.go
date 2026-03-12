// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeConfigHistoryResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*DescribeConfigHistoryResponseBodyData) *DescribeConfigHistoryResponseBody
	GetData() []*DescribeConfigHistoryResponseBodyData
	SetErrCode(v string) *DescribeConfigHistoryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeConfigHistoryResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeConfigHistoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeConfigHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeConfigHistoryResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeConfigHistoryResponseBody
	GetTotal() *int32
}

type DescribeConfigHistoryResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*DescribeConfigHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeConfigHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeConfigHistoryResponseBody) GetData() []*DescribeConfigHistoryResponseBodyData {
	return s.Data
}

func (s *DescribeConfigHistoryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeConfigHistoryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeConfigHistoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeConfigHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfigHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeConfigHistoryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeConfigHistoryResponseBody) SetAccessDeniedDetail(v string) *DescribeConfigHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetData(v []*DescribeConfigHistoryResponseBodyData) *DescribeConfigHistoryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetErrCode(v string) *DescribeConfigHistoryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetErrMessage(v string) *DescribeConfigHistoryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetHttpStatusCode(v int32) *DescribeConfigHistoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetRequestId(v string) *DescribeConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetSuccess(v bool) *DescribeConfigHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetTotal(v int32) *DescribeConfigHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) Validate() error {
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

type DescribeConfigHistoryResponseBodyData struct {
	ConfigHistoryEffectDetails []*DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails `json:"ConfigHistoryEffectDetails,omitempty" xml:"ConfigHistoryEffectDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 6838
	ConfigHistoryId *string                                                `json:"ConfigHistoryId,omitempty" xml:"ConfigHistoryId,omitempty"`
	ConfigMementos  []*DescribeConfigHistoryResponseBodyDataConfigMementos `json:"ConfigMementos,omitempty" xml:"ConfigMementos,omitempty" type:"Repeated"`
	// example:
	//
	// effective
	EffectStatus *string `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// example:
	//
	// true
	Effected *bool `json:"Effected,omitempty" xml:"Effected,omitempty"`
	// example:
	//
	// 1742178604000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 149920818483****
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// example:
	//
	// Test.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// false
	Rollback *bool `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
}

func (s DescribeConfigHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBodyData) GetConfigHistoryEffectDetails() []*DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails {
	return s.ConfigHistoryEffectDetails
}

func (s *DescribeConfigHistoryResponseBodyData) GetConfigHistoryId() *string {
	return s.ConfigHistoryId
}

func (s *DescribeConfigHistoryResponseBodyData) GetConfigMementos() []*DescribeConfigHistoryResponseBodyDataConfigMementos {
	return s.ConfigMementos
}

func (s *DescribeConfigHistoryResponseBodyData) GetEffectStatus() *string {
	return s.EffectStatus
}

func (s *DescribeConfigHistoryResponseBodyData) GetEffected() *bool {
	return s.Effected
}

func (s *DescribeConfigHistoryResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeConfigHistoryResponseBodyData) GetOperatorId() *string {
	return s.OperatorId
}

func (s *DescribeConfigHistoryResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *DescribeConfigHistoryResponseBodyData) GetRollback() *bool {
	return s.Rollback
}

func (s *DescribeConfigHistoryResponseBodyData) SetConfigHistoryEffectDetails(v []*DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) *DescribeConfigHistoryResponseBodyData {
	s.ConfigHistoryEffectDetails = v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetConfigHistoryId(v string) *DescribeConfigHistoryResponseBodyData {
	s.ConfigHistoryId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetConfigMementos(v []*DescribeConfigHistoryResponseBodyDataConfigMementos) *DescribeConfigHistoryResponseBodyData {
	s.ConfigMementos = v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetEffectStatus(v string) *DescribeConfigHistoryResponseBodyData {
	s.EffectStatus = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetEffected(v bool) *DescribeConfigHistoryResponseBodyData {
	s.Effected = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetGmtCreate(v int64) *DescribeConfigHistoryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetOperatorId(v string) *DescribeConfigHistoryResponseBodyData {
	s.OperatorId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetReason(v string) *DescribeConfigHistoryResponseBodyData {
	s.Reason = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) SetRollback(v bool) *DescribeConfigHistoryResponseBodyData {
	s.Rollback = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyData) Validate() error {
	if s.ConfigHistoryEffectDetails != nil {
		for _, item := range s.ConfigHistoryEffectDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConfigMementos != nil {
		for _, item := range s.ConfigMementos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails struct {
	// example:
	//
	// effective
	EffectStatus *string `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// example:
	//
	// ng-e6e15d2cdefdb38c
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// 10000367486
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) GetEffectStatus() *string {
	return s.EffectStatus
}

func (s *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) SetEffectStatus(v string) *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails {
	s.EffectStatus = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) SetNodeGroupId(v string) *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) SetNodeId(v string) *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails {
	s.NodeId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigHistoryEffectDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigHistoryResponseBodyDataConfigMementos struct {
	// example:
	//
	// MODIFY
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// 4
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// example:
	//
	// 3
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
	// example:
	//
	// create_tablet_worker_count
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// example:
	//
	// FE
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
}

func (s DescribeConfigHistoryResponseBodyDataConfigMementos) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryResponseBodyDataConfigMementos) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) GetAction() *string {
	return s.Action
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) GetAfter() *string {
	return s.After
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) GetBefore() *string {
	return s.Before
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) GetConfigType() *string {
	return s.ConfigType
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) SetAction(v string) *DescribeConfigHistoryResponseBodyDataConfigMementos {
	s.Action = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) SetAfter(v string) *DescribeConfigHistoryResponseBodyDataConfigMementos {
	s.After = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) SetBefore(v string) *DescribeConfigHistoryResponseBodyDataConfigMementos {
	s.Before = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) SetConfigKey(v string) *DescribeConfigHistoryResponseBodyDataConfigMementos {
	s.ConfigKey = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) SetConfigType(v string) *DescribeConfigHistoryResponseBodyDataConfigMementos {
	s.ConfigType = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyDataConfigMementos) Validate() error {
	return dara.Validate(s)
}
