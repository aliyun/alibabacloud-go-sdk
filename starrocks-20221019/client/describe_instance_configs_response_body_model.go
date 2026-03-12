// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeInstanceConfigsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*DescribeInstanceConfigsResponseBodyData) *DescribeInstanceConfigsResponseBody
	GetData() []*DescribeInstanceConfigsResponseBodyData
	SetErrCode(v string) *DescribeInstanceConfigsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeInstanceConfigsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeInstanceConfigsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeInstanceConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceConfigsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeInstanceConfigsResponseBody
	GetTotal() *int32
}

type DescribeInstanceConfigsResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                    `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*DescribeInstanceConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 4
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeInstanceConfigsResponseBody) GetData() []*DescribeInstanceConfigsResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceConfigsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeInstanceConfigsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeInstanceConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceConfigsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeInstanceConfigsResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceConfigsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetData(v []*DescribeInstanceConfigsResponseBodyData) *DescribeInstanceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetErrCode(v string) *DescribeInstanceConfigsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetErrMessage(v string) *DescribeInstanceConfigsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetRequestId(v string) *DescribeInstanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetSuccess(v bool) *DescribeInstanceConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetTotal(v int32) *DescribeInstanceConfigsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) Validate() error {
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

type DescribeInstanceConfigsResponseBodyData struct {
	// example:
	//
	// true
	AllowModify *string `json:"AllowModify,omitempty" xml:"AllowModify,omitempty"`
	// example:
	//
	// enable_udf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// example:
	//
	// FE
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// example:
	//
	// false
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// example:
	//
	// false
	Custom *bool `json:"Custom,omitempty" xml:"Custom,omitempty"`
	// example:
	//
	// true
	DefaultNodeGroup *bool `json:"DefaultNodeGroup,omitempty" xml:"DefaultNodeGroup,omitempty"`
	// example:
	//
	// true
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// A boolean value to control whether to enable the synchronization of the tablet metadata. true indicates enabling synchronization, and false indicates disabling it.
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// ng_1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// example:
	//
	// s
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// true,false
	ValueRange *string `json:"ValueRange,omitempty" xml:"ValueRange,omitempty"`
	// example:
	//
	// INT
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeInstanceConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsResponseBodyData) GetAllowModify() *string {
	return s.AllowModify
}

func (s *DescribeInstanceConfigsResponseBodyData) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeInstanceConfigsResponseBodyData) GetConfigType() *string {
	return s.ConfigType
}

func (s *DescribeInstanceConfigsResponseBodyData) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeInstanceConfigsResponseBodyData) GetCustom() *bool {
	return s.Custom
}

func (s *DescribeInstanceConfigsResponseBodyData) GetDefaultNodeGroup() *bool {
	return s.DefaultNodeGroup
}

func (s *DescribeInstanceConfigsResponseBodyData) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeInstanceConfigsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceConfigsResponseBodyData) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *DescribeInstanceConfigsResponseBodyData) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeInstanceConfigsResponseBodyData) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeInstanceConfigsResponseBodyData) GetRestart() *bool {
	return s.Restart
}

func (s *DescribeInstanceConfigsResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *DescribeInstanceConfigsResponseBodyData) GetValueRange() *string {
	return s.ValueRange
}

func (s *DescribeInstanceConfigsResponseBodyData) GetValueType() *string {
	return s.ValueType
}

func (s *DescribeInstanceConfigsResponseBodyData) SetAllowModify(v string) *DescribeInstanceConfigsResponseBodyData {
	s.AllowModify = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetConfigKey(v string) *DescribeInstanceConfigsResponseBodyData {
	s.ConfigKey = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetConfigType(v string) *DescribeInstanceConfigsResponseBodyData {
	s.ConfigType = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetConfigValue(v string) *DescribeInstanceConfigsResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetCustom(v bool) *DescribeInstanceConfigsResponseBodyData {
	s.Custom = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetDefaultNodeGroup(v bool) *DescribeInstanceConfigsResponseBodyData {
	s.DefaultNodeGroup = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetDefaultValue(v string) *DescribeInstanceConfigsResponseBodyData {
	s.DefaultValue = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetDescription(v string) *DescribeInstanceConfigsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetDescriptionEn(v string) *DescribeInstanceConfigsResponseBodyData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetNodeGroupId(v string) *DescribeInstanceConfigsResponseBodyData {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetNodeGroupName(v string) *DescribeInstanceConfigsResponseBodyData {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetRestart(v bool) *DescribeInstanceConfigsResponseBodyData {
	s.Restart = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetUnit(v string) *DescribeInstanceConfigsResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetValueRange(v string) *DescribeInstanceConfigsResponseBodyData {
	s.ValueRange = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) SetValueType(v string) *DescribeInstanceConfigsResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
