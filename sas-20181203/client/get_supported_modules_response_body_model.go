// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupportedModulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSupportedModulesResponseBody
	GetRequestId() *string
	SetSupportedModuleResponse(v []*GetSupportedModulesResponseBodySupportedModuleResponse) *GetSupportedModulesResponseBody
	GetSupportedModuleResponse() []*GetSupportedModulesResponseBodySupportedModuleResponse
}

type GetSupportedModulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C699E4E4-F2F4-58FC-A949-457FFE59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supported modules. The module information is classified by cloud service provider.
	SupportedModuleResponse []*GetSupportedModulesResponseBodySupportedModuleResponse `json:"SupportedModuleResponse,omitempty" xml:"SupportedModuleResponse,omitempty" type:"Repeated"`
}

func (s GetSupportedModulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedModulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupportedModulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupportedModulesResponseBody) GetSupportedModuleResponse() []*GetSupportedModulesResponseBodySupportedModuleResponse {
	return s.SupportedModuleResponse
}

func (s *GetSupportedModulesResponseBody) SetRequestId(v string) *GetSupportedModulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupportedModulesResponseBody) SetSupportedModuleResponse(v []*GetSupportedModulesResponseBodySupportedModuleResponse) *GetSupportedModulesResponseBody {
	s.SupportedModuleResponse = v
	return s
}

func (s *GetSupportedModulesResponseBody) Validate() error {
	if s.SupportedModuleResponse != nil {
		for _, item := range s.SupportedModuleResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSupportedModulesResponseBodySupportedModuleResponse struct {
	// The modules supported by the cloud service provider.
	SupportedModules []*GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules `json:"SupportedModules,omitempty" xml:"SupportedModules,omitempty" type:"Repeated"`
	// The cloud service provider. Valid values:
	//
	// 	- **Tencent**: Tencent Cloud
	//
	// 	- **HUAWEICLOUD**:Huawei Cloud
	//
	// 	- **Azure**: Microsoft Azure
	//
	// 	- **AWS**: Amazon Web Services (AWS)
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetSupportedModulesResponseBodySupportedModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedModulesResponseBodySupportedModuleResponse) GoString() string {
	return s.String()
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponse) GetSupportedModules() []*GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules {
	return s.SupportedModules
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponse) GetVendor() *string {
	return s.Vendor
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponse) SetSupportedModules(v []*GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) *GetSupportedModulesResponseBodySupportedModuleResponse {
	s.SupportedModules = v
	return s
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponse) SetVendor(v string) *GetSupportedModulesResponseBodySupportedModuleResponse {
	s.Vendor = &v
	return s
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponse) Validate() error {
	if s.SupportedModules != nil {
		for _, item := range s.SupportedModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules struct {
	// The code of the module. Valid values:
	//
	// 	- **HOST**: host
	//
	// 	- **CSPM**: configuration assessment
	//
	// 	- **SIEM**: CloudSiem
	//
	// 	- **TRIAL**: log audit
	//
	// example:
	//
	// HOST
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The display name of the module.
	//
	// example:
	//
	// Configuration assessment
	ModuleDisp *string `json:"ModuleDisp,omitempty" xml:"ModuleDisp,omitempty"`
}

func (s GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) GoString() string {
	return s.String()
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) GetModule() *string {
	return s.Module
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) GetModuleDisp() *string {
	return s.ModuleDisp
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) SetModule(v string) *GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules {
	s.Module = &v
	return s
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) SetModuleDisp(v string) *GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules {
	s.ModuleDisp = &v
	return s
}

func (s *GetSupportedModulesResponseBodySupportedModuleResponseSupportedModules) Validate() error {
	return dara.Validate(s)
}
