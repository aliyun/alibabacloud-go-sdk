// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSubTaskResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *GetSubTaskResultResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetSubTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSubTaskResultResponseBody
	GetRequestId() *string
	SetResultObject(v *GetSubTaskResultResponseBodyResultObject) *GetSubTaskResultResponseBody
	GetResultObject() *GetSubTaskResultResponseBodyResultObject
}

type GetSubTaskResultResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0A519CFA-0EEC-580A-A5C1-F9C653FB2354
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *GetSubTaskResultResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s GetSubTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSubTaskResultResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetSubTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubTaskResultResponseBody) GetResultObject() *GetSubTaskResultResponseBodyResultObject {
	return s.ResultObject
}

func (s *GetSubTaskResultResponseBody) SetCode(v string) *GetSubTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubTaskResultResponseBody) SetHttpStatusCode(v string) *GetSubTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSubTaskResultResponseBody) SetMessage(v string) *GetSubTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubTaskResultResponseBody) SetRequestId(v string) *GetSubTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubTaskResultResponseBody) SetResultObject(v *GetSubTaskResultResponseBodyResultObject) *GetSubTaskResultResponseBody {
	s.ResultObject = v
	return s
}

func (s *GetSubTaskResultResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubTaskResultResponseBodyResultObject struct {
	Config    []*GetSubTaskResultResponseBodyResultObjectConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
	ExtraInfo *string                                           `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	File      *GetSubTaskResultResponseBodyResultObjectFile     `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// example:
	//
	// enorl-20w-0926.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// CSV
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// saf/cpoc/34cd7959590ef568086035b956210495/1760580976089_XN_test_1016_100000.csv
	FileUrl  *string                                        `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	IsCharge *bool                                          `json:"IsCharge,omitempty" xml:"IsCharge,omitempty"`
	Log      []*GetSubTaskResultResponseBodyResultObjectLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Repeated"`
	// example:
	//
	// managed.by.apig
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// http://shuanglu-prod.oss-cn-shanghai-finance-1-pub.aliyuncs.com/idrs/24/local/remoteresult?Expires=1756436489&OSSAccessKeyId=****&Signature=****
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// anti_fraud_customed_v3
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// CLASS_CHANGING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 19150
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// icekredit_model_A_2025a_508185
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 5129547232
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSubTaskResultResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultResponseBodyResultObject) GetConfig() []*GetSubTaskResultResponseBodyResultObjectConfig {
	return s.Config
}

func (s *GetSubTaskResultResponseBodyResultObject) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetSubTaskResultResponseBodyResultObject) GetFile() *GetSubTaskResultResponseBodyResultObjectFile {
	return s.File
}

func (s *GetSubTaskResultResponseBodyResultObject) GetFileName() *string {
	return s.FileName
}

func (s *GetSubTaskResultResponseBodyResultObject) GetFileType() *string {
	return s.FileType
}

func (s *GetSubTaskResultResponseBodyResultObject) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetSubTaskResultResponseBodyResultObject) GetIsCharge() *bool {
	return s.IsCharge
}

func (s *GetSubTaskResultResponseBodyResultObject) GetLog() []*GetSubTaskResultResponseBodyResultObjectLog {
	return s.Log
}

func (s *GetSubTaskResultResponseBodyResultObject) GetReason() *string {
	return s.Reason
}

func (s *GetSubTaskResultResponseBodyResultObject) GetResultUrl() *string {
	return s.ResultUrl
}

func (s *GetSubTaskResultResponseBodyResultObject) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetSubTaskResultResponseBodyResultObject) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetSubTaskResultResponseBodyResultObject) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetSubTaskResultResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *GetSubTaskResultResponseBodyResultObject) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSubTaskResultResponseBodyResultObject) GetTaskName() *string {
	return s.TaskName
}

func (s *GetSubTaskResultResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *GetSubTaskResultResponseBodyResultObject) SetConfig(v []*GetSubTaskResultResponseBodyResultObjectConfig) *GetSubTaskResultResponseBodyResultObject {
	s.Config = v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetExtraInfo(v string) *GetSubTaskResultResponseBodyResultObject {
	s.ExtraInfo = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetFile(v *GetSubTaskResultResponseBodyResultObjectFile) *GetSubTaskResultResponseBodyResultObject {
	s.File = v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetFileName(v string) *GetSubTaskResultResponseBodyResultObject {
	s.FileName = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetFileType(v string) *GetSubTaskResultResponseBodyResultObject {
	s.FileType = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetFileUrl(v string) *GetSubTaskResultResponseBodyResultObject {
	s.FileUrl = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetIsCharge(v bool) *GetSubTaskResultResponseBodyResultObject {
	s.IsCharge = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetLog(v []*GetSubTaskResultResponseBodyResultObjectLog) *GetSubTaskResultResponseBodyResultObject {
	s.Log = v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetReason(v string) *GetSubTaskResultResponseBodyResultObject {
	s.Reason = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetResultUrl(v string) *GetSubTaskResultResponseBodyResultObject {
	s.ResultUrl = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetScheduleType(v string) *GetSubTaskResultResponseBodyResultObject {
	s.ScheduleType = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetServiceCode(v string) *GetSubTaskResultResponseBodyResultObject {
	s.ServiceCode = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetServiceName(v string) *GetSubTaskResultResponseBodyResultObject {
	s.ServiceName = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetStatus(v string) *GetSubTaskResultResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetTaskId(v string) *GetSubTaskResultResponseBodyResultObject {
	s.TaskId = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetTaskName(v string) *GetSubTaskResultResponseBodyResultObject {
	s.TaskName = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) SetUserId(v int64) *GetSubTaskResultResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObject) Validate() error {
	if s.Config != nil {
		for _, item := range s.Config {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.Log != nil {
		for _, item := range s.Log {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubTaskResultResponseBodyResultObjectConfig struct {
	// example:
	//
	// repl_lag
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// example:
	//
	// des
	ItemDesc *string `json:"ItemDesc,omitempty" xml:"ItemDesc,omitempty"`
	// example:
	//
	// sample
	SampleItem  *string   `json:"SampleItem,omitempty" xml:"SampleItem,omitempty"`
	SampleItems []*string `json:"SampleItems,omitempty" xml:"SampleItems,omitempty" type:"Repeated"`
}

func (s GetSubTaskResultResponseBodyResultObjectConfig) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultResponseBodyResultObjectConfig) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) GetItem() *string {
	return s.Item
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) GetItemDesc() *string {
	return s.ItemDesc
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) GetSampleItem() *string {
	return s.SampleItem
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) GetSampleItems() []*string {
	return s.SampleItems
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) SetItem(v string) *GetSubTaskResultResponseBodyResultObjectConfig {
	s.Item = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) SetItemDesc(v string) *GetSubTaskResultResponseBodyResultObjectConfig {
	s.ItemDesc = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) SetSampleItem(v string) *GetSubTaskResultResponseBodyResultObjectConfig {
	s.SampleItem = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) SetSampleItems(v []*string) *GetSubTaskResultResponseBodyResultObjectConfig {
	s.SampleItems = v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectConfig) Validate() error {
	return dara.Validate(s)
}

type GetSubTaskResultResponseBodyResultObjectFile struct {
	Col   []*GetSubTaskResultResponseBodyResultObjectFileCol `json:"Col,omitempty" xml:"Col,omitempty" type:"Repeated"`
	Title []*string                                          `json:"Title,omitempty" xml:"Title,omitempty" type:"Repeated"`
}

func (s GetSubTaskResultResponseBodyResultObjectFile) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultResponseBodyResultObjectFile) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultResponseBodyResultObjectFile) GetCol() []*GetSubTaskResultResponseBodyResultObjectFileCol {
	return s.Col
}

func (s *GetSubTaskResultResponseBodyResultObjectFile) GetTitle() []*string {
	return s.Title
}

func (s *GetSubTaskResultResponseBodyResultObjectFile) SetCol(v []*GetSubTaskResultResponseBodyResultObjectFileCol) *GetSubTaskResultResponseBodyResultObjectFile {
	s.Col = v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFile) SetTitle(v []*string) *GetSubTaskResultResponseBodyResultObjectFile {
	s.Title = v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFile) Validate() error {
	if s.Col != nil {
		for _, item := range s.Col {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubTaskResultResponseBodyResultObjectFileCol struct {
	// A0。
	//
	// example:
	//
	// A0
	A0 *string `json:"A0,omitempty" xml:"A0,omitempty"`
	// A1。
	//
	// example:
	//
	// A1
	A1 *string `json:"A1,omitempty" xml:"A1,omitempty"`
	// A10。
	//
	// example:
	//
	// A10
	A10 *string `json:"A10,omitempty" xml:"A10,omitempty"`
	// A11。
	//
	// example:
	//
	// A11
	A11 *string `json:"A11,omitempty" xml:"A11,omitempty"`
	// A2。
	//
	// example:
	//
	// A2
	A2 *string `json:"A2,omitempty" xml:"A2,omitempty"`
	// A3。
	//
	// example:
	//
	// A3
	A3 *string `json:"A3,omitempty" xml:"A3,omitempty"`
	// A4。
	//
	// example:
	//
	// A4
	A4 *string `json:"A4,omitempty" xml:"A4,omitempty"`
	// A5。
	//
	// example:
	//
	// A5
	A5 *string `json:"A5,omitempty" xml:"A5,omitempty"`
	// A6。
	//
	// example:
	//
	// A6
	A6 *string `json:"A6,omitempty" xml:"A6,omitempty"`
	// A7。
	//
	// example:
	//
	// A7
	A7 *string `json:"A7,omitempty" xml:"A7,omitempty"`
	// A8。
	//
	// example:
	//
	// A8
	A8 *string `json:"A8,omitempty" xml:"A8,omitempty"`
	// A9。
	//
	// example:
	//
	// A9
	A9 *string `json:"A9,omitempty" xml:"A9,omitempty"`
}

func (s GetSubTaskResultResponseBodyResultObjectFileCol) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultResponseBodyResultObjectFileCol) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA0() *string {
	return s.A0
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA1() *string {
	return s.A1
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA10() *string {
	return s.A10
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA11() *string {
	return s.A11
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA2() *string {
	return s.A2
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA3() *string {
	return s.A3
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA4() *string {
	return s.A4
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA5() *string {
	return s.A5
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA6() *string {
	return s.A6
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA7() *string {
	return s.A7
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA8() *string {
	return s.A8
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) GetA9() *string {
	return s.A9
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA0(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A0 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA1(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A1 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA10(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A10 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA11(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A11 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA2(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A2 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA3(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A3 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA4(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A4 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA5(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A5 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA6(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A6 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA7(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A7 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA8(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A8 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) SetA9(v string) *GetSubTaskResultResponseBodyResultObjectFileCol {
	s.A9 = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectFileCol) Validate() error {
	return dara.Validate(s)
}

type GetSubTaskResultResponseBodyResultObjectLog struct {
	// example:
	//
	// RELEASE
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// PasswordExpired
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 1760408725312
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetSubTaskResultResponseBodyResultObjectLog) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultResponseBodyResultObjectLog) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultResponseBodyResultObjectLog) GetOperateType() *string {
	return s.OperateType
}

func (s *GetSubTaskResultResponseBodyResultObjectLog) GetReason() *string {
	return s.Reason
}

func (s *GetSubTaskResultResponseBodyResultObjectLog) GetTime() *int64 {
	return s.Time
}

func (s *GetSubTaskResultResponseBodyResultObjectLog) SetOperateType(v string) *GetSubTaskResultResponseBodyResultObjectLog {
	s.OperateType = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectLog) SetReason(v string) *GetSubTaskResultResponseBodyResultObjectLog {
	s.Reason = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectLog) SetTime(v int64) *GetSubTaskResultResponseBodyResultObjectLog {
	s.Time = &v
	return s
}

func (s *GetSubTaskResultResponseBodyResultObjectLog) Validate() error {
	return dara.Validate(s)
}
