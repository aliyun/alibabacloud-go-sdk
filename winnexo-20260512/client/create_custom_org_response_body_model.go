// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCustomOrgResponseBody
	GetCode() *string
	SetCorpId(v string) *CreateCustomOrgResponseBody
	GetCorpId() *string
	SetCorpName(v string) *CreateCustomOrgResponseBody
	GetCorpName() *string
	SetMessage(v string) *CreateCustomOrgResponseBody
	GetMessage() *string
	SetPlatformType(v string) *CreateCustomOrgResponseBody
	GetPlatformType() *string
	SetRequestId(v string) *CreateCustomOrgResponseBody
	GetRequestId() *string
}

type CreateCustomOrgResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 注册成功的组织标识
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 组织展示名称
	//
	// example:
	//
	// string_value
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 平台类型，固定为 custom
	//
	// example:
	//
	// string_value
	PlatformType *string `json:"platformType,omitempty" xml:"platformType,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateCustomOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomOrgResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomOrgResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomOrgResponseBody) GetCorpId() *string {
	return s.CorpId
}

func (s *CreateCustomOrgResponseBody) GetCorpName() *string {
	return s.CorpName
}

func (s *CreateCustomOrgResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomOrgResponseBody) GetPlatformType() *string {
	return s.PlatformType
}

func (s *CreateCustomOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomOrgResponseBody) SetCode(v string) *CreateCustomOrgResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomOrgResponseBody) SetCorpId(v string) *CreateCustomOrgResponseBody {
	s.CorpId = &v
	return s
}

func (s *CreateCustomOrgResponseBody) SetCorpName(v string) *CreateCustomOrgResponseBody {
	s.CorpName = &v
	return s
}

func (s *CreateCustomOrgResponseBody) SetMessage(v string) *CreateCustomOrgResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomOrgResponseBody) SetPlatformType(v string) *CreateCustomOrgResponseBody {
	s.PlatformType = &v
	return s
}

func (s *CreateCustomOrgResponseBody) SetRequestId(v string) *CreateCustomOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomOrgResponseBody) Validate() error {
	return dara.Validate(s)
}
