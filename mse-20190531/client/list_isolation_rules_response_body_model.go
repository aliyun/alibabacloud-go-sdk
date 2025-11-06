// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIsolationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListIsolationRulesResponseBody
	GetCode() *int32
	SetData(v *ListIsolationRulesResponseBodyData) *ListIsolationRulesResponseBody
	GetData() *ListIsolationRulesResponseBodyData
	SetHttpStatusCode(v int32) *ListIsolationRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIsolationRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIsolationRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIsolationRulesResponseBody
	GetSuccess() *bool
}

type ListIsolationRulesResponseBody struct {
	// example:
	//
	// 200
	Code *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListIsolationRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIsolationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIsolationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListIsolationRulesResponseBody) GetData() *ListIsolationRulesResponseBodyData {
	return s.Data
}

func (s *ListIsolationRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIsolationRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIsolationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIsolationRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIsolationRulesResponseBody) SetCode(v int32) *ListIsolationRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIsolationRulesResponseBody) SetData(v *ListIsolationRulesResponseBodyData) *ListIsolationRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListIsolationRulesResponseBody) SetHttpStatusCode(v int32) *ListIsolationRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIsolationRulesResponseBody) SetMessage(v string) *ListIsolationRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIsolationRulesResponseBody) SetRequestId(v string) *ListIsolationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIsolationRulesResponseBody) SetSuccess(v bool) *ListIsolationRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListIsolationRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIsolationRulesResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*ListIsolationRulesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListIsolationRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIsolationRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIsolationRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIsolationRulesResponseBodyData) GetResult() []*ListIsolationRulesResponseBodyDataResult {
	return s.Result
}

func (s *ListIsolationRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListIsolationRulesResponseBodyData) SetPageNumber(v int32) *ListIsolationRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIsolationRulesResponseBodyData) SetPageSize(v int32) *ListIsolationRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIsolationRulesResponseBodyData) SetResult(v []*ListIsolationRulesResponseBodyDataResult) *ListIsolationRulesResponseBodyData {
	s.Result = v
	return s
}

func (s *ListIsolationRulesResponseBodyData) SetTotalSize(v int32) *ListIsolationRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListIsolationRulesResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIsolationRulesResponseBodyDataResult struct {
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// "{\\"appName\\":\\"spring-cloud-a\\",\\"fallbackBehavior\\":{\\"webFallbackMode\\":0,\\"webRespContentType\\":0,\\"webRespMessage\\":\\"Blocked\\",\\"webRespStatusCode\\":429},\\"id\\":977,\\"name\\":\\"Fallback\\",\\"namespace\\":\\"default\\",\\"resourceClassification\\":1}"
	FallbackObject *string `json:"FallbackObject,omitempty" xml:"FallbackObject,omitempty"`
	LimitApp       *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListIsolationRulesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListIsolationRulesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *ListIsolationRulesResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *ListIsolationRulesResponseBodyDataResult) GetEnable() *bool {
	return s.Enable
}

func (s *ListIsolationRulesResponseBodyDataResult) GetFallbackObject() *string {
	return s.FallbackObject
}

func (s *ListIsolationRulesResponseBodyDataResult) GetLimitApp() *string {
	return s.LimitApp
}

func (s *ListIsolationRulesResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIsolationRulesResponseBodyDataResult) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIsolationRulesResponseBodyDataResult) GetResource() *string {
	return s.Resource
}

func (s *ListIsolationRulesResponseBodyDataResult) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListIsolationRulesResponseBodyDataResult) GetThreshold() *int32 {
	return s.Threshold
}

func (s *ListIsolationRulesResponseBodyDataResult) SetAppId(v string) *ListIsolationRulesResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetAppName(v string) *ListIsolationRulesResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetEnable(v bool) *ListIsolationRulesResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetFallbackObject(v string) *ListIsolationRulesResponseBodyDataResult {
	s.FallbackObject = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetLimitApp(v string) *ListIsolationRulesResponseBodyDataResult {
	s.LimitApp = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetNamespace(v string) *ListIsolationRulesResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetRegionId(v string) *ListIsolationRulesResponseBodyDataResult {
	s.RegionId = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetResource(v string) *ListIsolationRulesResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetRuleId(v int64) *ListIsolationRulesResponseBodyDataResult {
	s.RuleId = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) SetThreshold(v int32) *ListIsolationRulesResponseBodyDataResult {
	s.Threshold = &v
	return s
}

func (s *ListIsolationRulesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
