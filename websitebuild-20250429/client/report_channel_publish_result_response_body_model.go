// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportChannelPublishResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReportChannelPublishResultResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ReportChannelPublishResultResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ReportChannelPublishResultResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ReportChannelPublishResultResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ReportChannelPublishResultResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ReportChannelPublishResultResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ReportChannelPublishResultResponseBodyModule) *ReportChannelPublishResultResponseBody
	GetModule() *ReportChannelPublishResultResponseBodyModule
	SetRequestId(v string) *ReportChannelPublishResultResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ReportChannelPublishResultResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ReportChannelPublishResultResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ReportChannelPublishResultResponseBody
	GetSynchro() *bool
}

type ReportChannelPublishResultResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the %s placeholder in the ErrMessage response parameter.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *ReportChannelPublishResultResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
	//
	// example:
	//
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ReportChannelPublishResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportChannelPublishResultResponseBody) GoString() string {
	return s.String()
}

func (s *ReportChannelPublishResultResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReportChannelPublishResultResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ReportChannelPublishResultResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ReportChannelPublishResultResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ReportChannelPublishResultResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ReportChannelPublishResultResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ReportChannelPublishResultResponseBody) GetModule() *ReportChannelPublishResultResponseBodyModule {
	return s.Module
}

func (s *ReportChannelPublishResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportChannelPublishResultResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ReportChannelPublishResultResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ReportChannelPublishResultResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ReportChannelPublishResultResponseBody) SetAccessDeniedDetail(v string) *ReportChannelPublishResultResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetAllowRetry(v bool) *ReportChannelPublishResultResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetAppName(v string) *ReportChannelPublishResultResponseBody {
	s.AppName = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetDynamicCode(v string) *ReportChannelPublishResultResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetDynamicMessage(v string) *ReportChannelPublishResultResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetErrorArgs(v []interface{}) *ReportChannelPublishResultResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetModule(v *ReportChannelPublishResultResponseBodyModule) *ReportChannelPublishResultResponseBody {
	s.Module = v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetRequestId(v string) *ReportChannelPublishResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetRootErrorCode(v string) *ReportChannelPublishResultResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetRootErrorMsg(v string) *ReportChannelPublishResultResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) SetSynchro(v bool) *ReportChannelPublishResultResponseBody {
	s.Synchro = &v
	return s
}

func (s *ReportChannelPublishResultResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReportChannelPublishResultResponseBodyModule struct {
	// The article status.
	//
	// example:
	//
	// xxx
	ArticleStatus *string `json:"ArticleStatus,omitempty" xml:"ArticleStatus,omitempty"`
	// The channel draft ID.
	//
	// example:
	//
	// draft-789
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// The draft status. Valid values: SUCCESS and FAILED.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ReportChannelPublishResultResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ReportChannelPublishResultResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ReportChannelPublishResultResponseBodyModule) GetArticleStatus() *string {
	return s.ArticleStatus
}

func (s *ReportChannelPublishResultResponseBodyModule) GetDraftId() *string {
	return s.DraftId
}

func (s *ReportChannelPublishResultResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *ReportChannelPublishResultResponseBodyModule) SetArticleStatus(v string) *ReportChannelPublishResultResponseBodyModule {
	s.ArticleStatus = &v
	return s
}

func (s *ReportChannelPublishResultResponseBodyModule) SetDraftId(v string) *ReportChannelPublishResultResponseBodyModule {
	s.DraftId = &v
	return s
}

func (s *ReportChannelPublishResultResponseBodyModule) SetStatus(v string) *ReportChannelPublishResultResponseBodyModule {
	s.Status = &v
	return s
}

func (s *ReportChannelPublishResultResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
