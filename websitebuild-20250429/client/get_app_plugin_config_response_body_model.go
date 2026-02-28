// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAppPluginConfigResponseBody
	GetCode() *string
	SetMessage(v string) *GetAppPluginConfigResponseBody
	GetMessage() *string
	SetModule(v *GetAppPluginConfigResponseBodyModule) *GetAppPluginConfigResponseBody
	GetModule() *GetAppPluginConfigResponseBodyModule
	SetRequestId(v string) *GetAppPluginConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppPluginConfigResponseBody
	GetSuccess() *bool
}

type GetAppPluginConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Instance `wget h33E1En5.popscan.xaliyun.com` does not exist.
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Module  *GetAppPluginConfigResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppPluginConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAppPluginConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppPluginConfigResponseBody) GetModule() *GetAppPluginConfigResponseBodyModule {
	return s.Module
}

func (s *GetAppPluginConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppPluginConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppPluginConfigResponseBody) SetCode(v string) *GetAppPluginConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppPluginConfigResponseBody) SetMessage(v string) *GetAppPluginConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppPluginConfigResponseBody) SetModule(v *GetAppPluginConfigResponseBodyModule) *GetAppPluginConfigResponseBody {
	s.Module = v
	return s
}

func (s *GetAppPluginConfigResponseBody) SetRequestId(v string) *GetAppPluginConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppPluginConfigResponseBody) SetSuccess(v bool) *GetAppPluginConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppPluginConfigResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppPluginConfigResponseBodyModule struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 16257
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// example:
	//
	// a simple test plugin
	PluginDesc *string `json:"PluginDesc,omitempty" xml:"PluginDesc,omitempty"`
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// example:
	//
	// tf_testaccapigatewayplugin29311
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAppPluginConfigResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetAppPluginConfigResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppPluginConfigResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppPluginConfigResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *GetAppPluginConfigResponseBodyModule) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *GetAppPluginConfigResponseBodyModule) GetPluginDesc() *string {
	return s.PluginDesc
}

func (s *GetAppPluginConfigResponseBodyModule) GetPluginId() *string {
	return s.PluginId
}

func (s *GetAppPluginConfigResponseBodyModule) GetPluginName() *string {
	return s.PluginName
}

func (s *GetAppPluginConfigResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *GetAppPluginConfigResponseBodyModule) SetBizId(v string) *GetAppPluginConfigResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetGmtCreate(v string) *GetAppPluginConfigResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetGmtModified(v string) *GetAppPluginConfigResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetId(v int64) *GetAppPluginConfigResponseBodyModule {
	s.Id = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetPluginConfig(v string) *GetAppPluginConfigResponseBodyModule {
	s.PluginConfig = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetPluginDesc(v string) *GetAppPluginConfigResponseBodyModule {
	s.PluginDesc = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetPluginId(v string) *GetAppPluginConfigResponseBodyModule {
	s.PluginId = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetPluginName(v string) *GetAppPluginConfigResponseBodyModule {
	s.PluginName = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) SetUserId(v string) *GetAppPluginConfigResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *GetAppPluginConfigResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
