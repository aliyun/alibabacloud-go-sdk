// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceAccessProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVoiceAccessProfilesResponseBody
	GetCode() *string
	SetData(v *ListVoiceAccessProfilesResponseBodyData) *ListVoiceAccessProfilesResponseBody
	GetData() *ListVoiceAccessProfilesResponseBodyData
	SetHttpStatusCode(v int32) *ListVoiceAccessProfilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVoiceAccessProfilesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListVoiceAccessProfilesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListVoiceAccessProfilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVoiceAccessProfilesResponseBody
	GetSuccess() *bool
}

type ListVoiceAccessProfilesResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ListVoiceAccessProfilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=outb001,.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVoiceAccessProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVoiceAccessProfilesResponseBody) GetData() *ListVoiceAccessProfilesResponseBodyData {
	return s.Data
}

func (s *ListVoiceAccessProfilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVoiceAccessProfilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVoiceAccessProfilesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListVoiceAccessProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoiceAccessProfilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVoiceAccessProfilesResponseBody) SetCode(v string) *ListVoiceAccessProfilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBody) SetData(v *ListVoiceAccessProfilesResponseBodyData) *ListVoiceAccessProfilesResponseBody {
	s.Data = v
	return s
}

func (s *ListVoiceAccessProfilesResponseBody) SetHttpStatusCode(v int32) *ListVoiceAccessProfilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBody) SetMessage(v string) *ListVoiceAccessProfilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBody) SetParams(v []*string) *ListVoiceAccessProfilesResponseBody {
	s.Params = v
	return s
}

func (s *ListVoiceAccessProfilesResponseBody) SetRequestId(v string) *ListVoiceAccessProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBody) SetSuccess(v bool) *ListVoiceAccessProfilesResponseBody {
	s.Success = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVoiceAccessProfilesResponseBodyData struct {
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the conditions.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The data list.
	VoiceAccessProfiles []*ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles `json:"VoiceAccessProfiles,omitempty" xml:"VoiceAccessProfiles,omitempty" type:"Repeated"`
}

func (s ListVoiceAccessProfilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfilesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceAccessProfilesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceAccessProfilesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVoiceAccessProfilesResponseBodyData) GetVoiceAccessProfiles() []*ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	return s.VoiceAccessProfiles
}

func (s *ListVoiceAccessProfilesResponseBodyData) SetPageNumber(v int32) *ListVoiceAccessProfilesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyData) SetPageSize(v int32) *ListVoiceAccessProfilesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyData) SetTotalCount(v int32) *ListVoiceAccessProfilesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyData) SetVoiceAccessProfiles(v []*ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) *ListVoiceAccessProfilesResponseBodyData {
	s.VoiceAccessProfiles = v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyData) Validate() error {
	if s.VoiceAccessProfiles != nil {
		for _, item := range s.VoiceAccessProfiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles struct {
	// The access profile ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// The list of supported capabilities.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The creation time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1735660800000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The voice engine vendor. Valid values:
	//
	// - BAILIAN: Bailian.
	//
	// - VOLC: Doubao.
	//
	// - IFLYTEK: iFLYTEK.
	//
	// - TENCENT: Tencent.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The display name of the voice engine vendor.
	//
	// example:
	//
	// Bailian
	NlsEngineName *string `json:"NlsEngineName,omitempty" xml:"NlsEngineName,omitempty"`
	// The configuration.
	Profile *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// The update time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1735660800000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetNlsEngineName() *string {
	return s.NlsEngineName
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetProfile() *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	return s.Profile
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetAccessProfileId(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.AccessProfileId = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetCapabilities(v []*string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.Capabilities = v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetCreatedTime(v int64) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.CreatedTime = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetInstanceId(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.InstanceId = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetNlsEngine(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.NlsEngine = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetNlsEngineName(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.NlsEngineName = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetProfile(v *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.Profile = v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) SetUpdatedTime(v int64) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles {
	s.UpdatedTime = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfiles) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile struct {
	// Required when NlsEngine=VOLC.
	//
	// example:
	//
	// ****
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// Required when NlsEngine=IFLYTEK or NlsEngine=BAILIAN.
	//
	// example:
	//
	// a9872e2342952e248727798f642936c7
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// Required when NlsEngine=IFLYTEK.
	//
	// example:
	//
	// c0358c6e51c1013b446fdeb21a3a5d2e
	ApiSecret *string `json:"ApiSecret,omitempty" xml:"ApiSecret,omitempty"`
	// Required when NlsEngine=IFLYTEK or NlsEngine=TENCENT.
	//
	// example:
	//
	// 9479688350
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Required when NlsEngine=VOLC.
	//
	// example:
	//
	// DW0yKRHQEe1nAd8c
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Not currently in use.
	//
	// example:
	//
	// Not currently in use
	AsrAppKey *string `json:"AsrAppKey,omitempty" xml:"AsrAppKey,omitempty"`
	// Required when NlsEngine=TENCENT.
	//
	// example:
	//
	// sci_r3b3e62udqcujnkerrorqztnpu
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// Required when NlsEngine=TENCENT.
	//
	// example:
	//
	// y5MZfFdW6yBZgJdKonHZBA
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// Not currently in use.
	//
	// example:
	//
	// Not currently in use
	TtsApiKey *string `json:"TtsApiKey,omitempty" xml:"TtsApiKey,omitempty"`
}

func (s ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetApiSecret() *string {
	return s.ApiSecret
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetAppId() *string {
	return s.AppId
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetAppKey() *string {
	return s.AppKey
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetAsrAppKey() *string {
	return s.AsrAppKey
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetSecretId() *string {
	return s.SecretId
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetSecretKey() *string {
	return s.SecretKey
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) GetTtsApiKey() *string {
	return s.TtsApiKey
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetAccessKey(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.AccessKey = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetApiKey(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.ApiKey = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetApiSecret(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.ApiSecret = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetAppId(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.AppId = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetAppKey(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.AppKey = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetAsrAppKey(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.AsrAppKey = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetSecretId(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.SecretId = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetSecretKey(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.SecretKey = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) SetTtsApiKey(v string) *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile {
	s.TtsApiKey = &v
	return s
}

func (s *ListVoiceAccessProfilesResponseBodyDataVoiceAccessProfilesProfile) Validate() error {
	return dara.Validate(s)
}
