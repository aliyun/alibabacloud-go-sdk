// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspiciousOverallConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOverallConfig(v *DescribeSuspiciousOverallConfigResponseBodyOverallConfig) *DescribeSuspiciousOverallConfigResponseBody
	GetOverallConfig() *DescribeSuspiciousOverallConfigResponseBodyOverallConfig
	SetRequestId(v string) *DescribeSuspiciousOverallConfigResponseBody
	GetRequestId() *string
}

type DescribeSuspiciousOverallConfigResponseBody struct {
	// The configuration.
	OverallConfig *DescribeSuspiciousOverallConfigResponseBodyOverallConfig `json:"OverallConfig,omitempty" xml:"OverallConfig,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSuspiciousOverallConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousOverallConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousOverallConfigResponseBody) GetOverallConfig() *DescribeSuspiciousOverallConfigResponseBodyOverallConfig {
	return s.OverallConfig
}

func (s *DescribeSuspiciousOverallConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspiciousOverallConfigResponseBody) SetOverallConfig(v *DescribeSuspiciousOverallConfigResponseBodyOverallConfig) *DescribeSuspiciousOverallConfigResponseBody {
	s.OverallConfig = v
	return s
}

func (s *DescribeSuspiciousOverallConfigResponseBody) SetRequestId(v string) *DescribeSuspiciousOverallConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspiciousOverallConfigResponseBody) Validate() error {
	if s.OverallConfig != nil {
		if err := s.OverallConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSuspiciousOverallConfigResponseBodyOverallConfig struct {
	// The status of the feature. Valid values:
	//
	// 	- **off**: disabled
	//
	// 	- **on**: enabled
	//
	// example:
	//
	// on
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- **auto_breaking**: Anti-Virus
	//
	// 	- **ransomware_breaking**: Anti-ransomware (Bait Capture)
	//
	// 	- **webshell_cloud_breaking**: Webshell Protection
	//
	// 	- **alinet**: Behavior prevention
	//
	// 	- **k8s_log_analysis**: K8s Threat Detection
	//
	// 	- **alisecguard**: Defense mode for Client Protection
	//
	// example:
	//
	// auto_breaking
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSuspiciousOverallConfigResponseBodyOverallConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousOverallConfigResponseBodyOverallConfig) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousOverallConfigResponseBodyOverallConfig) GetConfig() *string {
	return s.Config
}

func (s *DescribeSuspiciousOverallConfigResponseBodyOverallConfig) GetType() *string {
	return s.Type
}

func (s *DescribeSuspiciousOverallConfigResponseBodyOverallConfig) SetConfig(v string) *DescribeSuspiciousOverallConfigResponseBodyOverallConfig {
	s.Config = &v
	return s
}

func (s *DescribeSuspiciousOverallConfigResponseBodyOverallConfig) SetType(v string) *DescribeSuspiciousOverallConfigResponseBodyOverallConfig {
	s.Type = &v
	return s
}

func (s *DescribeSuspiciousOverallConfigResponseBodyOverallConfig) Validate() error {
	return dara.Validate(s)
}
