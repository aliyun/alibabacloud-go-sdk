// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmAccessProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLlmAccessProfilesResponseBody
	GetCode() *string
	SetData(v *ListLlmAccessProfilesResponseBodyData) *ListLlmAccessProfilesResponseBody
	GetData() *ListLlmAccessProfilesResponseBodyData
	SetHttpStatusCode(v int32) *ListLlmAccessProfilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListLlmAccessProfilesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListLlmAccessProfilesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListLlmAccessProfilesResponseBody
	GetRequestId() *string
}

type ListLlmAccessProfilesResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListLlmAccessProfilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// ED56B723-5802-5C32-865F-6E20B06D2455
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLlmAccessProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLlmAccessProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLlmAccessProfilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLlmAccessProfilesResponseBody) GetData() *ListLlmAccessProfilesResponseBodyData {
	return s.Data
}

func (s *ListLlmAccessProfilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLlmAccessProfilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLlmAccessProfilesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListLlmAccessProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLlmAccessProfilesResponseBody) SetCode(v string) *ListLlmAccessProfilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBody) SetData(v *ListLlmAccessProfilesResponseBodyData) *ListLlmAccessProfilesResponseBody {
	s.Data = v
	return s
}

func (s *ListLlmAccessProfilesResponseBody) SetHttpStatusCode(v int32) *ListLlmAccessProfilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBody) SetMessage(v string) *ListLlmAccessProfilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBody) SetParams(v []*string) *ListLlmAccessProfilesResponseBody {
	s.Params = v
	return s
}

func (s *ListLlmAccessProfilesResponseBody) SetRequestId(v string) *ListLlmAccessProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLlmAccessProfilesResponseBodyData struct {
	LlmAccessProfiles []*ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles `json:"LlmAccessProfiles,omitempty" xml:"LlmAccessProfiles,omitempty" type:"Repeated"`
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
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLlmAccessProfilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLlmAccessProfilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLlmAccessProfilesResponseBodyData) GetLlmAccessProfiles() []*ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles {
	return s.LlmAccessProfiles
}

func (s *ListLlmAccessProfilesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLlmAccessProfilesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLlmAccessProfilesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLlmAccessProfilesResponseBodyData) SetLlmAccessProfiles(v []*ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) *ListLlmAccessProfilesResponseBodyData {
	s.LlmAccessProfiles = v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyData) SetPageNumber(v int32) *ListLlmAccessProfilesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyData) SetPageSize(v int32) *ListLlmAccessProfilesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyData) SetTotalCount(v int32) *ListLlmAccessProfilesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyData) Validate() error {
	if s.LlmAccessProfiles != nil {
		for _, item := range s.LlmAccessProfiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// 1747620752000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 8fe2b8b1cdd446318610ccbc70bcfff0
	InstanceId *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Profile    *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// example:
	//
	// 1768788798
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) String() string {
	return dara.Prettify(s)
}

func (s ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) GoString() string {
	return s.String()
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) GetProfile() *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile {
	return s.Profile
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) SetAccessProfileId(v string) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles {
	s.AccessProfileId = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) SetCreatedTime(v int64) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles {
	s.CreatedTime = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) SetInstanceId(v string) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles {
	s.InstanceId = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) SetProfile(v *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles {
	s.Profile = v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) SetUpdatedTime(v int64) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles {
	s.UpdatedTime = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile struct {
	// example:
	//
	// akm-93929110-d7c1-4014-8678-613aacd58fa2
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// api.llm.enpoint.example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
}

func (s ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) String() string {
	return dara.Prettify(s)
}

func (s ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) GoString() string {
	return s.String()
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) GetName() *string {
	return s.Name
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) SetApiKey(v string) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile {
	s.ApiKey = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) SetEndpoint(v string) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile {
	s.Endpoint = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) SetName(v string) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile {
	s.Name = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) SetNluAccessType(v string) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile {
	s.NluAccessType = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) SetNluEngine(v string) *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile {
	s.NluEngine = &v
	return s
}

func (s *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile) Validate() error {
	return dara.Validate(s)
}
