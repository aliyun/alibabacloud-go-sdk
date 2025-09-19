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
	// The type of proactive defense. Valid values:
	//
	// 	- **auto_breaking**: virus defense
	//
	// 	- **ransomware_breaking**: ransomware capture
	//
	// 	- **webshell_cloud_breaking**: webshell defense
	//
	// 	- **alinet**: malicious behavior defense
	//
	// 	- **alisecguard**: client protection
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
