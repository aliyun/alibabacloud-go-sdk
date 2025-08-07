// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *GetIssueResponseBody
	GetArgs() map[string]interface{}
	SetErrorCode(v int32) *GetIssueResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *GetIssueResponseBody
	GetMessage() *string
	SetModel(v *GetIssueResponseBodyModel) *GetIssueResponseBody
	GetModel() *GetIssueResponseBodyModel
	SetRequestId(v string) *GetIssueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIssueResponseBody
	GetSuccess() *bool
}

type GetIssueResponseBody struct {
	// Args
	Args map[string]interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetIssueResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// A8313212-EB4E-4E15-A7F9-D9C8F3FE8E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIssueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIssueResponseBody) GoString() string {
	return s.String()
}

func (s *GetIssueResponseBody) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *GetIssueResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetIssueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIssueResponseBody) GetModel() *GetIssueResponseBodyModel {
	return s.Model
}

func (s *GetIssueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIssueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIssueResponseBody) SetArgs(v map[string]interface{}) *GetIssueResponseBody {
	s.Args = v
	return s
}

func (s *GetIssueResponseBody) SetErrorCode(v int32) *GetIssueResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetIssueResponseBody) SetMessage(v string) *GetIssueResponseBody {
	s.Message = &v
	return s
}

func (s *GetIssueResponseBody) SetModel(v *GetIssueResponseBodyModel) *GetIssueResponseBody {
	s.Model = v
	return s
}

func (s *GetIssueResponseBody) SetRequestId(v string) *GetIssueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIssueResponseBody) SetSuccess(v bool) *GetIssueResponseBody {
	s.Success = &v
	return s
}

func (s *GetIssueResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIssueResponseBodyModel struct {
	AffectedVersions []*string `json:"AffectedVersions,omitempty" xml:"AffectedVersions,omitempty" type:"Repeated"`
	AllocSizeMax     *int64    `json:"AllocSizeMax,omitempty" xml:"AllocSizeMax,omitempty"`
	AllocSizePct50   *int64    `json:"AllocSizePct50,omitempty" xml:"AllocSizePct50,omitempty"`
	AllocSizePct70   *int64    `json:"AllocSizePct70,omitempty" xml:"AllocSizePct70,omitempty"`
	AllocSizePct90   *int64    `json:"AllocSizePct90,omitempty" xml:"AllocSizePct90,omitempty"`
	// example:
	//
	// java.lang.NullPointerException: Attempt to invoke virtual method \\"java.lang.Object java.lang.ref.WeakReference.get()\\" on a null object reference
	//
	// 	at e.l.a.a.d.h(BasePresenter.java:1)
	//
	// 	at e.l.a.g.c.e.g.s1(GoodsPigLoadPresenter.java:1)
	//
	// 	at e.l.a.h.d.a$a.a(AliOssManager.java:2)
	CruxStack *string `json:"CruxStack,omitempty" xml:"CruxStack,omitempty"`
	// example:
	//
	// -6428474329608402395
	DigestHash *string `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	// example:
	//
	// 2
	ErrorColumn *int32 `json:"ErrorColumn,omitempty" xml:"ErrorColumn,omitempty"`
	// example:
	//
	// 2
	ErrorCount *int32 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// example:
	//
	// 0.2
	ErrorCountGrowthRate *float64 `json:"ErrorCountGrowthRate,omitempty" xml:"ErrorCountGrowthRate,omitempty"`
	// example:
	//
	// 4
	ErrorDeviceCount *int32 `json:"ErrorDeviceCount,omitempty" xml:"ErrorDeviceCount,omitempty"`
	// example:
	//
	// 0.2
	ErrorDeviceCountGrowthRate *float64 `json:"ErrorDeviceCountGrowthRate,omitempty" xml:"ErrorDeviceCountGrowthRate,omitempty"`
	// example:
	//
	// 0.2
	ErrorDeviceRate *float64 `json:"ErrorDeviceRate,omitempty" xml:"ErrorDeviceRate,omitempty"`
	// example:
	//
	// 0.2
	ErrorDeviceRateGrowthRate *float64 `json:"ErrorDeviceRateGrowthRate,omitempty" xml:"ErrorDeviceRateGrowthRate,omitempty"`
	// example:
	//
	// test.js
	ErrorFileName *string `json:"ErrorFileName,omitempty" xml:"ErrorFileName,omitempty"`
	// example:
	//
	// 1
	ErrorLine *string `json:"ErrorLine,omitempty" xml:"ErrorLine,omitempty"`
	// example:
	//
	// Error
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	// example:
	//
	// 0
	ErrorRate *float64 `json:"ErrorRate,omitempty" xml:"ErrorRate,omitempty"`
	// example:
	//
	// 0.2
	ErrorRateGrowthRate *float64 `json:"ErrorRateGrowthRate,omitempty" xml:"ErrorRateGrowthRate,omitempty"`
	// example:
	//
	// ErrorType
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	EventTime *int64  `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// 1.0.0
	FirstVersion *string `json:"FirstVersion,omitempty" xml:"FirstVersion,omitempty"`
	// example:
	//
	// 1673423227000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1740489688615
	GmtLatest *int64 `json:"GmtLatest,omitempty" xml:"GmtLatest,omitempty"`
	// example:
	//
	// 1
	KeyLine *int32  `json:"KeyLine,omitempty" xml:"KeyLine,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// java.lang.NullPointerException
	Stack *string `json:"Stack,omitempty" xml:"Stack,omitempty"`
	// example:
	//
	// SUCCESS
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// trustee instance
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// true
	SymbolicStatus *bool     `json:"SymbolicStatus,omitempty" xml:"SymbolicStatus,omitempty"`
	Tags           []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetIssueResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetIssueResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetIssueResponseBodyModel) GetAffectedVersions() []*string {
	return s.AffectedVersions
}

func (s *GetIssueResponseBodyModel) GetAllocSizeMax() *int64 {
	return s.AllocSizeMax
}

func (s *GetIssueResponseBodyModel) GetAllocSizePct50() *int64 {
	return s.AllocSizePct50
}

func (s *GetIssueResponseBodyModel) GetAllocSizePct70() *int64 {
	return s.AllocSizePct70
}

func (s *GetIssueResponseBodyModel) GetAllocSizePct90() *int64 {
	return s.AllocSizePct90
}

func (s *GetIssueResponseBodyModel) GetCruxStack() *string {
	return s.CruxStack
}

func (s *GetIssueResponseBodyModel) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetIssueResponseBodyModel) GetErrorColumn() *int32 {
	return s.ErrorColumn
}

func (s *GetIssueResponseBodyModel) GetErrorCount() *int32 {
	return s.ErrorCount
}

func (s *GetIssueResponseBodyModel) GetErrorCountGrowthRate() *float64 {
	return s.ErrorCountGrowthRate
}

func (s *GetIssueResponseBodyModel) GetErrorDeviceCount() *int32 {
	return s.ErrorDeviceCount
}

func (s *GetIssueResponseBodyModel) GetErrorDeviceCountGrowthRate() *float64 {
	return s.ErrorDeviceCountGrowthRate
}

func (s *GetIssueResponseBodyModel) GetErrorDeviceRate() *float64 {
	return s.ErrorDeviceRate
}

func (s *GetIssueResponseBodyModel) GetErrorDeviceRateGrowthRate() *float64 {
	return s.ErrorDeviceRateGrowthRate
}

func (s *GetIssueResponseBodyModel) GetErrorFileName() *string {
	return s.ErrorFileName
}

func (s *GetIssueResponseBodyModel) GetErrorLine() *string {
	return s.ErrorLine
}

func (s *GetIssueResponseBodyModel) GetErrorName() *string {
	return s.ErrorName
}

func (s *GetIssueResponseBodyModel) GetErrorRate() *float64 {
	return s.ErrorRate
}

func (s *GetIssueResponseBodyModel) GetErrorRateGrowthRate() *float64 {
	return s.ErrorRateGrowthRate
}

func (s *GetIssueResponseBodyModel) GetErrorType() *string {
	return s.ErrorType
}

func (s *GetIssueResponseBodyModel) GetEventTime() *int64 {
	return s.EventTime
}

func (s *GetIssueResponseBodyModel) GetFirstVersion() *string {
	return s.FirstVersion
}

func (s *GetIssueResponseBodyModel) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetIssueResponseBodyModel) GetGmtLatest() *int64 {
	return s.GmtLatest
}

func (s *GetIssueResponseBodyModel) GetKeyLine() *int32 {
	return s.KeyLine
}

func (s *GetIssueResponseBodyModel) GetName() *string {
	return s.Name
}

func (s *GetIssueResponseBodyModel) GetStack() *string {
	return s.Stack
}

func (s *GetIssueResponseBodyModel) GetStatus() *int32 {
	return s.Status
}

func (s *GetIssueResponseBodyModel) GetSummary() *string {
	return s.Summary
}

func (s *GetIssueResponseBodyModel) GetSymbolicStatus() *bool {
	return s.SymbolicStatus
}

func (s *GetIssueResponseBodyModel) GetTags() []*string {
	return s.Tags
}

func (s *GetIssueResponseBodyModel) SetAffectedVersions(v []*string) *GetIssueResponseBodyModel {
	s.AffectedVersions = v
	return s
}

func (s *GetIssueResponseBodyModel) SetAllocSizeMax(v int64) *GetIssueResponseBodyModel {
	s.AllocSizeMax = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetAllocSizePct50(v int64) *GetIssueResponseBodyModel {
	s.AllocSizePct50 = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetAllocSizePct70(v int64) *GetIssueResponseBodyModel {
	s.AllocSizePct70 = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetAllocSizePct90(v int64) *GetIssueResponseBodyModel {
	s.AllocSizePct90 = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetCruxStack(v string) *GetIssueResponseBodyModel {
	s.CruxStack = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetDigestHash(v string) *GetIssueResponseBodyModel {
	s.DigestHash = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorColumn(v int32) *GetIssueResponseBodyModel {
	s.ErrorColumn = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorCount(v int32) *GetIssueResponseBodyModel {
	s.ErrorCount = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorCountGrowthRate(v float64) *GetIssueResponseBodyModel {
	s.ErrorCountGrowthRate = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorDeviceCount(v int32) *GetIssueResponseBodyModel {
	s.ErrorDeviceCount = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorDeviceCountGrowthRate(v float64) *GetIssueResponseBodyModel {
	s.ErrorDeviceCountGrowthRate = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorDeviceRate(v float64) *GetIssueResponseBodyModel {
	s.ErrorDeviceRate = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorDeviceRateGrowthRate(v float64) *GetIssueResponseBodyModel {
	s.ErrorDeviceRateGrowthRate = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorFileName(v string) *GetIssueResponseBodyModel {
	s.ErrorFileName = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorLine(v string) *GetIssueResponseBodyModel {
	s.ErrorLine = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorName(v string) *GetIssueResponseBodyModel {
	s.ErrorName = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorRate(v float64) *GetIssueResponseBodyModel {
	s.ErrorRate = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorRateGrowthRate(v float64) *GetIssueResponseBodyModel {
	s.ErrorRateGrowthRate = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetErrorType(v string) *GetIssueResponseBodyModel {
	s.ErrorType = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetEventTime(v int64) *GetIssueResponseBodyModel {
	s.EventTime = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetFirstVersion(v string) *GetIssueResponseBodyModel {
	s.FirstVersion = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetGmtCreate(v int64) *GetIssueResponseBodyModel {
	s.GmtCreate = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetGmtLatest(v int64) *GetIssueResponseBodyModel {
	s.GmtLatest = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetKeyLine(v int32) *GetIssueResponseBodyModel {
	s.KeyLine = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetName(v string) *GetIssueResponseBodyModel {
	s.Name = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetStack(v string) *GetIssueResponseBodyModel {
	s.Stack = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetStatus(v int32) *GetIssueResponseBodyModel {
	s.Status = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetSummary(v string) *GetIssueResponseBodyModel {
	s.Summary = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetSymbolicStatus(v bool) *GetIssueResponseBodyModel {
	s.SymbolicStatus = &v
	return s
}

func (s *GetIssueResponseBodyModel) SetTags(v []*string) *GetIssueResponseBodyModel {
	s.Tags = v
	return s
}

func (s *GetIssueResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
