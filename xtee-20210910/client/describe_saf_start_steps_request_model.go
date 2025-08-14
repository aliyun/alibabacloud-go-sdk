// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafStartStepsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSafStartStepsRequest
	GetLang() *string
	SetAliyunServer(v bool) *DescribeSafStartStepsRequest
	GetAliyunServer() *bool
	SetDeviceTypesStr(v string) *DescribeSafStartStepsRequest
	GetDeviceTypesStr() *string
	SetEventCode(v string) *DescribeSafStartStepsRequest
	GetEventCode() *string
	SetLanguage(v string) *DescribeSafStartStepsRequest
	GetLanguage() *string
	SetRegId(v string) *DescribeSafStartStepsRequest
	GetRegId() *string
	SetServerRegion(v string) *DescribeSafStartStepsRequest
	GetServerRegion() *string
}

type DescribeSafStartStepsRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Whether the server is an Alibaba Cloud server
	//
	// true or false
	//
	// example:
	//
	// true
	AliyunServer *bool `json:"aliyunServer,omitempty" xml:"aliyunServer,omitempty"`
	// Used to receive a collection of strings from the frontend that POP cannot accept
	//
	//
	//
	// Device type
	//
	// example:
	//
	// ios
	DeviceTypesStr *string `json:"deviceTypesStr,omitempty" xml:"deviceTypesStr,omitempty"`
	// Event code
	//
	// example:
	//
	// de_ahqido8038
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Language, parameters can be passed
	//
	// - zh-CN: Chinese (default)
	//
	// - en-US: English
	//
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Server region
	//
	// example:
	//
	// cn-hangzhou
	ServerRegion *string `json:"serverRegion,omitempty" xml:"serverRegion,omitempty"`
}

func (s DescribeSafStartStepsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartStepsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSafStartStepsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSafStartStepsRequest) GetAliyunServer() *bool {
	return s.AliyunServer
}

func (s *DescribeSafStartStepsRequest) GetDeviceTypesStr() *string {
	return s.DeviceTypesStr
}

func (s *DescribeSafStartStepsRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSafStartStepsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeSafStartStepsRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSafStartStepsRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribeSafStartStepsRequest) SetLang(v string) *DescribeSafStartStepsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSafStartStepsRequest) SetAliyunServer(v bool) *DescribeSafStartStepsRequest {
	s.AliyunServer = &v
	return s
}

func (s *DescribeSafStartStepsRequest) SetDeviceTypesStr(v string) *DescribeSafStartStepsRequest {
	s.DeviceTypesStr = &v
	return s
}

func (s *DescribeSafStartStepsRequest) SetEventCode(v string) *DescribeSafStartStepsRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeSafStartStepsRequest) SetLanguage(v string) *DescribeSafStartStepsRequest {
	s.Language = &v
	return s
}

func (s *DescribeSafStartStepsRequest) SetRegId(v string) *DescribeSafStartStepsRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSafStartStepsRequest) SetServerRegion(v string) *DescribeSafStartStepsRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribeSafStartStepsRequest) Validate() error {
	return dara.Validate(s)
}
