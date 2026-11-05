// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspiciousUUIDConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeSuspiciousUUIDConfigRequest
	GetType() *string
}

type DescribeSuspiciousUUIDConfigRequest struct {
	// The defense type of the proactive defense to query. Valid values:
	//
	// - **auto_breaking**: Anti-virus.
	//
	// - **ransomware_breaking**: Anti-ransomware (bait capture).
	//
	// - **webshell_cloud_breaking**: Web shell connection defense.
	//
	// - **alinet**: Malicious network behavior defense.
	//
	// - **alisecguard**: Client self-protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// alinet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSuspiciousUUIDConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousUUIDConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousUUIDConfigRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSuspiciousUUIDConfigRequest) SetType(v string) *DescribeSuspiciousUUIDConfigRequest {
	s.Type = &v
	return s
}

func (s *DescribeSuspiciousUUIDConfigRequest) Validate() error {
	return dara.Validate(s)
}
