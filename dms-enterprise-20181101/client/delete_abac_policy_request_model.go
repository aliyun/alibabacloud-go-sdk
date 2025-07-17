// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAbacPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbacPolicyId(v int64) *DeleteAbacPolicyRequest
	GetAbacPolicyId() *int64
	SetTid(v int64) *DeleteAbacPolicyRequest
	GetTid() *int64
}

type DeleteAbacPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12****
	AbacPolicyId *int64 `json:"AbacPolicyId,omitempty" xml:"AbacPolicyId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteAbacPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAbacPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAbacPolicyRequest) GetAbacPolicyId() *int64 {
	return s.AbacPolicyId
}

func (s *DeleteAbacPolicyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteAbacPolicyRequest) SetAbacPolicyId(v int64) *DeleteAbacPolicyRequest {
	s.AbacPolicyId = &v
	return s
}

func (s *DeleteAbacPolicyRequest) SetTid(v int64) *DeleteAbacPolicyRequest {
	s.Tid = &v
	return s
}

func (s *DeleteAbacPolicyRequest) Validate() error {
	return dara.Validate(s)
}
