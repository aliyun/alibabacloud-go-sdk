// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGlobalizationConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalizationConfig(v *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) *GetInstanceGlobalizationConfigResponseBody
	GetGlobalizationConfig() *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig
	SetRequestId(v string) *GetInstanceGlobalizationConfigResponseBody
	GetRequestId() *string
}

type GetInstanceGlobalizationConfigResponseBody struct {
	GlobalizationConfig *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig `json:"GlobalizationConfig,omitempty" xml:"GlobalizationConfig,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceGlobalizationConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGlobalizationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceGlobalizationConfigResponseBody) GetGlobalizationConfig() *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig {
	return s.GlobalizationConfig
}

func (s *GetInstanceGlobalizationConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceGlobalizationConfigResponseBody) SetGlobalizationConfig(v *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) *GetInstanceGlobalizationConfigResponseBody {
	s.GlobalizationConfig = v
	return s
}

func (s *GetInstanceGlobalizationConfigResponseBody) SetRequestId(v string) *GetInstanceGlobalizationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceGlobalizationConfigResponseBody) Validate() error {
	if s.GlobalizationConfig != nil {
		if err := s.GlobalizationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig struct {
	// 语言
	//
	// example:
	//
	// zh-Hans-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 时区
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) GetLanguage() *string {
	return s.Language
}

func (s *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) SetLanguage(v string) *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig {
	s.Language = &v
	return s
}

func (s *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) SetTimeZone(v string) *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig {
	s.TimeZone = &v
	return s
}

func (s *GetInstanceGlobalizationConfigResponseBodyGlobalizationConfig) Validate() error {
	return dara.Validate(s)
}
