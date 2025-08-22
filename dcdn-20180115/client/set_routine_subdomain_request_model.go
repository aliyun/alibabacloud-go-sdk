// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineSubdomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubdomains(v map[string]interface{}) *SetRoutineSubdomainRequest
	GetSubdomains() map[string]interface{}
}

type SetRoutineSubdomainRequest struct {
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
	Subdomains map[string]interface{} `json:"Subdomains,omitempty" xml:"Subdomains,omitempty"`
}

func (s SetRoutineSubdomainRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineSubdomainRequest) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainRequest) GetSubdomains() map[string]interface{} {
	return s.Subdomains
}

func (s *SetRoutineSubdomainRequest) SetSubdomains(v map[string]interface{}) *SetRoutineSubdomainRequest {
	s.Subdomains = v
	return s
}

func (s *SetRoutineSubdomainRequest) Validate() error {
	return dara.Validate(s)
}
