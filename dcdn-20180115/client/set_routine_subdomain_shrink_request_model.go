// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineSubdomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubdomainsShrink(v string) *SetRoutineSubdomainShrinkRequest
	GetSubdomainsShrink() *string
}

type SetRoutineSubdomainShrinkRequest struct {
	// The parameters of the subdomain.
	//
	// The parameters are in the following format:
	//
	//     Subdomains: [
	//
	//         "subdomain-test"
	//
	//     ]
	//
	// This parameter is required.
	//
	// example:
	//
	// ["subdomain-test"]
	SubdomainsShrink *string `json:"Subdomains,omitempty" xml:"Subdomains,omitempty"`
}

func (s SetRoutineSubdomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineSubdomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainShrinkRequest) GetSubdomainsShrink() *string {
	return s.SubdomainsShrink
}

func (s *SetRoutineSubdomainShrinkRequest) SetSubdomainsShrink(v string) *SetRoutineSubdomainShrinkRequest {
	s.SubdomainsShrink = &v
	return s
}

func (s *SetRoutineSubdomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
