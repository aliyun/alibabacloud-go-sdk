// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlidingAssistantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlidingAssistantId(v string) *CreateAlidingAssistantResponseBody
	GetAlidingAssistantId() *string
	SetAppCode(v string) *CreateAlidingAssistantResponseBody
	GetAppCode() *string
	SetJumpUrl(v string) *CreateAlidingAssistantResponseBody
	GetJumpUrl() *string
	SetProcessInstanceId(v string) *CreateAlidingAssistantResponseBody
	GetProcessInstanceId() *string
	SetRequestId(v string) *CreateAlidingAssistantResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *CreateAlidingAssistantResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateAlidingAssistantResponseBody
	GetVendorType() *string
}

type CreateAlidingAssistantResponseBody struct {
	// example:
	//
	// 123456
	AlidingAssistantId *string `json:"alidingAssistantId,omitempty" xml:"alidingAssistantId,omitempty"`
	AppCode            *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	// example:
	//
	// https://www.baidu.com
	JumpUrl           *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreateAlidingAssistantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantResponseBody) GetAlidingAssistantId() *string {
	return s.AlidingAssistantId
}

func (s *CreateAlidingAssistantResponseBody) GetAppCode() *string {
	return s.AppCode
}

func (s *CreateAlidingAssistantResponseBody) GetJumpUrl() *string {
	return s.JumpUrl
}

func (s *CreateAlidingAssistantResponseBody) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *CreateAlidingAssistantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlidingAssistantResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateAlidingAssistantResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateAlidingAssistantResponseBody) SetAlidingAssistantId(v string) *CreateAlidingAssistantResponseBody {
	s.AlidingAssistantId = &v
	return s
}

func (s *CreateAlidingAssistantResponseBody) SetAppCode(v string) *CreateAlidingAssistantResponseBody {
	s.AppCode = &v
	return s
}

func (s *CreateAlidingAssistantResponseBody) SetJumpUrl(v string) *CreateAlidingAssistantResponseBody {
	s.JumpUrl = &v
	return s
}

func (s *CreateAlidingAssistantResponseBody) SetProcessInstanceId(v string) *CreateAlidingAssistantResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *CreateAlidingAssistantResponseBody) SetRequestId(v string) *CreateAlidingAssistantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlidingAssistantResponseBody) SetVendorRequestId(v string) *CreateAlidingAssistantResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateAlidingAssistantResponseBody) SetVendorType(v string) *CreateAlidingAssistantResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateAlidingAssistantResponseBody) Validate() error {
	return dara.Validate(s)
}
