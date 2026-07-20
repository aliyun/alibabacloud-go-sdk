// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitSasModuleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitSasModuleRuleResponseBody
	GetCode() *string
	SetData(v *InitSasModuleRuleResponseBodyData) *InitSasModuleRuleResponseBody
	GetData() *InitSasModuleRuleResponseBodyData
	SetMessage(v string) *InitSasModuleRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitSasModuleRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InitSasModuleRuleResponseBody
	GetSuccess() *bool
}

type InitSasModuleRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InitSasModuleRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2E130B0F-9E69-52FA-84FC-187FE1BA9489
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitSasModuleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitSasModuleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InitSasModuleRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitSasModuleRuleResponseBody) GetData() *InitSasModuleRuleResponseBodyData {
	return s.Data
}

func (s *InitSasModuleRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitSasModuleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitSasModuleRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InitSasModuleRuleResponseBody) SetCode(v string) *InitSasModuleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *InitSasModuleRuleResponseBody) SetData(v *InitSasModuleRuleResponseBodyData) *InitSasModuleRuleResponseBody {
	s.Data = v
	return s
}

func (s *InitSasModuleRuleResponseBody) SetMessage(v string) *InitSasModuleRuleResponseBody {
	s.Message = &v
	return s
}

func (s *InitSasModuleRuleResponseBody) SetRequestId(v string) *InitSasModuleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitSasModuleRuleResponseBody) SetSuccess(v bool) *InitSasModuleRuleResponseBody {
	s.Success = &v
	return s
}

func (s *InitSasModuleRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitSasModuleRuleResponseBodyData struct {
	// example:
	//
	// 14492571-0707-5130-85B4-4DDABB6BDF76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitSasModuleRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InitSasModuleRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitSasModuleRuleResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *InitSasModuleRuleResponseBodyData) SetRequestId(v string) *InitSasModuleRuleResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *InitSasModuleRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
