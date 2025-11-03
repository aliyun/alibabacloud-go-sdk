// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleForProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckServiceLinkedRoleForProductResponseBody
	GetCode() *string
	SetData(v *CheckServiceLinkedRoleForProductResponseBodyData) *CheckServiceLinkedRoleForProductResponseBody
	GetData() *CheckServiceLinkedRoleForProductResponseBodyData
	SetMessage(v string) *CheckServiceLinkedRoleForProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckServiceLinkedRoleForProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckServiceLinkedRoleForProductResponseBody
	GetSuccess() *bool
}

type CheckServiceLinkedRoleForProductResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CheckServiceLinkedRoleForProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckServiceLinkedRoleForProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForProductResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckServiceLinkedRoleForProductResponseBody) GetData() *CheckServiceLinkedRoleForProductResponseBodyData {
	return s.Data
}

func (s *CheckServiceLinkedRoleForProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckServiceLinkedRoleForProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceLinkedRoleForProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckServiceLinkedRoleForProductResponseBody) SetCode(v string) *CheckServiceLinkedRoleForProductResponseBody {
	s.Code = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponseBody) SetData(v *CheckServiceLinkedRoleForProductResponseBodyData) *CheckServiceLinkedRoleForProductResponseBody {
	s.Data = v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponseBody) SetMessage(v string) *CheckServiceLinkedRoleForProductResponseBody {
	s.Message = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleForProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponseBody) SetSuccess(v bool) *CheckServiceLinkedRoleForProductResponseBody {
	s.Success = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckServiceLinkedRoleForProductResponseBodyData struct {
	// Indicates whether the service-linked role exists.
	//
	// example:
	//
	// true
	CheckPass *bool `json:"CheckPass,omitempty" xml:"CheckPass,omitempty"`
	// The name of the service-linked role.
	//
	// example:
	//
	// AliyunServiceRoleForEventBridgeConnectVPC
	StsRoleName *string `json:"StsRoleName,omitempty" xml:"StsRoleName,omitempty"`
}

func (s CheckServiceLinkedRoleForProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForProductResponseBodyData) GetCheckPass() *bool {
	return s.CheckPass
}

func (s *CheckServiceLinkedRoleForProductResponseBodyData) GetStsRoleName() *string {
	return s.StsRoleName
}

func (s *CheckServiceLinkedRoleForProductResponseBodyData) SetCheckPass(v bool) *CheckServiceLinkedRoleForProductResponseBodyData {
	s.CheckPass = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponseBodyData) SetStsRoleName(v string) *CheckServiceLinkedRoleForProductResponseBodyData {
	s.StsRoleName = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponseBodyData) Validate() error {
	return dara.Validate(s)
}
