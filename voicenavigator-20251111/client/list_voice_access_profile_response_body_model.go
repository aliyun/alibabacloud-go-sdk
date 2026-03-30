// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVoiceAccessProfileResponseBody
	GetCode() *string
	SetData(v *ListVoiceAccessProfileResponseBodyData) *ListVoiceAccessProfileResponseBody
	GetData() *ListVoiceAccessProfileResponseBodyData
	SetHttpStatusCode(v int32) *ListVoiceAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVoiceAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListVoiceAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListVoiceAccessProfileResponseBody
	GetRequestId() *string
}

type ListVoiceAccessProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListVoiceAccessProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// C1788610-93C9-574C-AF2D-C053A50C3894
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVoiceAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVoiceAccessProfileResponseBody) GetData() *ListVoiceAccessProfileResponseBodyData {
	return s.Data
}

func (s *ListVoiceAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVoiceAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVoiceAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListVoiceAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoiceAccessProfileResponseBody) SetCode(v string) *ListVoiceAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBody) SetData(v *ListVoiceAccessProfileResponseBodyData) *ListVoiceAccessProfileResponseBody {
	s.Data = v
	return s
}

func (s *ListVoiceAccessProfileResponseBody) SetHttpStatusCode(v int32) *ListVoiceAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBody) SetMessage(v string) *ListVoiceAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBody) SetParams(v []*string) *ListVoiceAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *ListVoiceAccessProfileResponseBody) SetRequestId(v string) *ListVoiceAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVoiceAccessProfileResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 20
	TotalCount          *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoiceAccessProfiles []*ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles `json:"VoiceAccessProfiles,omitempty" xml:"VoiceAccessProfiles,omitempty" type:"Repeated"`
}

func (s ListVoiceAccessProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfileResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceAccessProfileResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceAccessProfileResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVoiceAccessProfileResponseBodyData) GetVoiceAccessProfiles() []*ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	return s.VoiceAccessProfiles
}

func (s *ListVoiceAccessProfileResponseBodyData) SetPageNumber(v int32) *ListVoiceAccessProfileResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyData) SetPageSize(v int32) *ListVoiceAccessProfileResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyData) SetTotalCount(v int32) *ListVoiceAccessProfileResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyData) SetVoiceAccessProfiles(v []*ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) *ListVoiceAccessProfileResponseBodyData {
	s.VoiceAccessProfiles = v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyData) Validate() error {
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

type ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string   `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	Capabilities    []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// example:
	//
	// 1747620752000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VOLC
	NlsEngine     *string                                                           `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	NlsEngineName *string                                                           `json:"NlsEngineName,omitempty" xml:"NlsEngineName,omitempty"`
	Profile       *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// example:
	//
	// 1747620752000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetNlsEngineName() *string {
	return s.NlsEngineName
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetProfile() *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	return s.Profile
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetAccessProfileId(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.AccessProfileId = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetCapabilities(v []*string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.Capabilities = v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetCreatedTime(v int64) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.CreatedTime = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetInstanceId(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.InstanceId = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetNlsEngine(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.NlsEngine = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetNlsEngineName(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.NlsEngineName = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetProfile(v *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.Profile = v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) SetUpdatedTime(v int64) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles {
	s.UpdatedTime = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfiles) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile struct {
	// example:
	//
	// HwRnTXgwnQOlsj68URDS5_VMm4Wtapq9
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// sk-12341e259b1049e8872b47981e545f78
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// c0358c6e51c1013b446fdeb21a3a1234
	ApiSecret *string `json:"ApiSecret,omitempty" xml:"ApiSecret,omitempty"`
	// example:
	//
	// 5b123bfb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2541370123
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AsrAppKey *string `json:"AsrAppKey,omitempty" xml:"AsrAppKey,omitempty"`
	TtsApiKey *string `json:"TtsApiKey,omitempty" xml:"TtsApiKey,omitempty"`
}

func (s ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GetApiSecret() *string {
	return s.ApiSecret
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GetAppId() *string {
	return s.AppId
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GetAppKey() *string {
	return s.AppKey
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GetAsrAppKey() *string {
	return s.AsrAppKey
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) GetTtsApiKey() *string {
	return s.TtsApiKey
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) SetAccessKey(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	s.AccessKey = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) SetApiKey(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	s.ApiKey = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) SetApiSecret(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	s.ApiSecret = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) SetAppId(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	s.AppId = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) SetAppKey(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	s.AppKey = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) SetAsrAppKey(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	s.AsrAppKey = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) SetTtsApiKey(v string) *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile {
	s.TtsApiKey = &v
	return s
}

func (s *ListVoiceAccessProfileResponseBodyDataVoiceAccessProfilesProfile) Validate() error {
	return dara.Validate(s)
}
