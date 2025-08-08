// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadUserAppToMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadUserAppToMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UploadUserAppToMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *UploadUserAppToMsaResponseBodyResultContent) *UploadUserAppToMsaResponseBody
	GetResultContent() *UploadUserAppToMsaResponseBodyResultContent
	SetResultMessage(v string) *UploadUserAppToMsaResponseBody
	GetResultMessage() *string
}

type UploadUserAppToMsaResponseBody struct {
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *UploadUserAppToMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UploadUserAppToMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadUserAppToMsaResponseBody) GoString() string {
	return s.String()
}

func (s *UploadUserAppToMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadUserAppToMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UploadUserAppToMsaResponseBody) GetResultContent() *UploadUserAppToMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *UploadUserAppToMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UploadUserAppToMsaResponseBody) SetRequestId(v string) *UploadUserAppToMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadUserAppToMsaResponseBody) SetResultCode(v string) *UploadUserAppToMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UploadUserAppToMsaResponseBody) SetResultContent(v *UploadUserAppToMsaResponseBodyResultContent) *UploadUserAppToMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *UploadUserAppToMsaResponseBody) SetResultMessage(v string) *UploadUserAppToMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UploadUserAppToMsaResponseBody) Validate() error {
	return dara.Validate(s)
}

type UploadUserAppToMsaResponseBodyResultContent struct {
	Code    *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *UploadUserAppToMsaResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Extra   *string                                          `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Message *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadUserAppToMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s UploadUserAppToMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *UploadUserAppToMsaResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *UploadUserAppToMsaResponseBodyResultContent) GetData() *UploadUserAppToMsaResponseBodyResultContentData {
	return s.Data
}

func (s *UploadUserAppToMsaResponseBodyResultContent) GetExtra() *string {
	return s.Extra
}

func (s *UploadUserAppToMsaResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *UploadUserAppToMsaResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *UploadUserAppToMsaResponseBodyResultContent) SetCode(v string) *UploadUserAppToMsaResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContent) SetData(v *UploadUserAppToMsaResponseBodyResultContentData) *UploadUserAppToMsaResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContent) SetExtra(v string) *UploadUserAppToMsaResponseBodyResultContent {
	s.Extra = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContent) SetMessage(v string) *UploadUserAppToMsaResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContent) SetSuccess(v bool) *UploadUserAppToMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type UploadUserAppToMsaResponseBodyResultContentData struct {
	ApkInfo       *UploadUserAppToMsaResponseBodyResultContentDataApkInfo `json:"ApkInfo,omitempty" xml:"ApkInfo,omitempty" type:"Struct"`
	EnhanceTaskId *int64                                                  `json:"EnhanceTaskId,omitempty" xml:"EnhanceTaskId,omitempty"`
	Id            *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Progress      *int64                                                  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status        *int64                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UploadUserAppToMsaResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s UploadUserAppToMsaResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) GetApkInfo() *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	return s.ApkInfo
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) GetEnhanceTaskId() *int64 {
	return s.EnhanceTaskId
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) GetId() *int64 {
	return s.Id
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) GetProgress() *int64 {
	return s.Progress
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) GetStatus() *int64 {
	return s.Status
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) SetApkInfo(v *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) *UploadUserAppToMsaResponseBodyResultContentData {
	s.ApkInfo = v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) SetEnhanceTaskId(v int64) *UploadUserAppToMsaResponseBodyResultContentData {
	s.EnhanceTaskId = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) SetId(v int64) *UploadUserAppToMsaResponseBodyResultContentData {
	s.Id = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) SetProgress(v int64) *UploadUserAppToMsaResponseBodyResultContentData {
	s.Progress = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) SetStatus(v int64) *UploadUserAppToMsaResponseBodyResultContentData {
	s.Status = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}

type UploadUserAppToMsaResponseBodyResultContentDataApkInfo struct {
	AfterMd5        *string                                                               `json:"AfterMd5,omitempty" xml:"AfterMd5,omitempty"`
	AfterSize       *int64                                                                `json:"AfterSize,omitempty" xml:"AfterSize,omitempty"`
	AppCode         *string                                                               `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppPackage      *string                                                               `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	BeforeMd5       *string                                                               `json:"BeforeMd5,omitempty" xml:"BeforeMd5,omitempty"`
	BeforeSize      *int64                                                                `json:"BeforeSize,omitempty" xml:"BeforeSize,omitempty"`
	ClassForest     *string                                                               `json:"ClassForest,omitempty" xml:"ClassForest,omitempty"`
	EnhanceMapping  *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping `json:"EnhanceMapping,omitempty" xml:"EnhanceMapping,omitempty" type:"Struct"`
	EnhanceRules    []*string                                                             `json:"EnhanceRules,omitempty" xml:"EnhanceRules,omitempty" type:"Repeated"`
	EnhancedClasses []*string                                                             `json:"EnhancedClasses,omitempty" xml:"EnhancedClasses,omitempty" type:"Repeated"`
	Id              *int64                                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	Label           *string                                                               `json:"Label,omitempty" xml:"Label,omitempty"`
	Progress        *int64                                                                `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status          *int64                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType        *string                                                               `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	VersionCode     *string                                                               `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionName     *string                                                               `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UploadUserAppToMsaResponseBodyResultContentDataApkInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GoString() string {
	return s.String()
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetAfterMd5() *string {
	return s.AfterMd5
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetAfterSize() *int64 {
	return s.AfterSize
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetAppPackage() *string {
	return s.AppPackage
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetBeforeMd5() *string {
	return s.BeforeMd5
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetBeforeSize() *int64 {
	return s.BeforeSize
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetClassForest() *string {
	return s.ClassForest
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetEnhanceMapping() *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	return s.EnhanceMapping
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetEnhanceRules() []*string {
	return s.EnhanceRules
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetEnhancedClasses() []*string {
	return s.EnhancedClasses
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetId() *int64 {
	return s.Id
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetLabel() *string {
	return s.Label
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetProgress() *int64 {
	return s.Progress
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetStatus() *int64 {
	return s.Status
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetTaskType() *string {
	return s.TaskType
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetVersionCode() *string {
	return s.VersionCode
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) GetVersionName() *string {
	return s.VersionName
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetAfterMd5(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.AfterMd5 = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetAfterSize(v int64) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.AfterSize = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetAppCode(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.AppCode = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetAppPackage(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.AppPackage = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetBeforeMd5(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.BeforeMd5 = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetBeforeSize(v int64) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.BeforeSize = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetClassForest(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.ClassForest = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetEnhanceMapping(v *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.EnhanceMapping = v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetEnhanceRules(v []*string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.EnhanceRules = v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetEnhancedClasses(v []*string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.EnhancedClasses = v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetId(v int64) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.Id = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetLabel(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.Label = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetProgress(v int64) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.Progress = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetStatus(v int64) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.Status = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetTaskType(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.TaskType = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetVersionCode(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.VersionCode = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) SetVersionName(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfo {
	s.VersionName = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfo) Validate() error {
	return dara.Validate(s)
}

type UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping struct {
	Info   *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) String() string {
	return dara.Prettify(s)
}

func (s UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GoString() string {
	return s.String()
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GetInfo() *string {
	return s.Info
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GetReason() *string {
	return s.Reason
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GetType() *string {
	return s.Type
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) SetInfo(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	s.Info = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) SetReason(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	s.Reason = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) SetType(v string) *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	s.Type = &v
	return s
}

func (s *UploadUserAppToMsaResponseBodyResultContentDataApkInfoEnhanceMapping) Validate() error {
	return dara.Validate(s)
}
