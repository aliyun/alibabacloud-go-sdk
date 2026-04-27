// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudInterruptIvrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckName(v string) *CloudInterruptIvrRequest
	GetCheckName() *string
	SetCheckValue(v string) *CloudInterruptIvrRequest
	GetCheckValue() *string
	SetEnterpriseId(v int64) *CloudInterruptIvrRequest
	GetEnterpriseId() *int64
	SetUniqueId(v string) *CloudInterruptIvrRequest
	GetUniqueId() *string
	SetUserField(v string) *CloudInterruptIvrRequest
	GetUserField() *string
}

type CloudInterruptIvrRequest struct {
	// 根据变量名去通道变量中取对应的变量值
	//
	// example:
	//
	// name1
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// 当checkName代表的变量值与checkValue一致，才打断
	//
	// example:
	//
	// value1
	CheckValue *string `json:"CheckValue,omitempty" xml:"CheckValue,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 通话唯一标识；从通道唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// 自定义字段；json字符串
	//
	// example:
	//
	// {}
	UserField *string `json:"UserField,omitempty" xml:"UserField,omitempty"`
}

func (s CloudInterruptIvrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudInterruptIvrRequest) GoString() string {
	return s.String()
}

func (s *CloudInterruptIvrRequest) GetCheckName() *string {
	return s.CheckName
}

func (s *CloudInterruptIvrRequest) GetCheckValue() *string {
	return s.CheckValue
}

func (s *CloudInterruptIvrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudInterruptIvrRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudInterruptIvrRequest) GetUserField() *string {
	return s.UserField
}

func (s *CloudInterruptIvrRequest) SetCheckName(v string) *CloudInterruptIvrRequest {
	s.CheckName = &v
	return s
}

func (s *CloudInterruptIvrRequest) SetCheckValue(v string) *CloudInterruptIvrRequest {
	s.CheckValue = &v
	return s
}

func (s *CloudInterruptIvrRequest) SetEnterpriseId(v int64) *CloudInterruptIvrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudInterruptIvrRequest) SetUniqueId(v string) *CloudInterruptIvrRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudInterruptIvrRequest) SetUserField(v string) *CloudInterruptIvrRequest {
	s.UserField = &v
	return s
}

func (s *CloudInterruptIvrRequest) Validate() error {
	return dara.Validate(s)
}
