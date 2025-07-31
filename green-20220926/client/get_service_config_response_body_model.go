// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetServiceConfigResponseBody
	GetCode() *int32
	SetData(v *GetServiceConfigResponseBodyData) *GetServiceConfigResponseBody
	GetData() *GetServiceConfigResponseBodyData
	SetMsg(v string) *GetServiceConfigResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetServiceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceConfigResponseBody
	GetSuccess() *bool
}

type GetServiceConfigResponseBody struct {
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetServiceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetServiceConfigResponseBody) GetData() *GetServiceConfigResponseBodyData {
	return s.Data
}

func (s *GetServiceConfigResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetServiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceConfigResponseBody) SetCode(v int32) *GetServiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceConfigResponseBody) SetData(v *GetServiceConfigResponseBodyData) *GetServiceConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceConfigResponseBody) SetMsg(v string) *GetServiceConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *GetServiceConfigResponseBody) SetRequestId(v string) *GetServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceConfigResponseBody) SetSuccess(v bool) *GetServiceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceConfigResponseBodyData struct {
	CustomServiceConf *GetServiceConfigResponseBodyDataCustomServiceConf `json:"CustomServiceConf,omitempty" xml:"CustomServiceConf,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-06 03:07:44
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// UID。
	//
	// example:
	//
	// 165379****31937
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetServiceConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponseBodyData) GetCustomServiceConf() *GetServiceConfigResponseBodyDataCustomServiceConf {
	return s.CustomServiceConf
}

func (s *GetServiceConfigResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetServiceConfigResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetServiceConfigResponseBodyData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetServiceConfigResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetServiceConfigResponseBodyData) SetCustomServiceConf(v *GetServiceConfigResponseBodyDataCustomServiceConf) *GetServiceConfigResponseBodyData {
	s.CustomServiceConf = v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetGmtModified(v string) *GetServiceConfigResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetResourceType(v string) *GetServiceConfigResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetServiceCode(v string) *GetServiceConfigResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetUid(v string) *GetServiceConfigResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetServiceConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetServiceConfigResponseBodyDataCustomServiceConf struct {
	KeywordFilterLibs   []*string                                                             `json:"KeywordFilterLibs,omitempty" xml:"KeywordFilterLibs,omitempty" type:"Repeated"`
	KeywordHitLibs      []*string                                                             `json:"KeywordHitLibs,omitempty" xml:"KeywordHitLibs,omitempty" type:"Repeated"`
	ManualMachineConfig *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig `json:"ManualMachineConfig,omitempty" xml:"ManualMachineConfig,omitempty" type:"Struct"`
	SimilarTextHitLibs  []*string                                                             `json:"SimilarTextHitLibs,omitempty" xml:"SimilarTextHitLibs,omitempty" type:"Repeated"`
}

func (s GetServiceConfigResponseBodyDataCustomServiceConf) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfigResponseBodyDataCustomServiceConf) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) GetKeywordFilterLibs() []*string {
	return s.KeywordFilterLibs
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) GetKeywordHitLibs() []*string {
	return s.KeywordHitLibs
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) GetManualMachineConfig() *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig {
	return s.ManualMachineConfig
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) GetSimilarTextHitLibs() []*string {
	return s.SimilarTextHitLibs
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) SetKeywordFilterLibs(v []*string) *GetServiceConfigResponseBodyDataCustomServiceConf {
	s.KeywordFilterLibs = v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) SetKeywordHitLibs(v []*string) *GetServiceConfigResponseBodyDataCustomServiceConf {
	s.KeywordHitLibs = v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) SetManualMachineConfig(v *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) *GetServiceConfigResponseBodyDataCustomServiceConf {
	s.ManualMachineConfig = v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) SetSimilarTextHitLibs(v []*string) *GetServiceConfigResponseBodyDataCustomServiceConf {
	s.SimilarTextHitLibs = v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) Validate() error {
	return dara.Validate(s)
}

type GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig struct {
	AuditRiskLevels []*string `json:"AuditRiskLevels,omitempty" xml:"AuditRiskLevels,omitempty" type:"Repeated"`
	CallbackId      *int64    `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	Enable          *bool     `json:"Enable,omitempty" xml:"Enable,omitempty"`
	ManualService   *string   `json:"ManualService,omitempty" xml:"ManualService,omitempty"`
}

func (s GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) GetAuditRiskLevels() []*string {
	return s.AuditRiskLevels
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) GetCallbackId() *int64 {
	return s.CallbackId
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) GetEnable() *bool {
	return s.Enable
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) GetManualService() *string {
	return s.ManualService
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) SetAuditRiskLevels(v []*string) *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig {
	s.AuditRiskLevels = v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) SetCallbackId(v int64) *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig {
	s.CallbackId = &v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) SetEnable(v bool) *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig {
	s.Enable = &v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) SetManualService(v string) *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig {
	s.ManualService = &v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConfManualMachineConfig) Validate() error {
	return dara.Validate(s)
}
