// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConcernNecessityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcernNecessity(v string) *ModifyConcernNecessityRequest
	GetConcernNecessity() *string
}

type ModifyConcernNecessityRequest struct {
	// The priorities to fix the vulnerabilities. Valid values:
	//
	// 	- asap: high
	//
	// 	- later: medium
	//
	// 	- nntf: low
	//
	// example:
	//
	// asap,nntf
	ConcernNecessity *string `json:"ConcernNecessity,omitempty" xml:"ConcernNecessity,omitempty"`
}

func (s ModifyConcernNecessityRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyConcernNecessityRequest) GoString() string {
	return s.String()
}

func (s *ModifyConcernNecessityRequest) GetConcernNecessity() *string {
	return s.ConcernNecessity
}

func (s *ModifyConcernNecessityRequest) SetConcernNecessity(v string) *ModifyConcernNecessityRequest {
	s.ConcernNecessity = &v
	return s
}

func (s *ModifyConcernNecessityRequest) Validate() error {
	return dara.Validate(s)
}
