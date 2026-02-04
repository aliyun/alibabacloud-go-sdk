// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceModuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *CheckInstanceModuleStatusResponseBodyModule) *CheckInstanceModuleStatusResponseBody
	GetModule() *CheckInstanceModuleStatusResponseBodyModule
	SetRequestId(v string) *CheckInstanceModuleStatusResponseBody
	GetRequestId() *string
}

type CheckInstanceModuleStatusResponseBody struct {
	Module *CheckInstanceModuleStatusResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInstanceModuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceModuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceModuleStatusResponseBody) GetModule() *CheckInstanceModuleStatusResponseBodyModule {
	return s.Module
}

func (s *CheckInstanceModuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstanceModuleStatusResponseBody) SetModule(v *CheckInstanceModuleStatusResponseBodyModule) *CheckInstanceModuleStatusResponseBody {
	s.Module = v
	return s
}

func (s *CheckInstanceModuleStatusResponseBody) SetRequestId(v string) *CheckInstanceModuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceModuleStatusResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckInstanceModuleStatusResponseBodyModule struct {
	// 模块状态
	//
	// example:
	//
	// enabled
	ModuleStatus *string `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
}

func (s CheckInstanceModuleStatusResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceModuleStatusResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckInstanceModuleStatusResponseBodyModule) GetModuleStatus() *string {
	return s.ModuleStatus
}

func (s *CheckInstanceModuleStatusResponseBodyModule) SetModuleStatus(v string) *CheckInstanceModuleStatusResponseBodyModule {
	s.ModuleStatus = &v
	return s
}

func (s *CheckInstanceModuleStatusResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
