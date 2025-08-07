// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *GetIssuesResponseBody
	GetArgs() map[string]interface{}
	SetErrorCode(v int32) *GetIssuesResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *GetIssuesResponseBody
	GetMessage() *string
	SetModel(v *GetIssuesResponseBodyModel) *GetIssuesResponseBody
	GetModel() *GetIssuesResponseBodyModel
	SetRequestId(v string) *GetIssuesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIssuesResponseBody
	GetSuccess() *bool
}

type GetIssuesResponseBody struct {
	// Args
	Args map[string]interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// successful
	Message *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetIssuesResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 22111548-55D2-4258-9B18-273E4C134444
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIssuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesResponseBody) GoString() string {
	return s.String()
}

func (s *GetIssuesResponseBody) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *GetIssuesResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetIssuesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIssuesResponseBody) GetModel() *GetIssuesResponseBodyModel {
	return s.Model
}

func (s *GetIssuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIssuesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIssuesResponseBody) SetArgs(v map[string]interface{}) *GetIssuesResponseBody {
	s.Args = v
	return s
}

func (s *GetIssuesResponseBody) SetErrorCode(v int32) *GetIssuesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetIssuesResponseBody) SetMessage(v string) *GetIssuesResponseBody {
	s.Message = &v
	return s
}

func (s *GetIssuesResponseBody) SetModel(v *GetIssuesResponseBodyModel) *GetIssuesResponseBody {
	s.Model = v
	return s
}

func (s *GetIssuesResponseBody) SetRequestId(v string) *GetIssuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIssuesResponseBody) SetSuccess(v bool) *GetIssuesResponseBody {
	s.Success = &v
	return s
}

func (s *GetIssuesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIssuesResponseBodyModel struct {
	Items []*GetIssuesResponseBodyModelItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Pages *int32 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetIssuesResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetIssuesResponseBodyModel) GetItems() []*GetIssuesResponseBodyModelItems {
	return s.Items
}

func (s *GetIssuesResponseBodyModel) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetIssuesResponseBodyModel) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetIssuesResponseBodyModel) GetPages() *int32 {
	return s.Pages
}

func (s *GetIssuesResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *GetIssuesResponseBodyModel) SetItems(v []*GetIssuesResponseBodyModelItems) *GetIssuesResponseBodyModel {
	s.Items = v
	return s
}

func (s *GetIssuesResponseBodyModel) SetPageNum(v int32) *GetIssuesResponseBodyModel {
	s.PageNum = &v
	return s
}

func (s *GetIssuesResponseBodyModel) SetPageSize(v int32) *GetIssuesResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *GetIssuesResponseBodyModel) SetPages(v int32) *GetIssuesResponseBodyModel {
	s.Pages = &v
	return s
}

func (s *GetIssuesResponseBodyModel) SetTotal(v int64) *GetIssuesResponseBodyModel {
	s.Total = &v
	return s
}

func (s *GetIssuesResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

type GetIssuesResponseBodyModelItems struct {
	// example:
	//
	// 1
	AffectedUserCount *int32 `json:"AffectedUserCount,omitempty" xml:"AffectedUserCount,omitempty"`
	AllocSizeMax      *int64 `json:"AllocSizeMax,omitempty" xml:"AllocSizeMax,omitempty"`
	AllocSizePct50    *int64 `json:"AllocSizePct50,omitempty" xml:"AllocSizePct50,omitempty"`
	AllocSizePct70    *int64 `json:"AllocSizePct70,omitempty" xml:"AllocSizePct70,omitempty"`
	AllocSizePct90    *int64 `json:"AllocSizePct90,omitempty" xml:"AllocSizePct90,omitempty"`
	// example:
	//
	// -3481243636390427020
	DigestHash *string `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	// example:
	//
	// 1
	DomScore *string `json:"DomScore,omitempty" xml:"DomScore,omitempty"`
	// example:
	//
	// 1
	ErrorColumn *int32 `json:"ErrorColumn,omitempty" xml:"ErrorColumn,omitempty"`
	// example:
	//
	// 2
	ErrorCount *int32 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// example:
	//
	// 1
	ErrorDeviceCount *int32 `json:"ErrorDeviceCount,omitempty" xml:"ErrorDeviceCount,omitempty"`
	// example:
	//
	// 1.0
	ErrorDeviceRate *float64 `json:"ErrorDeviceRate,omitempty" xml:"ErrorDeviceRate,omitempty"`
	// example:
	//
	// test.js
	ErrorFileName *string `json:"ErrorFileName,omitempty" xml:"ErrorFileName,omitempty"`
	// example:
	//
	// 1
	ErrorLine *int32 `json:"ErrorLine,omitempty" xml:"ErrorLine,omitempty"`
	// example:
	//
	// ErrorName
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	// example:
	//
	// 0
	ErrorRate *float64 `json:"ErrorRate,omitempty" xml:"ErrorRate,omitempty"`
	// example:
	//
	// Error
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// example:
	//
	// 1691745496851
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// 1.0.0
	FirstVersion *string `json:"FirstVersion,omitempty" xml:"FirstVersion,omitempty"`
	// example:
	//
	// ServiceType
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// java.lang.NullPointerException: Attempt to invoke virtual method \\"java.lang.Object java.lang.ref.WeakReference.get()\\" on a null object reference
	//
	// 	at e.l.a.a.d.h(BasePresenter.java:1)
	//
	// 	at e.l.a.g.c.e.g.s1(GoodsPigLoadPresenter.java:1)
	//
	// 	at e.l.a.h.d.a$a.a(AliOssManager.java:2)
	Stack *string `json:"Stack,omitempty" xml:"Stack,omitempty"`
	// example:
	//
	// CREATE_COMPLETE
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIssuesResponseBodyModelItems) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesResponseBodyModelItems) GoString() string {
	return s.String()
}

func (s *GetIssuesResponseBodyModelItems) GetAffectedUserCount() *int32 {
	return s.AffectedUserCount
}

func (s *GetIssuesResponseBodyModelItems) GetAllocSizeMax() *int64 {
	return s.AllocSizeMax
}

func (s *GetIssuesResponseBodyModelItems) GetAllocSizePct50() *int64 {
	return s.AllocSizePct50
}

func (s *GetIssuesResponseBodyModelItems) GetAllocSizePct70() *int64 {
	return s.AllocSizePct70
}

func (s *GetIssuesResponseBodyModelItems) GetAllocSizePct90() *int64 {
	return s.AllocSizePct90
}

func (s *GetIssuesResponseBodyModelItems) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetIssuesResponseBodyModelItems) GetDomScore() *string {
	return s.DomScore
}

func (s *GetIssuesResponseBodyModelItems) GetErrorColumn() *int32 {
	return s.ErrorColumn
}

func (s *GetIssuesResponseBodyModelItems) GetErrorCount() *int32 {
	return s.ErrorCount
}

func (s *GetIssuesResponseBodyModelItems) GetErrorDeviceCount() *int32 {
	return s.ErrorDeviceCount
}

func (s *GetIssuesResponseBodyModelItems) GetErrorDeviceRate() *float64 {
	return s.ErrorDeviceRate
}

func (s *GetIssuesResponseBodyModelItems) GetErrorFileName() *string {
	return s.ErrorFileName
}

func (s *GetIssuesResponseBodyModelItems) GetErrorLine() *int32 {
	return s.ErrorLine
}

func (s *GetIssuesResponseBodyModelItems) GetErrorName() *string {
	return s.ErrorName
}

func (s *GetIssuesResponseBodyModelItems) GetErrorRate() *float64 {
	return s.ErrorRate
}

func (s *GetIssuesResponseBodyModelItems) GetErrorType() *string {
	return s.ErrorType
}

func (s *GetIssuesResponseBodyModelItems) GetEventTime() *string {
	return s.EventTime
}

func (s *GetIssuesResponseBodyModelItems) GetFirstVersion() *string {
	return s.FirstVersion
}

func (s *GetIssuesResponseBodyModelItems) GetName() *string {
	return s.Name
}

func (s *GetIssuesResponseBodyModelItems) GetStack() *string {
	return s.Stack
}

func (s *GetIssuesResponseBodyModelItems) GetStatus() *int32 {
	return s.Status
}

func (s *GetIssuesResponseBodyModelItems) SetAffectedUserCount(v int32) *GetIssuesResponseBodyModelItems {
	s.AffectedUserCount = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetAllocSizeMax(v int64) *GetIssuesResponseBodyModelItems {
	s.AllocSizeMax = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetAllocSizePct50(v int64) *GetIssuesResponseBodyModelItems {
	s.AllocSizePct50 = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetAllocSizePct70(v int64) *GetIssuesResponseBodyModelItems {
	s.AllocSizePct70 = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetAllocSizePct90(v int64) *GetIssuesResponseBodyModelItems {
	s.AllocSizePct90 = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetDigestHash(v string) *GetIssuesResponseBodyModelItems {
	s.DigestHash = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetDomScore(v string) *GetIssuesResponseBodyModelItems {
	s.DomScore = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorColumn(v int32) *GetIssuesResponseBodyModelItems {
	s.ErrorColumn = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorCount(v int32) *GetIssuesResponseBodyModelItems {
	s.ErrorCount = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorDeviceCount(v int32) *GetIssuesResponseBodyModelItems {
	s.ErrorDeviceCount = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorDeviceRate(v float64) *GetIssuesResponseBodyModelItems {
	s.ErrorDeviceRate = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorFileName(v string) *GetIssuesResponseBodyModelItems {
	s.ErrorFileName = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorLine(v int32) *GetIssuesResponseBodyModelItems {
	s.ErrorLine = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorName(v string) *GetIssuesResponseBodyModelItems {
	s.ErrorName = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorRate(v float64) *GetIssuesResponseBodyModelItems {
	s.ErrorRate = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetErrorType(v string) *GetIssuesResponseBodyModelItems {
	s.ErrorType = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetEventTime(v string) *GetIssuesResponseBodyModelItems {
	s.EventTime = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetFirstVersion(v string) *GetIssuesResponseBodyModelItems {
	s.FirstVersion = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetName(v string) *GetIssuesResponseBodyModelItems {
	s.Name = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetStack(v string) *GetIssuesResponseBodyModelItems {
	s.Stack = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) SetStatus(v int32) *GetIssuesResponseBodyModelItems {
	s.Status = &v
	return s
}

func (s *GetIssuesResponseBodyModelItems) Validate() error {
	return dara.Validate(s)
}
