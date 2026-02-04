// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceGlobalizationConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetInstanceGlobalizationConfigRequest
	GetInstanceId() *string
	SetLanguage(v string) *SetInstanceGlobalizationConfigRequest
	GetLanguage() *string
	SetTimeZone(v string) *SetInstanceGlobalizationConfigRequest
	GetTimeZone() *string
}

type SetInstanceGlobalizationConfigRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 语言类型
	//
	// This parameter is required.
	//
	// example:
	//
	// zh-Hans-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 时区
	//
	// This parameter is required.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s SetInstanceGlobalizationConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceGlobalizationConfigRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceGlobalizationConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetInstanceGlobalizationConfigRequest) GetLanguage() *string {
	return s.Language
}

func (s *SetInstanceGlobalizationConfigRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *SetInstanceGlobalizationConfigRequest) SetInstanceId(v string) *SetInstanceGlobalizationConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceGlobalizationConfigRequest) SetLanguage(v string) *SetInstanceGlobalizationConfigRequest {
	s.Language = &v
	return s
}

func (s *SetInstanceGlobalizationConfigRequest) SetTimeZone(v string) *SetInstanceGlobalizationConfigRequest {
	s.TimeZone = &v
	return s
}

func (s *SetInstanceGlobalizationConfigRequest) Validate() error {
	return dara.Validate(s)
}
