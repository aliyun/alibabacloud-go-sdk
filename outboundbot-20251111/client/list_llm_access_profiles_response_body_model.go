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
	SetSuccess(v bool) *ListLlmAccessProfilesResponseBody
	GetSuccess() *bool
}

type ListLlmAccessProfilesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListLlmAccessProfilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *ListLlmAccessProfilesResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *ListLlmAccessProfilesResponseBody) SetSuccess(v bool) *ListLlmAccessProfilesResponseBody {
	s.Success = &v
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
	// The list of third-party language model access configurations.
	LlmAccessProfiles []*ListLlmAccessProfilesResponseBodyDataLlmAccessProfiles `json:"LlmAccessProfiles,omitempty" xml:"LlmAccessProfiles,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 15
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
	// The ID of the access configuration.
	//
	// example:
	//
	// 5dcc2dc6-57b4-488e-b6c8-6115cbc2a21e
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1779700982000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 7a8ef16a-c1a6-4c97-acbe-3e53521beb39
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language model access configuration.
	Profile *ListLlmAccessProfilesResponseBodyDataLlmAccessProfilesProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// The update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1779700982000
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
	// The API key.
	//
	// example:
	//
	// example-api-key
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The endpoint URL of the model service.
	//
	// example:
	//
	// wss://example.com/api/realtime
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// Sample configuration
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the access configuration.
	//
	// example:
	//
	// Provided
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// The type of the natural language understanding (NLU) engine.
	//
	// example:
	//
	// PROMPTS
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
