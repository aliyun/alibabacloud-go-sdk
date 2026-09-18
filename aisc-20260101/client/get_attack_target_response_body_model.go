// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAttackTargetResponseBodyData) *GetAttackTargetResponseBody
	GetData() *GetAttackTargetResponseBodyData
	SetRequestId(v string) *GetAttackTargetResponseBody
	GetRequestId() *string
}

type GetAttackTargetResponseBody struct {
	// The scan target details.
	Data *GetAttackTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can use this ID for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttackTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttackTargetResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttackTargetResponseBody) GetData() *GetAttackTargetResponseBodyData {
	return s.Data
}

func (s *GetAttackTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttackTargetResponseBody) SetData(v *GetAttackTargetResponseBodyData) *GetAttackTargetResponseBody {
	s.Data = v
	return s
}

func (s *GetAttackTargetResponseBody) SetRequestId(v string) *GetAttackTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttackTargetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAttackTargetResponseBodyData struct {
	// The advanced connection configuration (JSON character string). For common fields and provider configuration templates, refer to the ConnectionConfig parameter of CreateAttackTarget.
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
	// The time when the target was created. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1735689600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the scan target. This value may be empty.
	//
	// example:
	//
	// Production Bailian application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The HTTP or HTTPS endpoint of the target model service. When ConnectionMethod is set to enterprise_relay, this value is a fixed internal endpoint.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The time of the first scan. This value is a UNIX timestamp in milliseconds. This is an aggregate field that is not populated by this operation and returns an empty value. Refer to ListAttackTargets.
	//
	// example:
	//
	// 1735689600000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The failure reason of the most recent scan task. This is an aggregate field that is not populated by this operation and returns an empty value. Refer to ListAttackTargets.
	//
	// example:
	//
	// Execution timed out
	LastScanFailMessage *string `json:"LastScanFailMessage,omitempty" xml:"LastScanFailMessage,omitempty"`
	// The status of the most recent scan task. This is an aggregate field that is not populated by this operation and returns an empty value. Refer to ListAttackTargets.
	//
	// example:
	//
	// completed
	LastScanStatus *string `json:"LastScanStatus,omitempty" xml:"LastScanStatus,omitempty"`
	// The time of the most recent scan. This value is a UNIX timestamp in milliseconds. This is an aggregate field that is not populated by this operation and returns an empty value. Refer to ListAttackTargets.
	//
	// example:
	//
	// 1735689600000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The name of the target model. When ConnectionMethod is set to enterprise_relay, this value is the fixed platform value Agent.
	//
	// example:
	//
	// qwen-flash
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The time when the target was last modified. This value is a UNIX timestamp in milliseconds.
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
	// The risk level of the task result. This is an aggregate field that is not populated by this operation and returns an empty value. Refer to ListAttackTargets.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The cumulative number of scans. This is an aggregate field that is not populated by this operation and returns an empty value. For the meaning and example values, refer to the ListAttackTargets response.
	//
	// example:
	//
	// 12
	ScanCount *int64 `json:"ScanCount,omitempty" xml:"ScanCount,omitempty"`
	// The scan node configuration (JSON character string). Common fields include scanType (scan pattern: attack for security attack scan, tc260 for TC260 filing scan), scannerType (execute DPI engine: classic for per-sample execute, agent for multi-round autonomous attack), and sampleScope (sample scope: version for the current effective version, all for full samples). If the target is not configured, the default configurations are returned.
	//
	// example:
	//
	// {"scanType":"attack","scannerType":"classic","sampleScope":"all"}
	ScanTaskConfig *string `json:"ScanTaskConfig,omitempty" xml:"ScanTaskConfig,omitempty"`
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
}

func (s GetAttackTargetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAttackTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttackTargetResponseBodyData) GetConnectionConfig() *string {
	return s.ConnectionConfig
}

func (s *GetAttackTargetResponseBodyData) GetConnectionMethod() *string {
	return s.ConnectionMethod
}

func (s *GetAttackTargetResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAttackTargetResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAttackTargetResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetAttackTargetResponseBodyData) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *GetAttackTargetResponseBodyData) GetLastScanFailMessage() *string {
	return s.LastScanFailMessage
}

func (s *GetAttackTargetResponseBodyData) GetLastScanStatus() *string {
	return s.LastScanStatus
}

func (s *GetAttackTargetResponseBodyData) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *GetAttackTargetResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *GetAttackTargetResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetAttackTargetResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *GetAttackTargetResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetAttackTargetResponseBodyData) GetScanCount() *int64 {
	return s.ScanCount
}

func (s *GetAttackTargetResponseBodyData) GetScanTaskConfig() *string {
	return s.ScanTaskConfig
}

func (s *GetAttackTargetResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *GetAttackTargetResponseBodyData) GetTargetName() *string {
	return s.TargetName
}

func (s *GetAttackTargetResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *GetAttackTargetResponseBodyData) SetConnectionConfig(v string) *GetAttackTargetResponseBodyData {
	s.ConnectionConfig = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetConnectionMethod(v string) *GetAttackTargetResponseBodyData {
	s.ConnectionMethod = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetCreateTime(v int64) *GetAttackTargetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetDescription(v string) *GetAttackTargetResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetEndpoint(v string) *GetAttackTargetResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetFirstScanTime(v int64) *GetAttackTargetResponseBodyData {
	s.FirstScanTime = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetLastScanFailMessage(v string) *GetAttackTargetResponseBodyData {
	s.LastScanFailMessage = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetLastScanStatus(v string) *GetAttackTargetResponseBodyData {
	s.LastScanStatus = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetLastScanTime(v int64) *GetAttackTargetResponseBodyData {
	s.LastScanTime = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetModelName(v string) *GetAttackTargetResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetModifyTime(v int64) *GetAttackTargetResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetProvider(v string) *GetAttackTargetResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetRiskLevel(v string) *GetAttackTargetResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetScanCount(v int64) *GetAttackTargetResponseBodyData {
	s.ScanCount = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetScanTaskConfig(v string) *GetAttackTargetResponseBodyData {
	s.ScanTaskConfig = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetTargetId(v string) *GetAttackTargetResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetTargetName(v string) *GetAttackTargetResponseBodyData {
	s.TargetName = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) SetTargetType(v string) *GetAttackTargetResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *GetAttackTargetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
