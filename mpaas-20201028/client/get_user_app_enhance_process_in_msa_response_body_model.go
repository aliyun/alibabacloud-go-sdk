// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppEnhanceProcessInMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserAppEnhanceProcessInMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetUserAppEnhanceProcessInMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) *GetUserAppEnhanceProcessInMsaResponseBody
	GetResultContent() *GetUserAppEnhanceProcessInMsaResponseBodyResultContent
	SetResultMessage(v string) *GetUserAppEnhanceProcessInMsaResponseBody
	GetResultMessage() *string
}

type GetUserAppEnhanceProcessInMsaResponseBody struct {
	RequestId     *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *GetUserAppEnhanceProcessInMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetUserAppEnhanceProcessInMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppEnhanceProcessInMsaResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) GetResultContent() *GetUserAppEnhanceProcessInMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) SetRequestId(v string) *GetUserAppEnhanceProcessInMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) SetResultCode(v string) *GetUserAppEnhanceProcessInMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) SetResultContent(v *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) *GetUserAppEnhanceProcessInMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) SetResultMessage(v string) *GetUserAppEnhanceProcessInMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAppEnhanceProcessInMsaResponseBodyResultContent struct {
	Code    *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserAppEnhanceProcessInMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppEnhanceProcessInMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) GetData() *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	return s.Data
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) SetCode(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) SetData(v *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) *GetUserAppEnhanceProcessInMsaResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) SetMessage(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) SetSuccess(v bool) *GetUserAppEnhanceProcessInMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAppEnhanceProcessInMsaResponseBodyResultContentData struct {
	AfterMd5            *string                                                                     `json:"AfterMd5,omitempty" xml:"AfterMd5,omitempty"`
	AfterSize           *int64                                                                      `json:"AfterSize,omitempty" xml:"AfterSize,omitempty"`
	AppCode             *string                                                                     `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppPackage          *string                                                                     `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	AssetsFileList      []*string                                                                   `json:"AssetsFileList,omitempty" xml:"AssetsFileList,omitempty" type:"Repeated"`
	BeforeMd5           *string                                                                     `json:"BeforeMd5,omitempty" xml:"BeforeMd5,omitempty"`
	BeforeSize          *int64                                                                      `json:"BeforeSize,omitempty" xml:"BeforeSize,omitempty"`
	ClassForest         []*string                                                                   `json:"ClassForest,omitempty" xml:"ClassForest,omitempty" type:"Repeated"`
	EnhanceMapping      []*GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping `json:"EnhanceMapping,omitempty" xml:"EnhanceMapping,omitempty" type:"Repeated"`
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

func (s GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetAfterMd5() *string {
	return s.AfterMd5
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetAfterSize() *int64 {
	return s.AfterSize
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetAppCode() *string {
	return s.AppCode
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetAppPackage() *string {
	return s.AppPackage
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetAssetsFileList() []*string {
	return s.AssetsFileList
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetBeforeMd5() *string {
	return s.BeforeMd5
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetBeforeSize() *int64 {
	return s.BeforeSize
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetClassForest() []*string {
	return s.ClassForest
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetEnhanceMapping() []*GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping {
	return s.EnhanceMapping
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetEnhanceRules() []*string {
	return s.EnhanceRules
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetEnhancedAssetsFiles() []*string {
	return s.EnhancedAssetsFiles
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetEnhancedClasses() []*string {
	return s.EnhancedClasses
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetEnhancedSoFiles() []*string {
	return s.EnhancedSoFiles
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetId() *int64 {
	return s.Id
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetLabel() *string {
	return s.Label
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetProgress() *int64 {
	return s.Progress
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetSoFileList() []*string {
	return s.SoFileList
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetStatus() *int64 {
	return s.Status
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) GetVersionName() *string {
	return s.VersionName
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetAfterMd5(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.AfterMd5 = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetAfterSize(v int64) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.AfterSize = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetAppCode(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.AppCode = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetAppPackage(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.AppPackage = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetAssetsFileList(v []*string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.AssetsFileList = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetBeforeMd5(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.BeforeMd5 = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetBeforeSize(v int64) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.BeforeSize = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetClassForest(v []*string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.ClassForest = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetEnhanceMapping(v []*GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.EnhanceMapping = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetEnhanceRules(v []*string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.EnhanceRules = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetEnhancedAssetsFiles(v []*string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.EnhancedAssetsFiles = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetEnhancedClasses(v []*string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.EnhancedClasses = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetEnhancedSoFiles(v []*string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.EnhancedSoFiles = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetId(v int64) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.Id = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetLabel(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.Label = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetProgress(v int64) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.Progress = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetSoFileList(v []*string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.SoFileList = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetStatus(v int64) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.Status = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetTaskType(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.TaskType = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetVersionCode(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.VersionCode = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) SetVersionName(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData {
	s.VersionName = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentData) Validate() error {
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

type GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping struct {
	Info   *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) GoString() string {
	return s.String()
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) GetInfo() *string {
	return s.Info
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) GetReason() *string {
	return s.Reason
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) GetType() *string {
	return s.Type
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) SetInfo(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping {
	s.Info = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) SetReason(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping {
	s.Reason = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) SetType(v string) *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping {
	s.Type = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponseBodyResultContentDataEnhanceMapping) Validate() error {
	return dara.Validate(s)
}
