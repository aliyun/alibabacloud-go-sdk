// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppUploadProcessInMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserAppUploadProcessInMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetUserAppUploadProcessInMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *GetUserAppUploadProcessInMsaResponseBodyResultContent) *GetUserAppUploadProcessInMsaResponseBody
	GetResultContent() *GetUserAppUploadProcessInMsaResponseBodyResultContent
	SetResultMessage(v string) *GetUserAppUploadProcessInMsaResponseBody
	GetResultMessage() *string
}

type GetUserAppUploadProcessInMsaResponseBody struct {
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *GetUserAppUploadProcessInMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetUserAppUploadProcessInMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppUploadProcessInMsaResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAppUploadProcessInMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserAppUploadProcessInMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetUserAppUploadProcessInMsaResponseBody) GetResultContent() *GetUserAppUploadProcessInMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *GetUserAppUploadProcessInMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetUserAppUploadProcessInMsaResponseBody) SetRequestId(v string) *GetUserAppUploadProcessInMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBody) SetResultCode(v string) *GetUserAppUploadProcessInMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBody) SetResultContent(v *GetUserAppUploadProcessInMsaResponseBodyResultContent) *GetUserAppUploadProcessInMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBody) SetResultMessage(v string) *GetUserAppUploadProcessInMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAppUploadProcessInMsaResponseBodyResultContent struct {
	Code    *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetUserAppUploadProcessInMsaResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) GetData() *GetUserAppUploadProcessInMsaResponseBodyResultContentData {
	return s.Data
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) SetCode(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) SetData(v *GetUserAppUploadProcessInMsaResponseBodyResultContentData) *GetUserAppUploadProcessInMsaResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) SetMessage(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) SetSuccess(v bool) *GetUserAppUploadProcessInMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAppUploadProcessInMsaResponseBodyResultContentData struct {
	ApkInfo       *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo `json:"ApkInfo,omitempty" xml:"ApkInfo,omitempty" type:"Struct"`
	EnhanceTaskId *int64                                                            `json:"EnhanceTaskId,omitempty" xml:"EnhanceTaskId,omitempty"`
	Id            *int64                                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Progress      *int64                                                            `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status        *int64                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) GetApkInfo() *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	return s.ApkInfo
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) GetEnhanceTaskId() *int64 {
	return s.EnhanceTaskId
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) GetId() *int64 {
	return s.Id
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) GetProgress() *int64 {
	return s.Progress
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) GetStatus() *int64 {
	return s.Status
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) SetApkInfo(v *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) *GetUserAppUploadProcessInMsaResponseBodyResultContentData {
	s.ApkInfo = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) SetEnhanceTaskId(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentData {
	s.EnhanceTaskId = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) SetId(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentData {
	s.Id = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) SetProgress(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentData {
	s.Progress = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) SetStatus(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentData {
	s.Status = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentData) Validate() error {
	if s.ApkInfo != nil {
		if err := s.ApkInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo struct {
	AfterMd5            *string                                                                           `json:"AfterMd5,omitempty" xml:"AfterMd5,omitempty"`
	AfterSize           *int64                                                                            `json:"AfterSize,omitempty" xml:"AfterSize,omitempty"`
	AppCode             *string                                                                           `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppPackage          *string                                                                           `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	AssetsFileList      []*string                                                                         `json:"AssetsFileList,omitempty" xml:"AssetsFileList,omitempty" type:"Repeated"`
	BeforeMd5           *string                                                                           `json:"BeforeMd5,omitempty" xml:"BeforeMd5,omitempty"`
	BeforeSize          *int64                                                                            `json:"BeforeSize,omitempty" xml:"BeforeSize,omitempty"`
	ClassForest         *string                                                                           `json:"ClassForest,omitempty" xml:"ClassForest,omitempty"`
	EnhanceMapping      []*GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping `json:"EnhanceMapping,omitempty" xml:"EnhanceMapping,omitempty" type:"Repeated"`
	EnhanceRules        []*string                                                                         `json:"EnhanceRules,omitempty" xml:"EnhanceRules,omitempty" type:"Repeated"`
	EnhancedAssetsFiles []*string                                                                         `json:"EnhancedAssetsFiles,omitempty" xml:"EnhancedAssetsFiles,omitempty" type:"Repeated"`
	EnhancedClasses     []*string                                                                         `json:"EnhancedClasses,omitempty" xml:"EnhancedClasses,omitempty" type:"Repeated"`
	EnhancedSoFiles     []*string                                                                         `json:"EnhancedSoFiles,omitempty" xml:"EnhancedSoFiles,omitempty" type:"Repeated"`
	Id                  *int64                                                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Label               *string                                                                           `json:"Label,omitempty" xml:"Label,omitempty"`
	Progress            *int64                                                                            `json:"Progress,omitempty" xml:"Progress,omitempty"`
	SoFileList          []*string                                                                         `json:"SoFileList,omitempty" xml:"SoFileList,omitempty" type:"Repeated"`
	Status              *int64                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType            *string                                                                           `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	VersionCode         *string                                                                           `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionName         *string                                                                           `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GoString() string {
	return s.String()
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetAfterMd5() *string {
	return s.AfterMd5
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetAfterSize() *int64 {
	return s.AfterSize
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetAppPackage() *string {
	return s.AppPackage
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetAssetsFileList() []*string {
	return s.AssetsFileList
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetBeforeMd5() *string {
	return s.BeforeMd5
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetBeforeSize() *int64 {
	return s.BeforeSize
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetClassForest() *string {
	return s.ClassForest
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetEnhanceMapping() []*GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	return s.EnhanceMapping
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetEnhanceRules() []*string {
	return s.EnhanceRules
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetEnhancedAssetsFiles() []*string {
	return s.EnhancedAssetsFiles
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetEnhancedClasses() []*string {
	return s.EnhancedClasses
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetEnhancedSoFiles() []*string {
	return s.EnhancedSoFiles
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetId() *int64 {
	return s.Id
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetLabel() *string {
	return s.Label
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetProgress() *int64 {
	return s.Progress
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetSoFileList() []*string {
	return s.SoFileList
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetStatus() *int64 {
	return s.Status
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetTaskType() *string {
	return s.TaskType
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) GetVersionName() *string {
	return s.VersionName
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetAfterMd5(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.AfterMd5 = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetAfterSize(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.AfterSize = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetAppCode(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.AppCode = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetAppPackage(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.AppPackage = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetAssetsFileList(v []*string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.AssetsFileList = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetBeforeMd5(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.BeforeMd5 = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetBeforeSize(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.BeforeSize = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetClassForest(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.ClassForest = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetEnhanceMapping(v []*GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.EnhanceMapping = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetEnhanceRules(v []*string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.EnhanceRules = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetEnhancedAssetsFiles(v []*string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.EnhancedAssetsFiles = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetEnhancedClasses(v []*string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.EnhancedClasses = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetEnhancedSoFiles(v []*string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.EnhancedSoFiles = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetId(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.Id = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetLabel(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.Label = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetProgress(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.Progress = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetSoFileList(v []*string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.SoFileList = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetStatus(v int64) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.Status = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetTaskType(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.TaskType = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetVersionCode(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.VersionCode = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) SetVersionName(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo {
	s.VersionName = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfo) Validate() error {
	if s.EnhanceMapping != nil {
		for _, item := range s.EnhanceMapping {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping struct {
	Info   *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GoString() string {
	return s.String()
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GetInfo() *string {
	return s.Info
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GetReason() *string {
	return s.Reason
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) GetType() *string {
	return s.Type
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) SetInfo(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	s.Info = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) SetReason(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	s.Reason = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) SetType(v string) *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping {
	s.Type = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponseBodyResultContentDataApkInfoEnhanceMapping) Validate() error {
	return dara.Validate(s)
}
