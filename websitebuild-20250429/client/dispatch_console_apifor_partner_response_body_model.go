// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDispatchConsoleAPIForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DispatchConsoleAPIForPartnerResponseBody
	GetErrorCode() *string
	SetModule(v *DispatchConsoleAPIForPartnerResponseBodyModule) *DispatchConsoleAPIForPartnerResponseBody
	GetModule() *DispatchConsoleAPIForPartnerResponseBodyModule
	SetRequestId(v string) *DispatchConsoleAPIForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DispatchConsoleAPIForPartnerResponseBody
	GetSuccess() *string
}

type DispatchConsoleAPIForPartnerResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Module    *DispatchConsoleAPIForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DispatchConsoleAPIForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DispatchConsoleAPIForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *DispatchConsoleAPIForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DispatchConsoleAPIForPartnerResponseBody) GetModule() *DispatchConsoleAPIForPartnerResponseBodyModule {
	return s.Module
}

func (s *DispatchConsoleAPIForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DispatchConsoleAPIForPartnerResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DispatchConsoleAPIForPartnerResponseBody) SetErrorCode(v string) *DispatchConsoleAPIForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponseBody) SetModule(v *DispatchConsoleAPIForPartnerResponseBodyModule) *DispatchConsoleAPIForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponseBody) SetRequestId(v string) *DispatchConsoleAPIForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponseBody) SetSuccess(v string) *DispatchConsoleAPIForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DispatchConsoleAPIForPartnerResponseBodyModule struct {
	// example:
	//
	// {\\"HasCustomRoleAuth\\": False}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DispatchConsoleAPIForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s DispatchConsoleAPIForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *DispatchConsoleAPIForPartnerResponseBodyModule) GetData() *string {
	return s.Data
}

func (s *DispatchConsoleAPIForPartnerResponseBodyModule) SetData(v string) *DispatchConsoleAPIForPartnerResponseBodyModule {
	s.Data = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
