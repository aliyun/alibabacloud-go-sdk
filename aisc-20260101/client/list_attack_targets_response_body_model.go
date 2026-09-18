// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAttackTargetsResponseBodyData) *ListAttackTargetsResponseBody
	GetData() []*ListAttackTargetsResponseBodyData
	SetPageNumber(v int64) *ListAttackTargetsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAttackTargetsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAttackTargetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAttackTargetsResponseBody
	GetTotalCount() *int64
}

type ListAttackTargetsResponseBody struct {
	// The list of scan targets on the current page.
	Data []*ListAttackTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The normalized page number that takes effect. The value may differ from the input parameter.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The normalized number of entries per page that takes effect. The value may differ from the input parameter.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID. You can use this ID for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of scan targets that match the filter conditions. In post-filtering scenarios, this value is the total count after in-memory filtering.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAttackTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAttackTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAttackTargetsResponseBody) GetData() []*ListAttackTargetsResponseBodyData {
	return s.Data
}

func (s *ListAttackTargetsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAttackTargetsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAttackTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAttackTargetsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAttackTargetsResponseBody) SetData(v []*ListAttackTargetsResponseBodyData) *ListAttackTargetsResponseBody {
	s.Data = v
	return s
}

func (s *ListAttackTargetsResponseBody) SetPageNumber(v int64) *ListAttackTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAttackTargetsResponseBody) SetPageSize(v int64) *ListAttackTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAttackTargetsResponseBody) SetRequestId(v string) *ListAttackTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAttackTargetsResponseBody) SetTotalCount(v int64) *ListAttackTargetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAttackTargetsResponseBody) Validate() error {
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

type ListAttackTargetsResponseBodyData struct {
	// The advanced connection configuration for the target (JSON character string). For common fields and provider configuration templates, see the ConnectionConfig parameter description of the CreateAttackTarget operation.
	//
	// example:
	//
	// {\\"httpMethod\\":\\"POST\\",\\"authType\\":\\"bearer\\",\\"timeoutMs\\":30000,\\"requestTemplate\\":\\"{\\\\\\"input\\\\\\":{\\\\\\"prompt\\\\\\":\\\\\\"{{prompt}}\\\\\\"},\\\\\\"parameters\\\\\\":{\\\\\\"incremental_output\\\\\\":true},\\\\\\"debug\\\\\\":{}}\\",\\"messageJsonPath\\":\\"$.output.text\\",\\"requestHeaders\\":\\"{\\\\\\"X-DashScope-SSE\\\\\\": \\\\\\"enable\\\\\\" }\\",\\"stream\\":true,\\"customAuthHeaderName\\":\\"\\"}
	ConnectionConfig *string `json:"ConnectionConfig,omitempty" xml:"ConnectionConfig,omitempty"`
	// The connection protocol type of the target service.
	//
	// example:
	//
	// openai
	ConnectionMethod *string `json:"ConnectionMethod,omitempty" xml:"ConnectionMethod,omitempty"`
	// The time when the target was created. The value is a millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1735689600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the scan target.
	//
	// example:
	//
	// Production Bailian application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The HTTP or HTTPS endpoint address of the target model service.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The time when the first scan was performed. The value is a millisecond-level UNIX timestamp. This parameter is null if no scan has been performed.
	//
	// example:
	//
	// 1735689600000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The failure reason of the most recent scan task. This parameter is empty if the most recent scan did not fail.
	//
	// example:
	//
	// Scan executor connection timed out
	LastScanFailMessage *string `json:"LastScanFailMessage,omitempty" xml:"LastScanFailMessage,omitempty"`
	// The detection status of the most recent scan task.
	//
	// example:
	//
	// completed
	LastScanStatus *string `json:"LastScanStatus,omitempty" xml:"LastScanStatus,omitempty"`
	// The time when the last scan was performed. The value is a millisecond-level UNIX timestamp. This parameter is null if no scan has been performed.
	//
	// example:
	//
	// 1735689600000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The detailed message of the most recent connectivity verification. If the verification succeeded, the value is a response content snippet returned by the target service. If the verification failed, the value is the error reason.
	//
	// example:
	//
	// Connectivity verification succeeded
	LastVerifyMessage *string `json:"LastVerifyMessage,omitempty" xml:"LastVerifyMessage,omitempty"`
	// The name of the target model.
	//
	// example:
	//
	// qwen-flash
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The time when the target was last modified. The value is a millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1735689600000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The business label of the model or agent provider.
	//
	// example:
	//
	// bailian
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The risk level derived from the most recent completed scan task. This parameter is null if no scan has been performed.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The total number of scans performed. The value is 0 if no scan has been performed.
	//
	// example:
	//
	// 12
	ScanCount *int64 `json:"ScanCount,omitempty" xml:"ScanCount,omitempty"`
	// The unique identifier of the scan target.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The display name of the scan target.
	//
	// example:
	//
	// My Bailian target
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the scan target.
	//
	// example:
	//
	// model
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The connectivity verification status of the target.
	//
	// example:
	//
	// verified
	VerifyStatus *string `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s ListAttackTargetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAttackTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAttackTargetsResponseBodyData) GetConnectionConfig() *string {
	return s.ConnectionConfig
}

func (s *ListAttackTargetsResponseBodyData) GetConnectionMethod() *string {
	return s.ConnectionMethod
}

func (s *ListAttackTargetsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAttackTargetsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListAttackTargetsResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListAttackTargetsResponseBodyData) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *ListAttackTargetsResponseBodyData) GetLastScanFailMessage() *string {
	return s.LastScanFailMessage
}

func (s *ListAttackTargetsResponseBodyData) GetLastScanStatus() *string {
	return s.LastScanStatus
}

func (s *ListAttackTargetsResponseBodyData) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *ListAttackTargetsResponseBodyData) GetLastVerifyMessage() *string {
	return s.LastVerifyMessage
}

func (s *ListAttackTargetsResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *ListAttackTargetsResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListAttackTargetsResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *ListAttackTargetsResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListAttackTargetsResponseBodyData) GetScanCount() *int64 {
	return s.ScanCount
}

func (s *ListAttackTargetsResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAttackTargetsResponseBodyData) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAttackTargetsResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAttackTargetsResponseBodyData) GetVerifyStatus() *string {
	return s.VerifyStatus
}

func (s *ListAttackTargetsResponseBodyData) SetConnectionConfig(v string) *ListAttackTargetsResponseBodyData {
	s.ConnectionConfig = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetConnectionMethod(v string) *ListAttackTargetsResponseBodyData {
	s.ConnectionMethod = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetCreateTime(v int64) *ListAttackTargetsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetDescription(v string) *ListAttackTargetsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetEndpoint(v string) *ListAttackTargetsResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetFirstScanTime(v int64) *ListAttackTargetsResponseBodyData {
	s.FirstScanTime = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetLastScanFailMessage(v string) *ListAttackTargetsResponseBodyData {
	s.LastScanFailMessage = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetLastScanStatus(v string) *ListAttackTargetsResponseBodyData {
	s.LastScanStatus = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetLastScanTime(v int64) *ListAttackTargetsResponseBodyData {
	s.LastScanTime = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetLastVerifyMessage(v string) *ListAttackTargetsResponseBodyData {
	s.LastVerifyMessage = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetModelName(v string) *ListAttackTargetsResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetModifyTime(v int64) *ListAttackTargetsResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetProvider(v string) *ListAttackTargetsResponseBodyData {
	s.Provider = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetRiskLevel(v string) *ListAttackTargetsResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetScanCount(v int64) *ListAttackTargetsResponseBodyData {
	s.ScanCount = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetTargetId(v string) *ListAttackTargetsResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetTargetName(v string) *ListAttackTargetsResponseBodyData {
	s.TargetName = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetTargetType(v string) *ListAttackTargetsResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) SetVerifyStatus(v string) *ListAttackTargetsResponseBodyData {
	s.VerifyStatus = &v
	return s
}

func (s *ListAttackTargetsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
