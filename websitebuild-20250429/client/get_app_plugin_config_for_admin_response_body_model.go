// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginConfigForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAppPluginConfigForAdminResponseBody
	GetCode() *string
	SetMessage(v string) *GetAppPluginConfigForAdminResponseBody
	GetMessage() *string
	SetModule(v *GetAppPluginConfigForAdminResponseBodyModule) *GetAppPluginConfigForAdminResponseBody
	GetModule() *GetAppPluginConfigForAdminResponseBodyModule
	SetRequestId(v string) *GetAppPluginConfigForAdminResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppPluginConfigForAdminResponseBody
	GetSuccess() *bool
}

type GetAppPluginConfigForAdminResponseBody struct {
	// The operation status or POP error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The response data.
	//
	// example:
	//
	// true
	Module *GetAppPluginConfigForAdminResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppPluginConfigForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigForAdminResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAppPluginConfigForAdminResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppPluginConfigForAdminResponseBody) GetModule() *GetAppPluginConfigForAdminResponseBodyModule {
	return s.Module
}

func (s *GetAppPluginConfigForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppPluginConfigForAdminResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppPluginConfigForAdminResponseBody) SetCode(v string) *GetAppPluginConfigForAdminResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBody) SetMessage(v string) *GetAppPluginConfigForAdminResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBody) SetModule(v *GetAppPluginConfigForAdminResponseBodyModule) *GetAppPluginConfigForAdminResponseBody {
	s.Module = v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBody) SetRequestId(v string) *GetAppPluginConfigForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBody) SetSuccess(v bool) *GetAppPluginConfigForAdminResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppPluginConfigForAdminResponseBodyModule struct {
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key.
	//
	// example:
	//
	// 16257
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The plugin configuration.
	//
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// The plugin description.
	//
	// example:
	//
	// a simple test plugin
	PluginDesc *string `json:"PluginDesc,omitempty" xml:"PluginDesc,omitempty"`
	// The plugin ID.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plugin name.
	//
	// example:
	//
	// tf_testaccapigatewayplugin29311
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAppPluginConfigForAdminResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigForAdminResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetPluginDesc() *string {
	return s.PluginDesc
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetPluginId() *string {
	return s.PluginId
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetPluginName() *string {
	return s.PluginName
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetBizId(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetGmtCreate(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetGmtModified(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetId(v int64) *GetAppPluginConfigForAdminResponseBodyModule {
	s.Id = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetPluginConfig(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.PluginConfig = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetPluginDesc(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.PluginDesc = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetPluginId(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.PluginId = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetPluginName(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.PluginName = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) SetUserId(v string) *GetAppPluginConfigForAdminResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *GetAppPluginConfigForAdminResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
