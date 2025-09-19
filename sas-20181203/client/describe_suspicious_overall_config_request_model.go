// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspiciousOverallConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeSuspiciousOverallConfigRequest
	GetSourceIp() *string
	SetType(v string) *DescribeSuspiciousOverallConfigRequest
	GetType() *string
}

type DescribeSuspiciousOverallConfigRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 39.161.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// auto_breaking
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSuspiciousOverallConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousOverallConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousOverallConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspiciousOverallConfigRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSuspiciousOverallConfigRequest) SetSourceIp(v string) *DescribeSuspiciousOverallConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspiciousOverallConfigRequest) SetType(v string) *DescribeSuspiciousOverallConfigRequest {
	s.Type = &v
	return s
}

func (s *DescribeSuspiciousOverallConfigRequest) Validate() error {
	return dara.Validate(s)
}
