// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTwoFactorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *GetInstanceTwoFactorResponseBodyConfig) *GetInstanceTwoFactorResponseBody
	GetConfig() *GetInstanceTwoFactorResponseBodyConfig
	SetRequestId(v string) *GetInstanceTwoFactorResponseBody
	GetRequestId() *string
}

type GetInstanceTwoFactorResponseBody struct {
	// The settings of two-factor authentication.
	Config *GetInstanceTwoFactorResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceTwoFactorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTwoFactorResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponseBody) GetConfig() *GetInstanceTwoFactorResponseBodyConfig {
	return s.Config
}

func (s *GetInstanceTwoFactorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceTwoFactorResponseBody) SetConfig(v *GetInstanceTwoFactorResponseBodyConfig) *GetInstanceTwoFactorResponseBody {
	s.Config = v
	return s
}

func (s *GetInstanceTwoFactorResponseBody) SetRequestId(v string) *GetInstanceTwoFactorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceTwoFactorResponseBodyConfig struct {
	// Indicates whether two-factor authentication is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTwoFactor *bool `json:"EnableTwoFactor,omitempty" xml:"EnableTwoFactor,omitempty"`
	// The duration within which two-factor authentication is not required after a local user passes two-factor authentication. Valid values: `0 to 168`. Unit: hours.
	//
	// > If 0 is returned, a local user must pass two-factor authentication every time the local user logs on to the bastion host.
	//
	// example:
	//
	// 1
	SkipTwoFactorTime *int64 `json:"SkipTwoFactorTime,omitempty" xml:"SkipTwoFactorTime,omitempty"`
	// The two-factor authentication methods.
	TwoFactorMethods []*string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
}

func (s GetInstanceTwoFactorResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTwoFactorResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponseBodyConfig) GetEnableTwoFactor() *bool {
	return s.EnableTwoFactor
}

func (s *GetInstanceTwoFactorResponseBodyConfig) GetSkipTwoFactorTime() *int64 {
	return s.SkipTwoFactorTime
}

func (s *GetInstanceTwoFactorResponseBodyConfig) GetTwoFactorMethods() []*string {
	return s.TwoFactorMethods
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetEnableTwoFactor(v bool) *GetInstanceTwoFactorResponseBodyConfig {
	s.EnableTwoFactor = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetSkipTwoFactorTime(v int64) *GetInstanceTwoFactorResponseBodyConfig {
	s.SkipTwoFactorTime = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetTwoFactorMethods(v []*string) *GetInstanceTwoFactorResponseBodyConfig {
	s.TwoFactorMethods = v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
