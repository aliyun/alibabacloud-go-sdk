// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppCommoditySpecificationsForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListAppCommoditySpecificationsForPartnerResponseBodyModule) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetModule() *ListAppCommoditySpecificationsForPartnerResponseBodyModule
	SetRequestId(v string) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppCommoditySpecificationsForPartnerResponseBody
	GetSynchro() *bool
}

type ListAppCommoditySpecificationsForPartnerResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                               `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *ListAppCommoditySpecificationsForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppCommoditySpecificationsForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetModule() *ListAppCommoditySpecificationsForPartnerResponseBodyModule {
	return s.Module
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetAccessDeniedDetail(v string) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetAllowRetry(v bool) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetAppName(v string) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetDynamicCode(v string) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetDynamicMessage(v string) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetErrorArgs(v []interface{}) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetModule(v *ListAppCommoditySpecificationsForPartnerResponseBodyModule) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetRequestId(v string) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetRootErrorCode(v string) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetRootErrorMsg(v string) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) SetSynchro(v bool) *ListAppCommoditySpecificationsForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppCommoditySpecificationsForPartnerResponseBodyModule struct {
	Versions map[string]*ModuleVersionsValue `json:"Versions,omitempty" xml:"Versions,omitempty"`
}

func (s ListAppCommoditySpecificationsForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBodyModule) GetVersions() map[string]*ModuleVersionsValue {
	return s.Versions
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBodyModule) SetVersions(v map[string]*ModuleVersionsValue) *ListAppCommoditySpecificationsForPartnerResponseBodyModule {
	s.Versions = v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
