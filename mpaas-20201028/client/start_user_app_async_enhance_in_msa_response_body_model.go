// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartUserAppAsyncEnhanceInMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartUserAppAsyncEnhanceInMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *StartUserAppAsyncEnhanceInMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) *StartUserAppAsyncEnhanceInMsaResponseBody
	GetResultContent() *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent
	SetResultMessage(v string) *StartUserAppAsyncEnhanceInMsaResponseBody
	GetResultMessage() *string
}

type StartUserAppAsyncEnhanceInMsaResponseBody struct {
	RequestId     *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s StartUserAppAsyncEnhanceInMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartUserAppAsyncEnhanceInMsaResponseBody) GoString() string {
	return s.String()
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) GetResultContent() *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) SetRequestId(v string) *StartUserAppAsyncEnhanceInMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) SetResultCode(v string) *StartUserAppAsyncEnhanceInMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) SetResultContent(v *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) *StartUserAppAsyncEnhanceInMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) SetResultMessage(v string) *StartUserAppAsyncEnhanceInMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartUserAppAsyncEnhanceInMsaResponseBodyResultContent struct {
	Code    *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) GetData() *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	return s.Data
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) SetCode(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) SetData(v *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) SetMessage(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) SetSuccess(v bool) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData struct {
	AfterMd5            *string                                                                     `json:"AfterMd5,omitempty" xml:"AfterMd5,omitempty"`
	AfterSize           *int64                                                                      `json:"AfterSize,omitempty" xml:"AfterSize,omitempty"`
	AppCode             *string                                                                     `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppPackage          *string                                                                     `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	AssetsFileList      []*string                                                                   `json:"AssetsFileList,omitempty" xml:"AssetsFileList,omitempty" type:"Repeated"`
	BeforeMd5           *string                                                                     `json:"BeforeMd5,omitempty" xml:"BeforeMd5,omitempty"`
	BeforeSize          *int64                                                                      `json:"BeforeSize,omitempty" xml:"BeforeSize,omitempty"`
	ClassForest         *string                                                                     `json:"ClassForest,omitempty" xml:"ClassForest,omitempty"`
	EnhanceMapping      []*StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping `json:"EnhanceMapping,omitempty" xml:"EnhanceMapping,omitempty" type:"Repeated"`
	EnhanceRules        []*string                                                                   `json:"EnhanceRules,omitempty" xml:"EnhanceRules,omitempty" type:"Repeated"`
	EnhancedAssetsFiles []*string                                                                   `json:"EnhancedAssetsFiles,omitempty" xml:"EnhancedAssetsFiles,omitempty" type:"Repeated"`
	EnhancedClasses     []*string                                                                   `json:"EnhancedClasses,omitempty" xml:"EnhancedClasses,omitempty" type:"Repeated"`
	EnhancedSoFiles     []*string                                                                   `json:"EnhancedSoFiles,omitempty" xml:"EnhancedSoFiles,omitempty" type:"Repeated"`
	Id                  *int64                                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Label               *string                                                                     `json:"Label,omitempty" xml:"Label,omitempty"`
	Progress            *int64                                                                      `json:"Progress,omitempty" xml:"Progress,omitempty"`
	SoFileList          []*string                                                                   `json:"SoFileList,omitempty" xml:"SoFileList,omitempty" type:"Repeated"`
	Status              *int64                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType            *string                                                                     `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	VersionCode         *string                                                                     `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionName         *string                                                                     `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetAfterMd5() *string {
	return s.AfterMd5
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetAfterSize() *int64 {
	return s.AfterSize
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetAppCode() *string {
	return s.AppCode
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetAppPackage() *string {
	return s.AppPackage
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetAssetsFileList() []*string {
	return s.AssetsFileList
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetBeforeMd5() *string {
	return s.BeforeMd5
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetBeforeSize() *int64 {
	return s.BeforeSize
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetClassForest() *string {
	return s.ClassForest
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetEnhanceMapping() []*StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping {
	return s.EnhanceMapping
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetEnhanceRules() []*string {
	return s.EnhanceRules
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetEnhancedAssetsFiles() []*string {
	return s.EnhancedAssetsFiles
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetEnhancedClasses() []*string {
	return s.EnhancedClasses
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetEnhancedSoFiles() []*string {
	return s.EnhancedSoFiles
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetId() *int64 {
	return s.Id
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetLabel() *string {
	return s.Label
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetProgress() *int64 {
	return s.Progress
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetSoFileList() []*string {
	return s.SoFileList
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetStatus() *int64 {
	return s.Status
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetTaskType() *string {
	return s.TaskType
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) GetVersionName() *string {
	return s.VersionName
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetAfterMd5(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.AfterMd5 = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetAfterSize(v int64) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.AfterSize = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetAppCode(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.AppCode = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetAppPackage(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.AppPackage = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetAssetsFileList(v []*string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.AssetsFileList = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetBeforeMd5(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.BeforeMd5 = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetBeforeSize(v int64) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.BeforeSize = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetClassForest(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.ClassForest = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetEnhanceMapping(v []*StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.EnhanceMapping = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetEnhanceRules(v []*string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.EnhanceRules = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetEnhancedAssetsFiles(v []*string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.EnhancedAssetsFiles = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetEnhancedClasses(v []*string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.EnhancedClasses = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetEnhancedSoFiles(v []*string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.EnhancedSoFiles = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetId(v int64) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.Id = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetLabel(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.Label = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetProgress(v int64) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.Progress = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetSoFileList(v []*string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.SoFileList = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetStatus(v int64) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.Status = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetTaskType(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.TaskType = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetVersionCode(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.VersionCode = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) SetVersionName(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData {
	s.VersionName = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}

type StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping struct {
	Info   *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) String() string {
	return dara.Prettify(s)
}

func (s StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) GoString() string {
	return s.String()
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) GetInfo() *string {
	return s.Info
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) GetReason() *string {
	return s.Reason
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) GetType() *string {
	return s.Type
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) SetInfo(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping {
	s.Info = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) SetReason(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping {
	s.Reason = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) SetType(v string) *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping {
	s.Type = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponseBodyResultContentDataEnhanceMapping) Validate() error {
	return dara.Validate(s)
}
